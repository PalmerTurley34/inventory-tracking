package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) loginUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't decode request body: %v", err))
		return
	}
	user, err := cfg.DB.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't get user data: %v", err))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		respondWithError(w, 401, "incorrect password, try again")
		return
	}
	respondWithJSON(w, 200, user)
}
