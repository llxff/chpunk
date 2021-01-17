package settings

import (
	"chpunk/internal/translators/yandex"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Deepl  string `json:"deepl"`
	Yandex string
}

func Get() *Config {
	data, err := ioutil.ReadFile("configs/translators.json")
	if err != nil {
		panic(err)
	}

	var c Config

	if err = json.Unmarshal(data, &c); err != nil {
		panic(err)
	}

	yandexAPIKey, err := yandex.GetAPIKey()
	if err != nil {
		panic(err)
	}

	c.Yandex = yandexAPIKey

	return &c
}
