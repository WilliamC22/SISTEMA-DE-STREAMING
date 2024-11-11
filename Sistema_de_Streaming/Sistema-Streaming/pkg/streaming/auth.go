package streaming

import (
	"encoding/json"
	"net/http"
)

// User representa un usuario
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthHandler maneja el inicio de sesión
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		// Simular validación de usuario
		if user.Username == "admin" && user.Password == "password" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Login successful"))
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		return
	}
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
}
