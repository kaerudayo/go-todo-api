package main

import "net/http"

type UserHandler interface {
	HandleUserGet(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserSignup(http.ResponseWriter, *http.Request, httprouter.Params)
}
