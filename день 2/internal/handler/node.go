package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"study/internal/model"
	"study/internal/repository"

	"github.com/gorilla/mux"
)

// дает кабель для подключения к бд
type NoteHandler struct {
	repo *repository.NoteRepository
}

// разрешение на подключению к этому кабелю
func NewNoteHandler(repo *repository.NoteRepository) *NoteHandler {
	return &NoteHandler{repo: repo}
}

// создание листа
func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var n model.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ctx := r.Context()
	if err := h.repo.Create(ctx, &n); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	FixN := FastProcessing(&n)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(FixN)
}

// показать всю таблицу
func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	Notes, err := h.repo.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Notes)

}

// найти по айди
func (h *NoteHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ctx := r.Context()
	n, err := h.repo.GetById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	FixN := FastProcessing(n)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(FixN)
}

// удаление по айди
func (h *NoteHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ctx := r.Context()
	if err := h.repo.DeleteById(ctx, id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	Notes, err := h.repo.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Notes)
}

// переписывает все за нас вручную
func FastProcessing(n *model.Note) *model.CreateNoteRequest {
	return &model.CreateNoteRequest{
		Id:      n.Id,
		Title:   n.Title,
		Content: n.Content,
	}
}
