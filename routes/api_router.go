package routes

import (
	"rest-api/golang/exercise/controllers"
	router "rest-api/golang/exercise/http"
	repository "rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/services"
)

var (
	port                      = ":8080"
	httpRouter router.IRouter = router.NewMuxRouter()

	/* User Variables */
	UserRepo       repository.IUserRepository = repository.NewMySQLRepo()
	UserService    services.IUserService      = services.NewUserService(UserRepo)
	UserController controllers.IController    = controllers.NewUserController(UserService)

	/* Kennel Variables */
	KennelRepo       repository.IKennelInterface = repository.NewKennelRepository()
	KennelService    services.IKennelService     = services.NewKennelService(KennelRepo)
	KennelController controllers.IController     = controllers.NewKennelController(KennelService)
)

var ()

func HandleKennelReq() {
	httpRouter.GET("/kennels/", KennelController.GetAll)
	httpRouter.GET("/kennels/{id}/", KennelController.GetById)
	httpRouter.POST("/kennels/create", KennelController.Create)
	httpRouter.DELETE("/kennels/delete/{id}/", KennelController.Delete)
	httpRouter.UPDATE("/kennels/update/{id}/", KennelController.Update)
}

func HandleAllReq() {
	HandleKennelReq()
	httpRouter.GET("/users", UserController.GetAll)
	httpRouter.GET("/users/{id}", UserController.GetById)
	httpRouter.POST("/users/create", UserController.Create)
	httpRouter.DELETE("/users/delete/{id}", UserController.Delete)
	httpRouter.UPDATE("/users/update/{id}", UserController.Update)
	httpRouter.SERVE(port)
}
