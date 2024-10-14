package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thatsbruno/blogo/internal/database"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		repondWithError(w, 500, fmt.Sprintf("Error parsing JSON %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		repondWithError(w, 400, fmt.Sprintf("Error creating feed %d", err))
	}

	respondWithJSON(w, 200, feed)
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		repondWithError(w, 400, fmt.Sprintf("Error creating feed follows %d", err))
	}

	respondWithJSON(w, 200, feeds)
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		repondWithError(w, 400, fmt.Sprintf("Couldnt parse feed follow ide %d", err))
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		repondWithError(w, 400, fmt.Sprintf("Couldnt delete feed follow ide %d", err))
	}

	respondWithJSON(w, 200, struct{}{})
}
