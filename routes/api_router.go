package routes

import (
	"rest-api/golang/exercise/controllers"
	router "rest-api/golang/exercise/http"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/services"
)

var (
	port                                      = ":8080"
	httpRouter     router.IRouter             = router.NewMuxRouter()
	UserRepo       repository.IUserRepository = repository.NewMySQLRepo()
	UserService    services.IUserService      = services.NewUserService(UserRepo)
	UserController controllers.IController    = controllers.NewUserController(UserService)
)

func HandleAllReq() {
	httpRouter.GET("/users", UserController.GetAll)
	httpRouter.GET("/users/{id}", UserController.GetById)
	httpRouter.POST("/users/create", UserController.Create)
	httpRouter.DELETE("/users/delete/{id}", UserController.Delete)
	httpRouter.UPDATE("/users/update/{id}", UserController.Update)
	httpRouter.SERVE(port)

}
