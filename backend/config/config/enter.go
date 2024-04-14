package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server server `yaml:"server"`
	Log    log    `yaml:"log"`
	Db     db     `yaml:"mysql"`
}

type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

var Config *config

type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

func init() {
	yamlFile, err := os.ReadFile("./settings.yaml")
	//有错就down机
	if err != nil {
		panic(err)
	}
	//绑定值
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
