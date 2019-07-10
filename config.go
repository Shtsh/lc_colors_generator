package main


import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)


type ConfigFile struct {
	Test string `yaml:"test"`
	ColorDefinitions map[string]string `yaml:"color_definitions"`
	DoNotProcess map[string]string `yaml:"do_not_process"`
	ExtensionGroups map[string][]string `yaml:"extension_groups"`
	ColorMapping map[string]string `yaml:"color_mapping"`
	
}

func ReadConfig(configFile string) (*ConfigFile, error) {
	configData := ConfigFile{}
	fileData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Cannot read", configFile, err)
	}
	err = yaml.Unmarshal([]byte(fileData), &configData)
	if err != nil {
		log.Fatal("Cannot unmarshall yaml from", configFile, err)
	}
	_, err = yaml.Marshal(&configData)
	return &configData, nil
}
