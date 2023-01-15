package main

import (
	"log"
	"os"

	"github.com/musaprg/annict-epgstation-connector/internal/cmd"
)

func main() {
	if err := cmd.Root().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
