package recommendation

import (
	"Sistema-Streaming/pkg/catalog"
	"encoding/json"
	"net/http"
)

// RecommendationHandler maneja las recomendaciones de contenido
func RecommendationHandler(w http.ResponseWriter, r *http.Request) {
	// Recomendaciones ficticias basadas en el catálogo
	recommendations := []catalog.Content{
		{ID: "3", Title: "Madagascar 2", Type: "movie"},
		{ID: "4", Title: "Peaky Blinders", Type: "series"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}
