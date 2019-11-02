package main

import (
	"log"

	"github.com/ildarusmanov/go-up/cmd/goup-gen/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Printf("Error: %s", err.Error())
	}
}
