package config

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var GlobalCFG *Config

func init() {
	var err error
	GlobalCFG, err = GetGlobalConfig("./config.toml")
	if err != nil {
		panic(err)
	}
}

type RabbitMQCFG struct {
	Host     string
	Port     string
	User     string
	Password string
}

type Config struct {
	Context   string
	Namespace string
	RabbitMQ  RabbitMQCFG
}

func GetGlobalConfig(filePath string) (cfg *Config, err error) {
	cfg = new(Config)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file error:", err)
		return nil, err
	}

	_, err = toml.Decode(string(file), cfg)
	if err != nil {
		fmt.Println("Decode error:", err)
		return nil, err
	}

	return
}
