package main

import (
	"Sistema-Streaming/pkg/auth"
	"Sistema-Streaming/pkg/catalog"
	"Sistema-Streaming/pkg/playback"
	"Sistema-Streaming/pkg/recommendation"
	"log"
	"net/http"
)

func main() {
	// Inicializa el servidor del sistema de streaming
	log.Println("Starting Sistema-Streaming server on port 8080...")

	// Rutas del sistema
	http.HandleFunc("/auth", auth.AuthHandler)
	http.HandleFunc("/catalog", catalog.CatalogHandler)
	http.HandleFunc("/playback", playback.PlaybackHandler)
	http.HandleFunc("/recommendation", recommendation.RecommendationHandler)

	// Ejecuta el servidor en el puerto 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
