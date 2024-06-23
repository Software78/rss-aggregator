package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/software78/rss-aggregator/internal/database"
)

func (apicfg *apiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to decode JSON: %v", err))
	}

	feedID, err := uuid.Parse(params.FeedID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to parse feed ID: %v", err))
		return
	}
	
	feedFollow, error := apicfg.db.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if error != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create feed follow: %v", error))
		return
	}
	respondWithJson(w, 200, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apicfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	
	feedFollows, error := apicfg.db.GetFeedFollowsByUserId(r.Context(), user.ID)
	if error != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to get feed follows: %v", error))
		return
	}
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}