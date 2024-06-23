package main

import (
	"net/http"

	"github.com/software78/rss-aggregator/internal/auth"
	"github.com/software78/rss-aggregator/internal/database"
)

type authedHandler func (http.ResponseWriter,  *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, error := auth.GetApiKey(r.Header)
		if error != nil {
			respondWithError(w, 401, "Failed to get API key")
			return
		}
		user, error := apicfg.db.GetUserByApiKey(r.Context(), apiKey)
		if error != nil {
			respondWithError(w, 500, "Failed to get user")
			return
		}
		handler(w, r, user)
	}
}