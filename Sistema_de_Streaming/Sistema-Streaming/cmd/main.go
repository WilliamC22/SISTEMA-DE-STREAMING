package main

import (
	"Sistema-Streaming/pkg/streaming"
	"log"
	"net/http"
)

func main() {
	// Inicializa el servidor del sistema de streaming
	log.Println("Starting streaming server on port 8080...")
	http.HandleFunc("/catalog", streaming.CatalogHandler)
	http.HandleFunc("/auth", streaming.AuthHandler)
	http.HandleFunc("/playback", streaming.PlaybackHandler)
	http.HandleFunc("/recommendation", streaming.RecommendationHandler)

	// Ejecuta el servidor en el puerto 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
