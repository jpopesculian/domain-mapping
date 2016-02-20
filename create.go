package main

import (
	"net/http"
)

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	var form DomainCreateForm
	if err := ReadJsonForm(r, &form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	form = TrimDomainCreateForm(form)
	if ok, err := ValidateDomainCreateForm(form); err != nil || !ok {
		http.Error(w, "Invalid Form", http.StatusInternalServerError)
		return
	}
	domain, err := RepoCreateDomainMapping(form.Name, form.Hash, form.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, domain)
}
