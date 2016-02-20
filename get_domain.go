package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetDomainByHash(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]
	domain, err := RepoGetDomainByHash(hash)
	if err != nil {
		http.Error(w, "domain Not Found!", http.StatusNotFound)
		return
	}
	WriteJson(w, domain)
}

func GetDomainByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	domain, err := RepoGetDomainByName(name)
	if err != nil {
		http.Error(w, "domain Not Found!", http.StatusNotFound)
		return
	}
	WriteJson(w, domain)
}
