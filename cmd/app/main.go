package main

import (
	"golang-demo/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Application terminated with error: %v", err)
	}
}
