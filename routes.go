package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

var routes = []Route{
	Route{
		"GetDomainByHash",
		"GET",
		"/hash/{hash}",
		GetDomainByHash,
	},
	Route{
		"GetDomainByName",
		"GET",
		"/name/{name}",
		GetDomainByName,
	},
	Route{
		"CreateDomain",
		"POST",
		"/create",
		GetDomainByName,
	},
}
