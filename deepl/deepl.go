package deepl

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type DeeplText struct {
	Text string
}

type DeeplTranslation struct {
	Translations []DeeplText
}

func Translate(text string) string {
	formData := url.Values{
		"auth_key":    {os.Getenv("DEEPL_KEY")},
		"text":        {text},
		"source_lang": {"en"},
		"target_lang": {"ru"},
	}

	resp, err := http.PostForm("https://api.deepl.com/v1/translate", formData)

	if err != nil {
		log.Fatalln(err)
	}

	var result DeeplTranslation

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	return result.Translations[0].Text
}
