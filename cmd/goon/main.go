package main

import (
	"log"

	"github.com/devkcud/goondl/internal/app"
)

func main() {
	if err := app.New(); err != nil {
		log.Fatal(err)
	}
}
