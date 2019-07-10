package main

import (
	"os"
)

func main() {
	extensions := PopulateExtensions()
	keys := make([]string, 0, len(extensions))
	for k:= range extensions {
		keys = append(keys, k)
	}
	if len(os.Args) != 3 {
		panic("Icorrect Parameters")
	}
	filePath := os.Args[1]
	savePath := os.Args[2]
	config, _ := ReadConfig(filePath)
	lcconfig, _ := GenerateLCConfigStruct(config)
	lcconfig.WriteConfigFile(savePath)
}
