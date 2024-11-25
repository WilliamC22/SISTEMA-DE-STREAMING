package utils

import "log"

// LogInfo registra mensajes informativos
func LogInfo(message string) {
	log.Println("INFO:", message)
}

// LogError registra errores
func LogError(message string, err error) {
	log.Printf("ERROR: %s - %v", message, err)
}
