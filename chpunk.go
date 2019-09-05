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

type Translation struct {
	Text   string
	Yandex string
	Deepl  string
}

func Translate(text string) *Translation {
	if text == "" {
		return &Translation{Text: "***"}
	} else {
		return &Translation{Text: text, Yandex: yandex(text), Deepl: deepl(text)}
	}
}

func (t *Translation) AppendToFile(f io.Writer) {
	fmt.Fprintln(f, t.Text)

	if t.Text != "***" {
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, t.Yandex)
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, t.Deepl)
	}

	fmt.Fprintln(f, "")
}

func deepl(text string) string {
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

func main() {
	lines := chapter()
	translations := translate(lines)

	f, err := os.Create("translation.txt")

	if err != nil {
		log.Fatalln(err)
	}

	for _, translation := range translations {
		translation.AppendToFile(f)
	}

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
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

func translate(lines []string) []*Translation {
	translations := make([]*Translation, len(lines), len(lines))

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	lock := sync.Mutex{}

	for ix, line := range lines {
		go func(ix int, line string) {
			translation := Translate(line)

			lock.Lock()
			translations[ix] = translation
			lock.Unlock()

			fmt.Printf("Line %d has been translated\n", ix)

			wg.Done()
		}(ix, line)
	}

	wg.Wait()

	return translations
}
