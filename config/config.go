package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func init() {
	loadConfiguration()
}

//Config config information
var Config config

type config struct {
	Name       string `yaml:"name"`
	Endpoint   string `yaml:"endpoint"`
	SecretName string `yaml:"secret_name"`
	Namespace  string `yaml:"namespace"`
	Mysql      mysql  `yaml:"mysql"`
}
type mysql struct {
	Addr   string `yaml:"addr"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	DBName string `yaml:"db_name"`
}

func loadConfiguration() error {
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	return nil
}
