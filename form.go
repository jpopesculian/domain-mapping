package main

import (
	"strings"
)

type DomainCreateForm struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Hash   string `json:"hash"`
}

func ValidateDomainCreateForm(form DomainCreateForm) (bool, error) {
	return true, nil
}

func TrimDomainCreateForm(form DomainCreateForm) DomainCreateForm {
	form.UserId = strings.TrimSpace(form.UserId)
	form.Name = strings.TrimSpace(form.Name)
	form.Hash = strings.TrimSpace(form.Hash)
	return form
}
