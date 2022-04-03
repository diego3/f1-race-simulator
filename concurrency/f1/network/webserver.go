package network

import (
	"log"
	"net/http"
)

func WebServer() {
	routes()

	log.Println("Starting web server on port 8080")

	go ListenPayloadChannel()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error trying to server", err)
	}
}
