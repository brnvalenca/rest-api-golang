package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() IRouter {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) GETBYID(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Delete(uri, f)
}

func (*chiRouter) UPDATE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Put(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v\n", port)
	http.ListenAndServe(port, chiDispatcher)
}
