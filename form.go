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
	if len(form.Name) < 1 {
		return false, nil
	}
	if len(form.Hash) < 1 {
		return false, nil
	}
	if len(form.UserId) < 1 {
		return false, nil
	}
	return true, nil
}

func TrimDomainCreateForm(form DomainCreateForm) DomainCreateForm {
	form.UserId = strings.TrimSpace(form.UserId)
	form.Name = strings.TrimSpace(form.Name)
	form.Hash = strings.TrimSpace(form.Hash)
	return form
}
