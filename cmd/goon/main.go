package main

import (
	"log"

	"github.com/devkcud/goondl/internal/app"
)

func main() {
	if err := app.New().Run(); err != nil {
		log.Fatal(err)
	}
}
