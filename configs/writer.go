package configs

import (
	"bytes"
	"github.com/BurntSushi/toml"
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
	config.JaneName = "jane"
	config.JaneEmoji = ":game_die:"
	config.JaneChannel = "#general"
	config.SlackToken = ""
	config.BambooUser = ""
	config.BambooPass = ""
	config.BambooChannels = make(map[string]string)
	config.BambooChannels["*"] = "#general"

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		log.Println(err)
	}

	if _, err := fo.Write(buf.Bytes()); err != nil {
		log.Println(err)
	}

}
