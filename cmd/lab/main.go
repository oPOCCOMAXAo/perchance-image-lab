package main

import (
	"log"

	"github.com/opoccomaxao/perchance-image-lab/pkg/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
