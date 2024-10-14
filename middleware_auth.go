package main

import (
	"fmt"
	"net/http"
	"thatsbruno/blogo/internal/auth"
	"thatsbruno/blogo/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			repondWithError(w, 403, fmt.Sprintf("Couldnt grab auth information %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			repondWithError(w, 404, fmt.Sprintf("User not found %v", err))
		}

		handler(w, r, user)
	}
}
