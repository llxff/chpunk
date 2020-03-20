# chpunk

Helps with books translation.

## Usage
```
$ go run cmd/chpunk/main.go
Usage:
  translate [command]

Available Commands:
  doc         Translates a given Google Spreadsheet and past translation to a given Google Doc
  file        Translates a given file
  help        Help about any command
  server      Starts server
  sheet       Translates a given Google Spreadsheet

Flags:
  -h, --help   help for translate

Use "translate [command] --help" for more information about a command.
```

Setup:

1. Enable Google Docs API https://developers.google.com/docs/api/quickstart/go
2. Set api keys for Deepl and Yandex translators to `translators.json` (The project uses public version of [Yandex Translate](https://translate.yandex.ru/), for some reasons it makes better translation than official API)

## This is a pet-project

I wrote it to help my wife 👩‍💻 with her work, don't expect much from it. No tests, DRY, SOLID and other buzzwords, this code has been written for fun.
