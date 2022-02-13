package main

import (
	"log"

	cmd "github.com/chalfel/go-lead-store/cmd/crypto-store"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
