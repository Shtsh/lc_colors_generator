package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const default_cache_path = ".cache/lc_colors_generator"

func PopulateExtensions(cachePath ...string) map[string]bool {
	homeDir := os.Getenv("HOME")
	filePath := filepath.Join(homeDir, default_cache_path)
	if len(cachePath) > 0 {
		filePath = cachePath[0]
	}
	extensions, err := LoadExtensionsFromFile(filePath)
	if err != nil {
		extensions = RereadExtensions(filePath)
	}
	return extensions
}

func RereadExtensions(filePath string) map[string]bool {
	extensions := runLocate()
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Cannot create cache file: ", err)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(extensions)
	if err != nil {
		log.Fatal("Cannot encode extensions list: ", err)
	}
	file.Close()
	return extensions
}

func LoadExtensionsFromFile(filePath string) (map[string]bool, error) {
	returnData := map[string]bool {}
	file, err := os.Open(filePath)
	if err != nil {
		return returnData, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&returnData)
	if err != nil {
		return returnData, err
	}
	file.Close()
	return returnData, nil
}

func processLine(line string, set map[string]bool) {
	filePath := filepath.Base(line)
	parts := strings.Split(filePath, ".")
	if len(parts) > 1 {
		extension := parts[len(parts) - 1]
		set[extension] = true
	}
}

func runLocate() map[string]bool {
	set := map[string]bool {}
	command := exec.Command("locate", "*.*")
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Fatal("StdoutPipe: ", err)
	}
	err = command.Start()
	if err != nil {
		log.Fatal("Start: ", err)
	}
	bufferSize := make([]byte, 1000)
	reader := bufio.NewReader(stdout)
	for {
		_, err = stdout.Read(bufferSize)
		if err != nil {
			break
		}
		line, _, _ := reader.ReadLine()
		processLine(string(line), set)		
	}
	
	err = command.Wait()
	if err != nil {
		log.Fatal("Wait: ", err, err.(*exec.ExitError).Stderr)
	}
	return set
	
}
