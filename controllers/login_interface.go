package controllers

import "net/http"

type LoginInterface interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}
