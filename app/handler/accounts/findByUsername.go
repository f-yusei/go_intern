package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/accounts`
func (h *handler) FindByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	accountDTO, err := h.accountUsecase.FindByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accountDTO.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
