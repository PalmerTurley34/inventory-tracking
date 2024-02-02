package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/PalmerTurley34/inventory-tracking/internal/models"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createInventoryItem(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse request body: %v", err))
		return
	}

	var dbDescription sql.NullString
	if params.Description == "" {
		dbDescription = sql.NullString{Valid: false}
	} else {
		dbDescription = sql.NullString{String: params.Description, Valid: true}
	}

	newItem, err := cfg.DB.CreateInventoryItem(r.Context(), database.CreateInventoryItemParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        params.Name,
		Description: dbDescription,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Couldn't create inventory item: %v", err))
		return
	}
	respondWithJSON(w, 201, models.DBInventoryItemToResponse(newItem))
}

func (cfg *apiConfig) getAllInventoryItems(w http.ResponseWriter, r *http.Request) {
	items, err := cfg.DB.GetAllInventoryItems(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("couldn't get inventory items: %v", err))
		return
	}
	respondWithJSON(w, 200, models.DBInventoryItemsToResponse(items))
}
