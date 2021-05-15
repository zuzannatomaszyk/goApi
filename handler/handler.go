package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zuzannatomaszyk/goApi/db"
	"github.com/zuzannatomaszyk/goApi/models"
)

var dbInstance db.Database

func HandleRequests(db db.Database) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	dbInstance = db

	router.HandleFunc("/messages", getMessages).Methods("GET")
	router.HandleFunc("/message", addMessage).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")

	return router
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := dbInstance.GetMessages()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, messages)
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	message := &models.Message{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if message.Text == "" || message.User == "" {
		respondWithError(w, http.StatusBadRequest, "Missing required fields. Required fields: user, text")
		return
	}

	if err := dbInstance.AddMessage(message); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, message)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dbInstance.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}
