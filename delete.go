package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteDomainByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	domain := RepoDeleteDomainByName(name)
	WriteJson(w, domain)
}
