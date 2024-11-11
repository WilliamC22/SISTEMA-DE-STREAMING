package streaming

import (
	"encoding/json"
	"net/http"
)

// RecommendationHandler maneja las recomendaciones de contenido
func RecommendationHandler(w http.ResponseWriter, r *http.Request) {
	recommendations := []Content{
		{ID: "3", Title: "Comedy Nights", Type: "movie"},
		{ID: "4", Title: "Thrill Seekers", Type: "series"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}
