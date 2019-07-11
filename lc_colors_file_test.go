package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

	
func TestGenerateLCConfig(t *testing.T) {
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
  does_not_exist: 9889
`
	origData, _ := readConfig("", mockReadFile)
	var lcData *LCConfig
	var err error
	definitions := map[string]string {}
	lcData, err = GenerateLCConfigStruct(origData, definitions)
	assert.NoError(t, err)
	assert.Contains(t, lcData.extensions, "ex1")

	data := lcData.extensions["ex1"]
	assert.Equal(t, data, "3999887")
	
	assert.Contains(t, lcData.extensions, "test")	
	data = lcData.extensions["test"]
	assert.Equal(t, data, "123;123")

	assert.Contains(t, lcData.extensions, "test2")	
	data = lcData.extensions["test2"]
	assert.Equal(t, data, "12345")

	assert.Contains(t, lcData.groups, "group1")
	assert.Contains(t, lcData.groups["group1"], "ex1")
	assert.Contains(t, lcData.groups["group1"], "ext2")
}


func TestGetExtensionLine(t *testing.T) {
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
  does_not_exist: 9889
`
	origData, _ := readConfig("", mockReadFile)
	var lcData *LCConfig
	var err error
	definitions := map[string]string {}
	lcData, err = GenerateLCConfigStruct(origData, definitions)
	assert.NoError(t, err)

	line, err := lcData.getExtensionLine("ex1")
	assert.NoError(t, err)
	assert.Equal(t, line, "ex1 3999887\n")

	line, err = lcData.getExtensionLine("test")
	assert.NoError(t, err)
	assert.Equal(t, line, "test 123;123\n")
}


func TestGetGroupContent(t *testing.T) {
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
  does_not_exist: 9889
`
	expected := `
Group group1
ex1 3999887
ext2 3999887
`
	origData, _ := readConfig("", mockReadFile)
	var lcData *LCConfig
	var err error
	definitions := map[string]string {}
	lcData, err = GenerateLCConfigStruct(origData, definitions)
	assert.NoError(t, err)

	line, err := lcData.getGroupContent("group1")
	assert.NoError(t, err)
	assert.Equal(t, *line, expected)
}


