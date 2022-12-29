package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"rest-api/golang/exercise/config"
	api_service "rest-api/golang/exercise/grpc-api"
	"rest-api/golang/exercise/middleware"
	grpc_api "rest-api/golang/exercise/proto/pb"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"
	"rest-api/golang/exercise/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// TODO: Estudar e pesquisar sobre o DOCKER

var (
	// Security Constants
	PasswordService security.IPasswordHash = security.NewMyHashPassword()
	// Breed Constants
	BreedRepo    repository.IBreedRepository = repository.NewBreedRepository()
	BreedService services.IBreedService      = services.NewBreedService(BreedRepo)
	// User Constants
	UserPrefsRepo repository.IPrefsRepository = repository.NewPrefsRepo()
	UserRepo      repository.IUserRepository  = repository.NewMySQLRepo()
	UserService   services.IUserService       = services.NewUserService(UserRepo, UserPrefsRepo)
	// Kennel Constants
	KennelAddrRepo repository.IAddressRepository = repository.NewAddrRepo()
	KennelRepo     repository.IKennelRepository  = repository.NewKennelRepository()
	KennelService  services.IKennelService       = services.NewKennelService(KennelRepo, KennelAddrRepo)
	// Dog Constants
	DogRepo    repository.IDogRepository = repository.NewSQL_D_Repo()
	DogService services.IDogService      = services.NewDogService(DogRepo, BreedRepo, KennelRepo)
	// Login Constant
	LoginService services.ILoginService = services.NewLoginService(PasswordService, UserService)
)

func main() {
	appConfig, err := config.New()
	if err != nil {
		log.Fatalln("error initialization config variables: %w", err)
	}
	utils.DB = utils.DBConn(appConfig)

	listener, err := net.Listen("tcp", appConfig.AppConfig.GrpcAddr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.Unary()),
	)

	dogService := api_service.NewDogService(DogService, BreedService)
	userService := api_service.NewUserGrpcService(UserService, PasswordService)
	kennelService := api_service.NewKennelGrpcService(KennelService, DogService)
	breedService := api_service.NewBreedGrpcService(BreedService)
	matchService := api_service.NewMatchGrpcService(UserService, DogService)
	loginService := api_service.NewLoginGrpcService(LoginService)

	grpc_api.RegisterDogServiceServer(grpcServer, dogService)
	grpc_api.RegisterUserServiceServer(grpcServer, userService)
	grpc_api.RegisterKennelServiceServer(grpcServer, kennelService)
	grpc_api.RegisterBreedServiceServer(grpcServer, breedService)
	grpc_api.RegisterMatchServiceServer(grpcServer, matchService)
	grpc_api.RegisterLoginServiceServer(grpcServer, loginService)

	reflection.Register(grpcServer)

	log.Printf("serving gRPC server at %s", listener.Addr().String())

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Println(err)
		}
	}()

	conn, err := grpc.Dial(appConfig.AppConfig.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	err = grpc_api.RegisterUserServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = grpc_api.RegisterDogServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}
	err = grpc_api.RegisterKennelServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}
	err = grpc_api.RegisterBreedServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}
	err = grpc_api.RegisterMatchServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}
	err = grpc_api.RegisterLoginServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}
	gwServer := &http.Server{
		Addr:    appConfig.AppConfig.HttpAddr,
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())

	defer utils.DB.Close()
}
