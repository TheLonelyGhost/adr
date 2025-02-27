package main

import (
	"log"

	"github.com/thelonelyghost/adr/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
