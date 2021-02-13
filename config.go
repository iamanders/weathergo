package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Check if config file exists
func configFileExists(home string, configFileName string) bool {
	info, err := os.Stat(path.Join(home, configFileName))
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return !info.IsDir()
}

// Create base config file
func createBaseConfigFile(home string, configFileName string) {
	content := "API_KEY=\nCITY=Stockholm\n"
	err := ioutil.WriteFile(path.Join(home, configFileName), []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
