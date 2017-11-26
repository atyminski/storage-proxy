package handlers

import(
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string
}

func NewErrorMessage(message string) *ErrorMessage {
	return &ErrorMessage{
		Message : message,
	}
}

func (msg ErrorMessage) AsJsonResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(msg)
}