package textfile

import (
	"chpunk/translation"
	"fmt"
	"io"
	"os"
)

type Container struct {
	FileName string
}

func (c *Container) Export(translations []*translation.Content) error {
	f, err := os.Create(c.FileName)

	if err != nil {
		return err
	}

	for _, t := range translations {
		appendToFile(t, f)
	}

	err = f.Close()

	if err != nil {
		return err
	}

	return nil
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
