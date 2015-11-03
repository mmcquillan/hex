package models

import (
	"encoding/json"
	"log"
	"os"
)

func WriteDefaultConfig(location string) {

	fo, err := os.Create(location)
	if err != nil {
		log.Println(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			log.Println(err)
		}
	}()

	var config Config

	b, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
	}

	if _, err := fo.Write(b); err != nil {
		log.Println(err)
	}

}
