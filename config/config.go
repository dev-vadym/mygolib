package config

import (
	"github.com/dev-vadym/mygolib/logger"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func NewConfig(path string, debug bool, out interface{}) {
	log := logger.Use.WithField("config", "config.yml")
	log.Info("Loading the configuration file...")

	yamlFile, errF := ioutil.ReadFile(path + "/config.yml")
	if errF != nil {
		log.Fatal(errF)
	}

	err := yaml.Unmarshal(yamlFile, out)
	if err != nil {
		log.Fatal(err)
	}
}
