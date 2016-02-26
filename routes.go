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
		"ListDomainsByUserId",
		"GET",
		"/user/{user_id}",
		GetDomainsByUserId,
	},
	Route{
		"GetUserIdByName",
		"GET",
		"/{name}/user",
		GetUserIdByName,
	},
	Route{
		"GetDomainByName",
		"GET",
		"/{name}",
		GetDomainByName,
	},
	Route{
		"CreateDomain",
		"POST",
		"/create",
		CreateDomain,
	},
	Route{
		"DeleteDomainByName",
		"DELETE",
		"/{name}",
		DeleteDomainByName,
	},
}
