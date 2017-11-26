package handlers

import (
	"net/http"
)

type IHandler interface {
	Upload(w http.ResponseWriter, r *http.Request)
	Download(w http.ResponseWriter, r *http.Request)
	Remove(w http.ResponseWriter, r *http.Request)
}

