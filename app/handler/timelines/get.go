package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Handle request for `GET /v1/timeline/public`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	limitStr := queryParams.Get("limit")
	limit, err := strconv.Atoi(limitStr)

	timelineDTO, err := h.timelineUsecase.Get(r.Context(), limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timelineDTO.Timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
