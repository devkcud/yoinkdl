package main

import (
	"log"

	"github.com/devkcud/yoinkdl/internal/cli/app"
)

func main() {
	if err := app.New().Run(); err != nil {
		log.Fatal(err)
	}
}
