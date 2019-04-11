package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config struct contains token related info from messenger
type Config struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

func parseContentFile() string {
	contentFile, err := ioutil.ReadFile("content.yml")

	if err != nil {
		log.Printf("Error opening content file: %s\n\n", err)
		panic(err)
	}

	er, _ := yaml.Marshal(contentFile)

	if er != nil {
		log.Printf("Couldn't marshal content file: $s\n\n", er)
	}

	return string(er)
}

func (c *Config) readYml() *Config {
	yamlFile, err := ioutil.ReadFile("resources/bot.config.yml")
	if err != nil {
		log.Printf("Error opening content file: %s\n\n", err)
	}

	er := yaml.Unmarshal(yamlFile, c)

	if er != nil {
		log.Printf("Couldn't marshal content file: $s\n\n", er)
	}

	log.Printf("parsed content: %s\n\n", c)
	return c
}

func getToken() string {
	var c Config
	c.readYml()
	v, err := json.Marshal(c)

	if err != nil {
		log.Printf("error marshalling json file: %s\n\n", err)
	}

	log.Printf(c.AccessToken)
	//parsedFile := string(v)
	return string(v)
}
