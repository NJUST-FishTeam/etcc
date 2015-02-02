package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "GetAllServices",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: GetServicesHandler,
	},
	Route{
		Name:        "AddService",
		Method:      "POST",
		Pattern:     "/",
		HandlerFunc: NewServiceHandler,
	},
	Route{
		Name:        "GetAllConfigs",
		Method:      "GET",
		Pattern:     "/{service}",
		HandlerFunc: GetConfigsHandler,
	},
	Route{
		Name:        "GetConfig",
		Method:      "GET",
		Pattern:     "/{service}/{config}",
		HandlerFunc: GetConfigHandler,
	},
	Route{
		Name:        "PutConfig",
		Method:      "POST",
		Pattern:     "/{service}/{config}",
		HandlerFunc: PostConfigHandler,
	},
}
