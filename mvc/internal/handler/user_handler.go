package handler

import (
	"encoding/json"
	"mvc/internal/controller"
	"mvc/internal/dto"
	"net/http"
	"strconv"
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

func (h *UserHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Lấy ID từ URL
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.WriteJson(w, http.StatusBadRequest, dto.GetUserReponse{
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}
	userResponse, err := h.Ctrl.GetUserByIDController(id)
	if err != nil {
		h.WriteJson(w, http.StatusInternalServerError, dto.GetUserReponse{
			Message: "Failed to get user: " + err.Error(),
			Data:    nil,
		})
		return
	}
	apiResponse := dto.GetUserReponse{
		Message: "Get user successfully",
		Data:    userResponse,
	}
	h.WriteJson(w, http.StatusOK, apiResponse)
}

func (h *UserHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Gọi xuống Controller để lấy dữ liệu
	userRes, err := h.Ctrl.GetAllUserController()
	if err != nil {
		// 2. Nếu có lỗi, trả về lỗi (nên dùng InternalServerError)
		h.WriteJson(w, http.StatusInternalServerError, dto.GetAllUserReponse{
			// Lưu ý: Struct DTO của bạn dùng 'Massage' (viết sai)
			Message: "Failed to get all users: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 3. Nếu thành công, tạo response thành công
	apiResponse := dto.GetAllUserReponse{
		Message: "Get all users successfully",
		Data:    userRes,
	}

	// 4. Trả về JSON với status OK
	h.WriteJson(w, http.StatusOK, apiResponse)
}
func (h *UserHandler) UpdateUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Lấy ID từ URL
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.WriteJson(w, http.StatusBadRequest, dto.UpdateUseReponse{
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}
	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.WriteJson(w, http.StatusBadRequest, dto.UpdateUseReponse{
			Message: "Invalid request body" + err.Error(),
			Data:    nil,
		})
		return
	}
	userRespone, err := h.Ctrl.UpdateUserByIDController(req, id)
	if err != nil {
		h.WriteJson(w, http.StatusInternalServerError, dto.UpdateUseReponse{
			Message: "Failed to update user: " + err.Error(), // Trả về lỗi
			Data:    nil,
		})
		return
	}
	apiResponse := dto.UpdateUseReponse{
		Message: "Update user successfully",
		Data:    userRespone,
	}
	h.WriteJson(w, http.StatusOK, apiResponse)
}
func (h *UserHandler) DeleteUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Lấy ID từ URL
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.WriteJson(w, http.StatusBadRequest, dto.DeleteReponse{
			Message: "Invalid user ID",
			Dele:    false,
		})
		return
	}
	err = h.Ctrl.DeleteUserByIDController(id)
	if err != nil {
		h.WriteJson(w, http.StatusInternalServerError, dto.DeleteReponse{
			Message: "Failed to delete user: " + err.Error(),
			Dele:    false,
		})
		return
	}
	apiResponse := dto.DeleteReponse{
		Message: "Delete user successfully",
		Dele:    true,
	}
	h.WriteJson(w, http.StatusOK, apiResponse)
}
