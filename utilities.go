package main

import (
	"log"
	"strings"
)

func printLine() {

	log.Println(strings.Repeat("_", config.PrintLineCharLength));
}
