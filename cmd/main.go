package main

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	PwFile struct {
		Name    string `yaml:"name"`
		Masukan map[string]struct {
			Fullname string `yaml:"fullname"`
			Password string `yaml:"password"`
		} `yaml:"entries"`
	} `yaml:"pw-file"`
}

func main() {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	c := Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("--- Config:\n%v\n\n", c)
}
