package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetDomainByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	domain, err := RepoGetDomainByName(name)
	if err != nil {
		http.Error(w, "Domain Not Found!", http.StatusNotFound)
		return
	}
	WriteJson(w, domain)
}
