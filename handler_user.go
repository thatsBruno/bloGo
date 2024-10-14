package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thatsbruno/blogo/internal/auth"
	"thatsbruno/blogo/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		repondWithError(w, 500, fmt.Sprintf("Error parsing JSON %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		repondWithError(w, 400, fmt.Sprintf("Error creating user %d", err))
	}

	respondWithJSON(w, 200, user)
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		repondWithError(w, 403, fmt.Sprintf("Couldnt grab auth information %v", err))
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		repondWithError(w, 404, fmt.Sprintf("User not found %v", err))
	}
	respondWithJSON(w, 200, user)
}
