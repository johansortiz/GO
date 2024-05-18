// handlers.go

package handlers

import (
	"encoding/json"
	"net/http"

	. "github.com/johansortiz/SuperHero/repository"
)

type SuperheroHandler struct {
	repo *SuperheroRepository
}

func NewSuperheroHandler(repo *SuperheroRepository) *SuperheroHandler {
	return &SuperheroHandler{repo: repo}
}

func (h *SuperheroHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	heroName := r.URL.Query().Get("hero")
	if heroName == "" {
		http.Error(w, "missing 'hero' query parameter", http.StatusBadRequest)
		return
	}

	superhero, err := h.repo.Get(heroName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(superhero)
}
