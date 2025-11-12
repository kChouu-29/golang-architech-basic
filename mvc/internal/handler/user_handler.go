package handler

import (
	"encoding/json"
	"mvc/internal/controller"
	"mvc/internal/dto"
	"net/http"
)

type UserHandler struct {
	Ctrl *controller.UserController
}

func NewUserHandler(r *controller.UserController) *UserHandler {
	return &UserHandler{
		Ctrl: r,
	}
}
func (h *UserHandler) WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.WriteJson(w, http.StatusBadRequest, dto.CreateUserReponse{
			Message: "Invalid request body" + err.Error(),
			Data:    nil,
		})
		return
	}
	userRespone, err := h.Ctrl.CreateUserController(req)
	if err != nil {
		h.WriteJson(w, http.StatusInternalServerError, dto.CreateUserReponse{
			Message: "Failed to create user: " + err.Error(), // Trả về lỗi
			Data:    nil,
		})
		return
	}
	apiResponse := dto.CreateUserReponse{
		Message: "Create pesonal successfully",
		Data:    userRespone,
	}
	h.WriteJson(w, http.StatusOK, apiResponse)
}
