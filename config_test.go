package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var mockReadContent string


func mockReadFile(configFile string) ([]byte, error) {
	return []byte(mockReadContent), nil
}

func TestReadConfig(t *testing.T) {
	mockReadContent = `---
color_definitions:
  blue: 123;123
  red: 0
  black: 3999887
do_not_process:
  test: blue
  test2: 12345
extension_groups:
  group1: [ex1, ext2]
color_mapping:
  group1: black
  does_not_exists: 9889
`
	configData, _ := readConfig("", mockReadFile)
	assert.Equal(t, configData.ColorDefinitions["blue"], "123;123")
}
