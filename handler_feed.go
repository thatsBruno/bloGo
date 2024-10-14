package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thatsbruno/blogo/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		repondWithError(w, 500, fmt.Sprintf("Error parsing JSON %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		repondWithError(w, 400, fmt.Sprintf("Error creating feed %d", err))
	}

	respondWithJSON(w, 200, feed)
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		repondWithError(w, 400, fmt.Sprintf("Error creating feed %d", err))
	}

	respondWithJSON(w, 200, feeds)
}
