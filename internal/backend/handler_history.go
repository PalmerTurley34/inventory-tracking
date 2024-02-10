package backend

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) getItemHistory(w http.ResponseWriter, r *http.Request) {
	id, err := GetUrlUUID(r)
	if err != nil {
		respondWithError(w, 400, "{ID} is not a valid uuid")
		return
	}
	historyItems, err := cfg.DB.GetItemHistory(r.Context(), id)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("couldn't get item history: %v", err))
		return
	}
	respondWithJSON(w, 200, historyItems)
}
