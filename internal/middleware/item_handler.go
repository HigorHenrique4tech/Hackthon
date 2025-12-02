package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	
	"github.com/go-chi/chi/v5"
	"hackathon-api/internal/models"
	"hackathon-api/internal/repository"
)
type ItemHandler struct {
	repo *repository.ItemRepository
}

func NewItemHandler() *ItemHandler {
	return &itemHandler{
		repo: repository.NewItemRepository(),
	}
}

func (h *ItemHandler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.create)
	r.Get("/", h.getAll)
	r.Get("/{id}", h.getByID)
	r.Put("/{id}", h.update)
	r.Delete("/{id}", h.delete)

	return r
}

func (h *ItemHandler) create(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	h.repo.Create(item)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) getAll(w http.ResponseWriter, r *http.Request) {
	items, _ := h.repo.GetAll()
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) getByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	item, err := h.repo.GetByID(id)
	if err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Item not found", 404)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	updated, err := h.repo.Update(id, item)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := h.repo.Delete(id)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	w.WriteHeader(204)
}
