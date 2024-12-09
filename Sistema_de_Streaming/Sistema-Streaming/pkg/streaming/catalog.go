package catalog

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Content representa un elemento del catálogo
type Content struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"` // "movie" o "series"
}

// Catálogo ficticio
var catalog = []Content{
	{ID: "1", Title: "Minions", Type: "movie"},
	{ID: "2", Title: "Peaky Blinders", Type: "series"},
}

// GetContentByID busca un contenido por su ID
func GetContentByID(id string) (*Content, error) {
	for _, content := range catalog {
		if content.ID == id {
			return &content, nil
		}
	}
	return nil, errors.New("content not found")
}

// CatalogHandler maneja las solicitudes del catálogo
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catalog)
}
