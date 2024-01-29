package main

import (
	"log"
	"os"

	"github.com/miscord-dev/annict-epgstation-connector/internal/cmd"
)

func main() {
	if err := cmd.Root().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
