package catalog

import (
	"encoding/json"
	"net/http"
)

// Content representa un elemento del catálogo
type Content struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"` // Ej: "movie" o "series"
}

// Catálogo ficticio
var catalog = []Content{
	{ID: "1", Title: "Minions", Type: "movie"},
	{ID: "2", Title: "Nuestros Oceanos", Type: "series"},
}

// CatalogHandler maneja las solicitudes del catálogo
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catalog)
}
