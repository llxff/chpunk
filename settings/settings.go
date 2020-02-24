package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Deepl  string `json:"deepl"`
	Yandex string `json:"yandex"`
}

func Get() *Config {
	data, err := ioutil.ReadFile("settings.json")
	if err != nil {
		log.Fatalln(err)
	}

	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Fatalln(err)
	}

	return &c
}
