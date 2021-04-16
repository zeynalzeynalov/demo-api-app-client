package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var config Configuration

type Configuration struct {

	env string

	PrintLineCharLength	int

	TaxdooAPIWithJavaURL string

	WebServerPort int

	GolangServerLogFilePath string
}

func readConfiguration() {

	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)
}
