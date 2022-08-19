package controllers

import "net/http"

type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

/*

	Controller interface will define all the methods that MUST be implemented by any
	controller that a define. This interface will be instaced in the api-router file inside
	the router folder and will be responsable to all the handle functions to deal with the
	endpoints of my API.

*/
