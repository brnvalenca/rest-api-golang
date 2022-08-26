package router

import (
	"fmt"
	"net/http"

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

/* 
	By doing this, every return of the NewMuxRouter (just another constructor function) will be 
	implementing the Router interface. And this interface is the middle point in communication between
	my api-router.go inside the routes folder, and the mux-router implementation. This make me independent
	of the implementation, making much more easier to switch the router if i want.
*/

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) GETBYID(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (*muxRouter) UPDATE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v\n", port)
	http.ListenAndServe(port, muxDispatcher)
}
