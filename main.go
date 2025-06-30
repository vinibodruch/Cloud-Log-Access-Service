package main

import (
	"Cloud-Log-Access-Service/app"
	"log"
)

func main() {
	application := app.NewApplication()

	if err := application.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
