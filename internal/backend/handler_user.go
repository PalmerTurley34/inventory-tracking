package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't decode request body: %v", err))
		return
	}

	if len(params.Password) < 8 {
		respondWithError(w, 400, "password must be at least 8 characters!")
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 0)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't hash password: %v", err))
		return
	}

	newUser, err := cfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Username:  params.Username,
		Password:  string(hashedPass),
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_username_key\"") {
			respondWithError(w, 400, fmt.Sprintf("Username: %v already exists. Please choose a unique username.", params.Username))
			return
		}
		respondWithError(w, 500, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	respondWithJSON(w, 201, newUser)
}

func (cfg *apiConfig) validateUsername(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't decode request body: %v", err))
		return
	}
	count, err := cfg.DB.CheckUsernameExists(r.Context(), params.Username)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Couldn't get user count: %v", err))
		return
	}
	type response struct {
		Valid bool `json:"valid"`
	}
	respondWithJSON(w, 200, response{Valid: !(count > 0)})
}
