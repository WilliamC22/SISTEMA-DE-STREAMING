package playback

import (
	"fmt"
	"net/http"
)

// Player define el comportamiento común para diferentes tipos de contenido
type Player interface {
	Play() string
}

// Video representa un contenido de video
type Video struct {
	Title string
}

// Audio representa un contenido de audio
type Audio struct {
	Title string
}

// Play implementa el método de la interfaz Player para Video
func (v Video) Play() string {
	return fmt.Sprintf("Playing video: %s", v.Title)
}

// Play implementa el método de la interfaz Player para Audio
func (a Audio) Play() string {
	return fmt.Sprintf("Playing audio: %s", a.Title)
}

// PlaybackHandler maneja la reproducción de contenido
func PlaybackHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.URL.Query().Get("type")
	title := r.URL.Query().Get("title")

	var player Player

	switch contentType {
	case "video":
		player = Video{Title: title}
	case "audio":
		player = Audio{Title: title}
	default:
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(player.Play()))
}
