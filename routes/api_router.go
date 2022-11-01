package routes

import (
	"rest-api/golang/exercise/controllers"
	router "rest-api/golang/exercise/http"
	repository "rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"
)

var (
	port                      = ":8080"
	httpRouter router.IRouter = router.NewMuxRouter()

	/* Security Constants */
	PasswordHash security.IPasswordHash = security.NewMyHashPassword()
	/* Breed Constants */
	BreedRepo       repository.IBreedRepository = repository.NewBreedRepository()
	BreedServ       services.IBreedService      = services.NewBreedService(BreedRepo)
	BreedController controllers.IController     = controllers.NewBreedController(BreedServ)

	/* User Constants */
	UserPrefsRepo  repository.IPrefsRepository = repository.NewPrefsRepo()
	UserRepo       repository.IUserRepository  = repository.NewMySQLRepo()
	UserService    services.IUserService       = services.NewUserService(UserRepo, UserPrefsRepo)
	UserController controllers.IController     = controllers.NewUserController(UserService, PasswordHash)

	/* Kennel Constants */
	KennelAddrRepo   repository.IAddressRepository = repository.NewAddrRepo()
	KennelRepo       repository.IKennelRepository  = repository.NewKennelRepository()
	KennelService    services.IKennelService       = services.NewKennelService(KennelRepo, KennelAddrRepo)
	KennelController controllers.IController       = controllers.NewKennelController(KennelService)

	/* Dog Constants */
	DogRepo       repository.IDogRepository = repository.NewSQL_D_Repo()
	DogService    services.IDogService      = services.NewDogService(DogRepo, BreedRepo, KennelRepo)
	DogController controllers.IController   = controllers.NewDogController(DogService)

	/* Login Constants */
	LoginCtrl controllers.LoginInterface = controllers.NewLoginController(UserService, PasswordHash)
)

func HandleBreedReq() {
	httpRouter.GET("/breeds/", BreedController.GetAll)
	httpRouter.GET("/breed/{id}/", BreedController.GetById)
	httpRouter.POST("/breed/create/", BreedController.Create)
	httpRouter.DELETE("/breed/delete/{id}/", BreedController.Delete)
	httpRouter.UPDATE("/breed/update/", BreedController.Update)
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
	httpRouter.POST("/login/", LoginCtrl.SignIn)
	httpRouter.GET("/users", UserController.GetAll)
	httpRouter.GET("/users/{id}", UserController.GetById)
	httpRouter.POST("/users/create", UserController.Create)
	httpRouter.DELETE("/users/delete/{id}", UserController.Delete)
	httpRouter.UPDATE("/users/update/{id}", UserController.Update)
	httpRouter.SERVE(port)
}
