package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/software78/rss-aggregator/internal/database"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to decode JSON: %v", err))
	}

	user,error := apicfg.db.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		FirstName: params.FirstName,
		LastName: params.LastName,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if error != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create user: %v", error))
		return
	}


	respondWithJson(w, 200, user)
}