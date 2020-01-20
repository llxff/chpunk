package textfile

import (
	"chpunk/translation"
	"fmt"
	"io"
	"log"
	"os"
)

func Export(c []*translation.Content, fileName string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	for _, t := range c {
		appendToFile(t, f)
	}

	err = f.Close()

	if err != nil {
		log.Fatalln(err)
	}
}

func appendToFile(c *translation.Content, f io.Writer) {
	fmt.Fprintln(f, c.Text)

	if !c.IsNewParagraph() {
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, c.Yandex)
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, c.Deepl)
	}

	fmt.Fprintln(f, "")

}
