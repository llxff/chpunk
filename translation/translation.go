package translation

import (
	"chpunk/settings"
	"chpunk/translators/deepl"
	"chpunk/translators/yandex"
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

func Translate(config settings.Config, lines []string) []*Content {
	translations := make([]*Content, len(lines))

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	lock := sync.Mutex{}

	for ix, line := range lines {
		go func(ix int, line string) {
			translation := translateContent(config, line)

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

func translateContent(config settings.Config, text string) *Content {
	if text == NewParagraph {
		return &Content{Text: NewParagraph}
	} else {
		return &Content{
			Text:   text,
			Yandex: yandex.Translate(config.Yandex, text),
			Deepl:  deepl.Translate(config.Deepl, text),
		}
	}
}
