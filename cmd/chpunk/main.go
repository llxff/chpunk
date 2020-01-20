package main

import (
	"chpunk/export/textfile"
	"chpunk/import/csvfile"
	"chpunk/translation"
	"os"
)

func main() {
	lines := csvfile.Import(os.Args[1])
	translations := translation.Translate(lines)

	textfile.Export(translations, "translation.txt")
}
