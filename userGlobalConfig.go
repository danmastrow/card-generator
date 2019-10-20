package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserGlobalConfig struct {
	cardWidth  float64
	cardHeight float64
}

func getUserGlobalConfig(filename string) (UserGlobalConfig, bool) {
	var file *os.File
	var result UserGlobalConfig
	newlyCreated := false

	if !fileExists(filename) {
		file = createFile(filename)
		newlyCreated = true
	} else {
		file, _ = os.Open(filename)
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println(err)
	}
	return result, newlyCreated
}

func fileExists(filename string) bool {
	_, err := os.Open(filename)
	if err != nil {
		return false
	}
	return true
}

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
