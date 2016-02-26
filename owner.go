package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Owner struct {
	UserId string `json:"user_id"`
}

func GetUserIdByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	userId, err := RepoGetUserIdByName(name)
	if err != nil {
		http.Error(w, "Domain Not Found!", http.StatusNotFound)
		return
	}
	WriteJson(w, Owner{userId})
}
