package nopaste

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Root    string `yaml:"root"`
	DataDir string `yaml:"data_dir"`
	Listen  string `yaml:"listen"`
}

func LoadConfig(file string) (*Config, error) {
	log.Println("loading config file", file)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	c := Config{}
	err = yaml.Unmarshal(content, &c)
	return &c, nil
}
