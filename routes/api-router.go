package routes

import (
	"fmt"
	"net/http"
	"rest-api/golang/exercise/controllers"
	router "rest-api/golang/exercise/http"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/services"
)

var (
	httpRouter     router.Router          = router.NewMuxRouter()
	userRepo       repository.Repository  = repository.NewMySQLRepo()
	userService    services.Service       = services.NewUserService(userRepo)
	userController controllers.Controller = controllers.NewUserController(userService)
)

const (
	port string = ":8080"
)

func HandleUserRequest() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Servers UP")
	})
	HandleDogRequests()
	httpRouter.GET("/users", userController.GetAll)
	httpRouter.GET("/users/{id}", userController.GetById)
	httpRouter.POST("/users/create", userController.Create)
	httpRouter.DELETE("/users/delete/{id}", userController.Delete)
	httpRouter.UPDATE("/users/update/{id}", userController.Update)
	httpRouter.SERVE(port)
}

func HandleDogRequests() {
	httpRouter.GET("/dogs", controllers.GetAll)
	httpRouter.GET("/dogs/{id}", controllers.GetById)
	httpRouter.POST("/dogs/create", controllers.Create)
	httpRouter.DELETE("/dogs/delete/{id}", controllers.Delete)
	httpRouter.UPDATE("/dogs/update/{id}", controllers.Update)

}
