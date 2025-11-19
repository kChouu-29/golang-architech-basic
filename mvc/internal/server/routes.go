package server

import (
	"mvc/internal/handler"
	"net/http"
)

func NewRouter(userHandler *handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", userHandler.CreateUserHandler)
	mux.HandleFunc("GET /user/{id}", userHandler.GetUserByIdHandler)
	mux.HandleFunc("GET /users", userHandler.GetAllUserHandler)
	mux.HandleFunc("PUT /user/{id}", userHandler.UpdateUserByIdHandler)
	mux.HandleFunc("DELETE /user/{id}", userHandler.DeleteUserByIdHandler)
	return mux
}
