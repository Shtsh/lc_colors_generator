package main


import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)


type ConfigFile struct {
	ColorDefinitions map[string]string `yaml:"color_definitions"`
	DoNotProcess map[string]string `yaml:"do_not_process"`
	ExtensionGroups map[string][]string `yaml:"extension_groups"`
	ColorMapping map[string]string `yaml:"color_mapping"`
	
}

type readFunction func(string) ([]byte, error)


func ReadConfig(configFile string) (*ConfigFile, error) {
	return readConfig(configFile, ioutil.ReadFile)
}


func readConfig(configFile string, read readFunction) (*ConfigFile, error) {
	configData := ConfigFile{}
	fileData, err := read(configFile)
	if err != nil {
		log.Fatal("Cannot read ", configFile, err)
	}
	err = yaml.Unmarshal([]byte(fileData), &configData)
	if err != nil {
		log.Fatal("Cannot unmarshall yaml from ", configFile, err)
	}
	return &configData, nil
}
