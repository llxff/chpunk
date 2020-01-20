package translation

import (
	"chpunk/deepl"
	"chpunk/yandex"
	"fmt"
	"sync"
)

const NewParagraph = "***"

type Content struct {
	Text   string
	Yandex string
	Deepl  string
}

func (c *Content) IsNewParagraph() bool {
	return c.Text == NewParagraph
}

func Translate(lines []string) []*Content {
	translations := make([]*Content, len(lines))

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	lock := sync.Mutex{}

	for ix, line := range lines {
		go func(ix int, line string) {
			translation := translateContent(line)

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

func translateContent(text string) *Content {
	if text == NewParagraph {
		return &Content{Text: NewParagraph}
	} else {
		return &Content{
			Text:   text,
			Yandex: yandex.Translate(text),
			Deepl:  deepl.Translate(text),
		}
	}
}
