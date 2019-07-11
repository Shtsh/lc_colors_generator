package main

import (
	"fmt"
	"os"
	"strings"
	"log"
)


const DefaultColorType = "00;38;5;"

type LCConfig struct {
	groups map[string][]string
	colors map[string]string
	extensions map[string]string
}

func (c *LCConfig) GetConfigContent() (*string, error) {
	var buffer strings.Builder
	var err error
	for group := range c.groups {
		var groupContent *string
		if groupContent, err = c.getGroupContent(group); err != nil {
			continue
		}
		buffer.WriteString(*groupContent)
	}
	result := buffer.String()
	return &result, nil
}


func (c *LCConfig) getGroupContent(name string) (*string, error) {
	var buffer strings.Builder
	var err error
	groupHeader := fmt.Sprintf("\nGroup %s\n", name)
	buffer.WriteString(groupHeader)
	for _, extension := range c.groups[name] {
		var extensionLine string
		if extensionLine, err = c.getExtensionLine(extension); err != nil {
			continue
		}
		buffer.WriteString(extensionLine)
	}
	result := buffer.String()
	return &result, nil
}

func (c *LCConfig) getExtensionLine(name string) (string, error) {
	if color, ok := c.extensions[name]; ok {
		return fmt.Sprintf("%s %s\n", name, color), nil
	}
	return "", fmt.Errorf("No extension %s", name)
}

func (c *LCConfig) WriteConfigFile(path string) error {
	log.Println("Saving LC_COLORS file to", path)
	var configContent *string
	var err error
	if configContent, err = c.GetConfigContent(); err != nil {
		return err
	}
	var file *os.File
	if file, err = os.Create(path); err != nil {
		return err
	}
	if _, err = file.WriteString(*configContent); err != nil {
		return err
	}
	file.Close()
	
	return nil
}

func (c *LCConfig) UpdateColors(newColors map[string]string) {
	for k, v := range newColors {
		c.colors[k] = v
	}
}

func (c *LCConfig) generateColorCode(name string) string {
	result := name
	if colorCode, ok := c.colors[name]; ok {
		result = colorCode
	}
	return result
}

func NewLCConfig(colorDefinitions map[string]string) *LCConfig {
	config := LCConfig{}
	config.groups = map[string][]string {}
	config.extensions = map[string]string {}
	config.colors = map[string]string {}
	for k, v := range colorDefinitions {
		config.colors[k] = DefaultColorType + v
	}
	return &config
}

func GenerateLCConfigStruct(orig *ConfigFile, definitions map[string]string) (*LCConfig, error) {
	config := NewLCConfig(definitions)
	config.groups = orig.ExtensionGroups
	config.UpdateColors(orig.ColorDefinitions)
	
	for groupName, groupExtensions := range orig.ExtensionGroups {
		var color string
		if colorName, ok := orig.ColorMapping[groupName]; ok {
			color = config.generateColorCode(colorName)
		}
		for _, extension := range groupExtensions {
			config.extensions[extension] = color
		}
	}
	groupName := "do_not_process"
	config.groups[groupName] = []string {}
	for extensionName, extensionColor := range orig.DoNotProcess {
		config.extensions[extensionName] = config.generateColorCode(extensionColor)
		config.groups[groupName] = append(config.groups[groupName], extensionName)
	}
		
	return config, nil
}
