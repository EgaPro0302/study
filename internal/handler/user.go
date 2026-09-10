package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"study/internal/model"
	"study/internal/repository"
	"study/internal/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// создает пользователя в бд
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

// показывает всех пользователей
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.service.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}

// находит по айдишнику (/users/{id})
func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	ctx := r.Context()
	user, err := h.service.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			http.Error(w, "user not found", 404)
			return
		}
		if errors.Is(err, service.ErrId) {
			http.Error(w, "invalid id", 400)
			return
		}
		http.Error(w, "internal error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
