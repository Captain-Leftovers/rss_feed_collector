package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Captain-Leftovers/rss_feed_collector/internal/auth"
	"github.com/Captain-Leftovers/rss_feed_collector/internal/database"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
    params := parameters{}

    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Could not decode parameters")
        return
    }

    if params.Name == "" {
        respondWithError(w, http.StatusInternalServerError, "Name is required")
        return
    }

    user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
        ID: uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Name: params.Name,
    })

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Could not create user")
        return
    }

    respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}


func (cfg apiConfig) handlerGetCurrentUserInfo(w http.ResponseWriter, r *http.Request){

   apiKey, err := auth.GetAPIKey(r.Header)
   
    if err != nil {
        respondWithError(w, http.StatusUnauthorized, "No authorization api key included in request")
        return
    }
  

    dbUser, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

    if err != nil {
        respondWithError(w, http.StatusNotFound, "User with the given ApiKey not found")
        return
    }



    

    respondWithJSON(w, http.StatusOK, databaseUserToUser(dbUser))
}