package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func parseJsonFile(filename string, mappedType interface{}) (interface{}, bool) {
	var file *os.File
	newlyCreated := false

	if !fileExists(filename) {
		file = createConfigFile(filename)
		newlyCreated = true
	}

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal([]byte(byteValue), &mappedType)

	return mappedType, newlyCreated
}

func fileExists(filename string) bool {
	_, err := os.Open(filename)
	if err != nil {
		return false
	}
	return true
}

func createConfigFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
