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

	/* Breed Variables */
	BreedRepo       repository.IBreedRepository = repos.NewBreedRepository()
	BreedServ       services.IBreedService      = services.NewBreedService(BreedRepo)
	BreedController controllers.IController     = controllers.NewBreedController(BreedServ)

	/* Dog Variables */
	DogRepo       repository.IDogRepository = repos.NewSQL_D_Repo()
	DogService    services.IDogService      = services.NewDogService(DogRepo, BreedRepo)
	DogController controllers.IController   = controllers.NewDogController(DogService)

	/* User Variables */
	UserPrefsRepo  repository.IPrefsRepository = repos.NewPrefsRepo()
	UserRepo       repository.IUserRepository  = repos.NewMySQLRepo()
	UserService    services.IUserService       = services.NewUserService(UserRepo, UserPrefsRepo)
	UserController controllers.IController     = controllers.NewUserController(UserService)

	/* Kennel Variables */
	KennelAddrRepo   repository.IAddressRepository = repos.NewAddrRepo()
	KennelRepo       repository.IKennelRepository  = repos.NewKennelRepository()
	KennelService    services.IKennelService       = services.NewKennelService(KennelRepo, KennelAddrRepo)
	KennelController controllers.IController       = controllers.NewKennelController(KennelService)
)

func HandleBreedReq() {
	httpRouter.GET("/breeds/", BreedController.GetAll)
	httpRouter.GET("/breed/{id}/", BreedController.GetById)
	httpRouter.POST("/breed/create/", BreedController.Create)
	httpRouter.DELETE("/breed/delete/{id}/", BreedController.Delete)
	httpRouter.UPDATE("/breed/update/{id}/", BreedController.Update)
}

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
	httpRouter.POST("/kennels/create/", KennelController.Create)
	httpRouter.DELETE("/kennels/delete/{id}/", KennelController.Delete)
	httpRouter.UPDATE("/kennels/update/{id}/", KennelController.Update)
}

func HandleAllReq() {
	HandleKennelReq()
	HandleDogReq()
	HandleBreedReq()
	httpRouter.GET("/users", UserController.GetAll)
	httpRouter.GET("/users/{id}", UserController.GetById)
	httpRouter.POST("/users/create", UserController.Create)
	httpRouter.DELETE("/users/delete/{id}", UserController.Delete)
	httpRouter.UPDATE("/users/update/{id}", UserController.Update)
	httpRouter.SERVE(port)
}
