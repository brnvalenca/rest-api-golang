package router

import "net/http"

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	UPDATE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

/*
 So here i specified the methods signatures for the Router that i'll be using in this API
 By doing this i can be able to decouple my external router from my logic business core,
 so in another source file i can specify these methods and define a structure that may
 implement this interface, with specific details of implementation but without make my
 core logic dependent to the external router framework.

*/
