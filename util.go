package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadJsonForm(r *http.Request, form interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, form); err != nil {
		return err
	}
	return nil
}

func WriteJson(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(object)
}
