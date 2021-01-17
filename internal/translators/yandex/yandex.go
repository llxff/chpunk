package yandex

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Translation struct {
	Text []string
}

func GetAPIKey() (string, error) {
	req, err := http.NewRequest("GET", "https://translate.yandex.ru/", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authority", "translate.yandex.ru")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "en-GB,en;q=0.9")
	req.Header.Set("Referer", "https://translate.yandex.ru/")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	html := string(r)
	ind := strings.Index(html, "Ya.reqid = '") + 12

	return html[ind:ind+41] + "-0-0", nil
}

func Translate(apiKey string, text string) string {
	u := url.URL{
		Scheme:   "https",
		Host:     "translate.yandex.net",
		Path:     "/api/v1/tr.json/translate",
		RawQuery: "id=" + apiKey + "&srv=tr-text&lang=en-ru&reason=auto",
	}

	formData := url.Values{
		"text": {text},
	}
	resp, err := http.PostForm(u.String(), formData)

	if err != nil {
		log.Fatalln(err)
	}

	var result Translation

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if len(result.Text) == 0 {
		return "https://translate.yandex.ru/?lang=en-ru&text=" + url.QueryEscape(text)
	}

	return result.Text[0]
}
