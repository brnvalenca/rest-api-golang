package router

import (
	"fmt"
	"net/http"
	"rest-api/golang/exercise/authentication"

	"github.com/gorilla/mux"
)

type muxRouter struct{} // struct that will implement the Router interface

var (
	muxDispatcher = mux.NewRouter() // instance of the muxRouter
)

// Contructor function that will return a struct implementing the Router interface
func NewMuxRouter() IRouter {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, authentication.IsAuthorized(f)).Methods("GET")
}

func (*muxRouter) GETBYID(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	if uri == "/login/" || uri == "/users/create" {
		muxDispatcher.HandleFunc(uri, f).Methods("POST")
	} else {
		muxDispatcher.HandleFunc(uri, authentication.IsAuthorized(f)).Methods("POST")
	}
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, authentication.IsAuthorized(f)).Methods("DELETE")
}

func (*muxRouter) UPDATE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, authentication.IsAuthorized(f)).Methods("PUT")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v\n", port)
	http.ListenAndServe(port, muxDispatcher)
}
