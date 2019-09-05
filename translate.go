package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type DeeplText struct {
	Text string
}

type DeeplTranslation struct {
	Translations []DeeplText
}

type YandexTranslation struct {
	Text []string
}

func main() {
	lines := chapter()
	translations := translate(lines)

	f, err := os.Create("translation.txt")

	if err != nil {
		log.Fatalln(err)
	}

	for ix := range translations {
		fmt.Fprintln(f, lines[ix])
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, translations[ix]["yandex"])
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, translations[ix]["deepl"])
		fmt.Fprintln(f, "")
	}

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func translate(lines []string) []map[string]string {
	translations := make([]map[string]string, len(lines), len(lines))

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	lock := sync.Mutex{}

	for ix, line := range lines {
		go func(ix int, line string) {
			translation := map[string]string{
				"yandex": yandex(line),
				"deepl":  deepl(line),
			}

			lock.Lock()
			fmt.Printf("Line %d has been translated\n", ix)

			translations[ix] = translation
			lock.Unlock()

			wg.Done()
		}(ix, line)
	}

	wg.Wait()

	return translations
}

func chapter() []string {
	csvfile, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	lines := make([]string, 0, 10)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, record[0])
	}

	return lines
}

func deepl(text string) string {
	if text == "" || text == "***" {
		return text
	}

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

	decodedValue, err := url.QueryUnescape(result.Translations[0].Text)

	if err != nil {
		log.Fatalln(err)
	}

	return decodedValue
}

func yandex(text string) string {
	if text == "" || text == "***" {
		return text
	}

	formData := url.Values{
		"text": {text},
	}

	resp, err := http.PostForm("https://translate.yandex.net/api/v1/tr.json/translate?id=6617361a.5d6f7b83.73baf1c9-2-0&srv=tr-text&lang=en-ru&reason=auto", formData)

	if err != nil {
		log.Fatalln(err)
	}

	var result YandexTranslation

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

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
