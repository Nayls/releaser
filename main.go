package main

import (
	"log"
	"os"

	"github.com/Nayls/releaser/cmd/releaser"
)

func init() {
	if err := releaser.InitCobraConfig().Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
}
