package main

import (
	"flag"
)

func main() {
	extensions := PopulateExtensions()
	sourceFlag := flag.String("source", "", "yaml configuration")
	destinationFlag := flag.String("destination", "","LC_CONFIG file destination")
	flag.Parse()

	if *sourceFlag == "" || *destinationFlag == "" {
		panic("Incorrect parameters")
	}
	
	keys := make([]string, 0, len(extensions))
	for k:= range extensions {
		keys = append(keys, k)
	}
	definitions, _ := loadDefinitions()
	
	config, _ := ReadConfig(*sourceFlag)
	lcconfig, _ := GenerateLCConfigStruct(config, definitions)
	lcconfig.WriteConfigFile(*destinationFlag)
}
