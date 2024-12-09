package main

import (
	"log"
	"net/http"

	"Sistema-Streaming/pkg/auth"
	"Sistema-Streaming/pkg/catalog"
	"Sistema-Streaming/pkg/playback"
	"Sistema-Streaming/pkg/recommendation"
)

func main() {
	log.Println("Starting Sistema-Streaming server on port 8080...")

	// Definición de rutas
	http.HandleFunc("/auth", auth.AuthHandler)                               // Ruta para autenticación
	http.HandleFunc("/catalog", catalog.CatalogHandler)                      // Ruta para listar catálogo
	http.HandleFunc("/playback", playback.PlaybackHandler)                   // Ruta para reproducción
	http.HandleFunc("/recommendation", recommendation.RecommendationHandler) // Ruta para recomendaciones

	// Inicia el servidor HTTP en el puerto 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

