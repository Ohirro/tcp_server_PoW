package main

import (
	"log"

	"word-of-wisdom-server/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
