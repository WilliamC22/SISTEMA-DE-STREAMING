package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

// User representa un usuario con campos encapsulados
type User struct {
	username string
	password string
}

// NewUser crea un nuevo usuario con validaciones
func NewUser(username, password string) (*User, error) {
	if len(username) < 3 {
		return nil, errors.New("username must be at least 3 characters")
	}
	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}
	return &User{username: username, password: password}, nil
}

// GetUsername devuelve el nombre de usuario
func (u *User) GetUsername() string {
	return u.username
}

// SetPassword permite cambiar la contraseña con validación
func (u *User) SetPassword(newPassword string) error {
	if len(newPassword) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	u.password = newPassword
	return nil
}

// Authenticate verifica si la contraseña es correcta
func (u *User) Authenticate(inputPassword string) bool {
	return u.password == inputPassword
}

// AuthHandler maneja la autenticación
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := NewUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !user.Authenticate(credentials.Password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

