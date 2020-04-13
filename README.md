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

1. Enable Google Docs API https://developers.google.com/docs/api/quickstart/go and put credentials file to `configs/google.json`
2. Set api keys for Deepl and Yandex translators to `configs/translators.json` (The project uses public version of [Yandex Translate](https://translate.yandex.ru/), for some reasons it makes better translation than official API)

## Frontend:

Frontend uses Vue.js and yarn.

To start just frontend:

`$ make frontend`

To start frontend with API:

`$ make`

<img src="https://user-images.githubusercontent.com/866273/79149739-e4deec00-7dc7-11ea-92d2-955569fb5988.gif" />

## This is a pet-project

I wrote it to help my wife 👩‍💻 with her work, don't expect much from it. No tests, DRY, SOLID and other buzzwords, this code has been written for fun.
