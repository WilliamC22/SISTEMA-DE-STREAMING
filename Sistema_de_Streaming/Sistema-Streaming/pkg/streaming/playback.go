package streaming

import (
	"fmt"
	"net/http"
)

// PlaybackHandler maneja la reproducción de contenido
func PlaybackHandler(w http.ResponseWriter, r *http.Request) {
	contentID := r.URL.Query().Get("id")
	if contentID == "" {
		http.Error(w, "Content ID is required", http.StatusBadRequest)
		return
	}

	// Simular reproducción
	fmt.Fprintf(w, "Streaming content with ID: %s", contentID)
}
