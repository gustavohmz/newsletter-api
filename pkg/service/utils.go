package service

import (
	"encoding/json"
	"net/http"
	"regexp"
)

// ErrorResponse estructura para respuestas de error
type ErrorResponse struct {
	Error string `json:"error"`
}

// Función de utilidad para responder con un error en formato JSON
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, ErrorResponse{Error: message})
}

// RespondWithJSON envía una respuesta JSON al cliente.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Función para validar el formato del correo electrónico
func IsValidEmail(email string) bool {
	// Utiliza una expresión regular simple para validar el formato del correo
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
