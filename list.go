package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetDomainsByUserId(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]
	domain, err := RepoGetDomainsByUserId(userId)
	if err != nil {
		http.Error(w, "Domains Not Found!", http.StatusNotFound)
		return
	}
	WriteJson(w, domain)
}
