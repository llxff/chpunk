package yandex

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

var apiID = os.Getenv("YANDEX_ID")

type YandexTranslation struct {
	Text []string
}

func Translate(text string) string {
	formData := url.Values{
		"text": {text},
	}

	resp, err := http.PostForm("https://translate.yandex.net/api/v1/tr.json/translate?id="+apiID+"&srv=tr-text&lang=en-ru&reason=auto", formData)

	if err != nil {
		log.Fatalln(err)
	}

	var result YandexTranslation

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if len(result.Text) == 0 {
		return "https://translate.yandex.ru/?lang=en-ru&text=" + url.QueryEscape(text)
	} else {
		decodedValue, err := url.QueryUnescape(result.Text[0])

		if err != nil {
			log.Fatalln(err)
		}

		return decodedValue
	}
}
