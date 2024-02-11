package backend

import (
	"fmt"
	"net/http"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Error with ApiKey: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Error with DB: %v", err))
			return
		}
		handler(w, r, user)
	}
}
