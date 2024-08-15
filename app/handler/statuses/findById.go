package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/status/{id}`
func (h *handler) FindByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// 数値変換エラー時のメッセージ
		http.Error(w, "無効なID形式です。整数である必要があります。", http.StatusBadRequest)
		return
	}

	statusDTO, err := h.statusUsecase.FindById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statusDTO.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
