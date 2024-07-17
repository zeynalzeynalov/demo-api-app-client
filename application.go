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

	log.Println("DemoApi Go App started!")
	fmt.Println("DemoApi Go App started!")

	handleRequests()
}
