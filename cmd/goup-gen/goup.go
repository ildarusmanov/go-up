package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()

	cmd := flag.Arg(0)
	wdir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Can not detect current directory: %s", err)
	}

	switch cmd {
	case "init":
		InitCommand(wdir)
	case "update":
		UpdateCommand(wdir)
	default:
		log.Printf("Invalid command")
	}

	log.Println("Done!")
}
