package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	readConfiguration()

	f, _ := os.Create(config.GolangServerLogFilePath)
	defer f.Close()
	log.SetOutput(f)

	log.Println("Taxdoo Go App started!")
	fmt.Println("Taxdoo Go App started!")

	handleRequests()
}
