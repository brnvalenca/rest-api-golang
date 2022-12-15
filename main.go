package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	apiservice "rest-api/golang/exercise/grpc-api"
	"rest-api/golang/exercise/grpc-api/interceptor"
	apipb "rest-api/golang/exercise/proto/pb"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"
	"rest-api/golang/exercise/utils"

	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// TODO: Create the match service between user and dog

type config struct {
	GrpcAddr string `env:"GRPCADDR" envDefault:"localhost:9090"`
	HttpAddr string `env:"HTTPADDR" envDefault:"localhost:8080"`
}

var (
	// Security Constants
	PasswordHash security.IPasswordHash = security.NewMyHashPassword()
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
)

func main() {
	utils.DB = utils.DBConn()

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	listener, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
	)
	// TODO : Study better ways to return log errors (log  libs)
	dogService := apiservice.NewDogService(DogService, BreedService)
	userService := apiservice.NewUserGrpcService(UserService, PasswordHash)
	kennelService := apiservice.NewKennelGrpcService(KennelService, DogService)
	breedService := apiservice.NewBreedGrpcService(BreedService)
	matchService := apiservice.NewMatchGrpcService(UserService, DogService)

	apipb.RegisterDogServiceServer(grpcServer, dogService)
	apipb.RegisterUserServiceServer(grpcServer, userService)
	apipb.RegisterKennelServiceServer(grpcServer, kennelService)
	apipb.RegisterBreedServiceServer(grpcServer, breedService)
	apipb.RegisterMatchServiceServer(grpcServer, matchService)

	reflection.Register(grpcServer)

	log.Printf("serving gRPC server at %s", listener.Addr().String())

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Println(err)
		}
	}()

	conn, err := grpc.Dial(cfg.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	err = apipb.RegisterUserServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = apipb.RegisterDogServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = apipb.RegisterKennelServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = apipb.RegisterBreedServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = apipb.RegisterMatchServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    cfg.HttpAddr,
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())

	defer utils.DB.Close()
}
