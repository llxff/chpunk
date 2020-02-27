package yandex

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Translation struct {
	Text []string
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

	decodedValue, err := url.QueryUnescape(result.Text[0])

	if err != nil {
		log.Fatalln(err)
	}

	return decodedValue
}
