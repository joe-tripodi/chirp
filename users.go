package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// unmarshal the response
	// execute the db command
	// return json data from the DB
	type CreateUserRequest struct {
		Email string `json:"email"`
	}

	var createUserJson CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createUserJson)
	if err != nil {
		log.Println("unable to parse request:", err)
		w.WriteHeader(500)
		return
	}
	log.Println(createUserJson)

	user, err := apiCfg.db.CreateUser(r.Context(), createUserJson.Email)
	if err != nil {
		// I could check if the user already exists and return a 4XX error
		log.Println("Failed to create user in the database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("user created:", user)

	userResponse := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	respBody, err := json.Marshal(userResponse)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(respBody)

}
