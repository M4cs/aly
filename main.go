package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("aly", "Alias manager for all platforms")
	// var load *bool = parser.Flag("l", "load", &argparse.Options{Required: false, Help: "Load aliases into shell session. Should be added to profile.", Default: false})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	config, err := checkForConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.Version)
}
