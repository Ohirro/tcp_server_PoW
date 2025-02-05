package main

import (
	"log"

	"word-of-wisdom-client/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
