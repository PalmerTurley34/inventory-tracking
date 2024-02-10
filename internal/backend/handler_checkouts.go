package backend

import (
	"fmt"
	"net/http"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) checkOutItem(w http.ResponseWriter, r *http.Request, user db.User) {
	id, err := GetUrlUUID(r)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error with {ID} param: %v", err))
		return
	}
	item, err := cfg.DB.CheckOutItem(r.Context(), db.CheckOutItemParams{
		ID:     id,
		UserID: &user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("couldn't check out item: %v", err))
		return
	}
	cfg.DB.LogCheckOut(r.Context(), db.LogCheckOutParams{
		ID:              uuid.New(),
		InventoryItemID: id,
		UserID:          user.ID,
	})
	respondWithJSON(w, 200, item)
}

func (cfg *apiConfig) checkInItem(w http.ResponseWriter, r *http.Request, user db.User) {
	id, err := GetUrlUUID(r)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error with {ID} param: %v", err))
		return
	}
	item, err := cfg.DB.CheckInItem(r.Context(), id)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("couldn't check in item: %v", err))
		return
	}
	cfg.DB.LogCheckIn(r.Context(), db.LogCheckInParams{
		InventoryItemID: id,
		UserID:          user.ID,
	})
	respondWithJSON(w, 200, item)
}

func (cfg *apiConfig) getUserInventory(w http.ResponseWriter, r *http.Request, user db.User) {
	items, err := cfg.DB.GetUserInventory(r.Context(), &user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("couldn't get inventory: %v", err))
	}
	respondWithJSON(w, 200, items)
}
