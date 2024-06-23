package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/software78/rss-aggregator/internal/database"
)


func (apicfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Url string `json:"url"`
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to decode JSON: %v", err))
	}

	feed, error := apicfg.db.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Url: params.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		UserID: user.ID,
	})
	if error != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create feed: %v", error))
		return
	}
	respondWithJson(w, 200, databaseFeedToFeed(feed))
}

func (apicfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	
	feeds, error := apicfg.db.GetFeed(r.Context())
	if error != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to get feeds: %v", error))
		return
	}
	respondWithJson(w, 200, databaseFeedsToFeeds(feeds))
}