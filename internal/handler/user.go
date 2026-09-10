package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"study/internal/model"
	"study/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Error", 400)
		return
	}
	ctx := r.Context()
	err := h.service.Create(ctx, &u)
	if errors.Is(err, service.ErrName) || errors.Is(err, service.ErrAge) {
		http.Error(w, err.Error(), 400)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)

}
