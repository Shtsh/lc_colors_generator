package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const colorsDefinitions = "https://jonasjacek.github.io/colors/data.json"
const defaultDefinitionsPath = ".cache/lc_colors_definitions.json"


type ColorDefinition struct {
	ColorId int `json:"colorId"`
	Name string `json:"name"`
}

func downloadDefinitions(destination string) error {
	log.Printf("Downloading %s -> %s", colorsDefinitions, destination) 
	resp, err := http.Get(colorsDefinitions)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}


func loadDefinitions() (map[string]string, error) {
	homeDir := os.Getenv("HOME")
	filePath := filepath.Join(homeDir, defaultDefinitionsPath)
	definitions := []ColorDefinition {}
	result := make(map[string]string, 256)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err = downloadDefinitions(filePath); err != nil {
			log.Println("Cannot download ", defaultDefinitionsPath)
			return result, err
		}		
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening ", filePath, err)
		return result, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&definitions)

	for _, definition := range definitions {
		result[definition.Name] = strconv.Itoa(definition.ColorId)
	}
	
	return result, nil
}	
