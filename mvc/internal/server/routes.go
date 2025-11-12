package server

import (
	"mvc/internal/handler"
	"net/http"
)

func NewRouter(userHandler *handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", userHandler.CreateUserHandler)
	return mux
}
