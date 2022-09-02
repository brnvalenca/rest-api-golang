package routes

import (
	"rest-api/golang/exercise/controllers"
	router "rest-api/golang/exercise/http"
	repository "rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/repository/repos"
	"rest-api/golang/exercise/services"
)

var (
	port                      = ":8080"
	httpRouter router.IRouter = router.NewMuxRouter()

	/* Dog Variables */

	DogRepo       repository.IDogRepository = repos.NewSQL_D_Repo()
	DogService    services.IDogService      = services.NewDogService(DogRepo)
	DogController controllers.IController   = controllers.NewDogController(DogService)

	/* User Variables */
	UserRepo       repository.IUserRepository = repos.NewMySQLRepo()
	UserService    services.IUserService      = services.NewUserService(UserRepo)
	UserController controllers.IController    = controllers.NewUserController(UserService)

	/* Kennel Variables */
	KennelRepo       repository.IKennelRepository = repos.NewKennelRepository()
	KennelService    services.IKennelService     = services.NewKennelService(KennelRepo)
	KennelController controllers.IController     = controllers.NewKennelController(KennelService)
)

func HandleDogReq() {
	httpRouter.GET("/dogs/", DogController.GetAll)
	httpRouter.GET("/dogs/{id}/", DogController.GetById)
	httpRouter.POST("/dogs/create/", DogController.Create)
	httpRouter.DELETE("/dogs/delete/{id}/", DogController.Delete)
	httpRouter.UPDATE("/dogs/update/{id}/", DogController.Update)
}

func HandleKennelReq() {
	httpRouter.GET("/kennels/", KennelController.GetAll)
	httpRouter.GET("/kennels/{id}/", KennelController.GetById)
	httpRouter.POST("/kennels/create", KennelController.Create)
	httpRouter.DELETE("/kennels/delete/{id}/", KennelController.Delete)
	httpRouter.UPDATE("/kennels/update/{id}/", KennelController.Update)
}

func HandleAllReq() {
	HandleKennelReq()
	HandleDogReq()
	httpRouter.GET("/users", UserController.GetAll)
	httpRouter.GET("/users/{id}", UserController.GetById)
	httpRouter.POST("/users/create", UserController.Create)
	httpRouter.DELETE("/users/delete/{id}", UserController.Delete)
	httpRouter.UPDATE("/users/update/{id}", UserController.Update)
	httpRouter.SERVE(port)
}
