package deepl

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Text struct {
	Text string
}

type Translation struct {
	Translations []Text
}

func Translate(apiKey string, text string) string {
	formData := url.Values{
		"auth_key":    {apiKey},
		"text":        {text},
		"source_lang": {"en"},
		"target_lang": {"ru"},
	}

	resp, err := http.PostForm("https://api.deepl.com/v2/translate", formData)

	if err != nil {
		log.Fatalln(err)
	}

	var result Translation

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	return result.Translations[0].Text
}
