package apiservice

import (
	"context"
	"rest-api/golang/exercise/proto/pb"
	"rest-api/golang/exercise/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginService struct {
	pb.UnimplementedLoginServiceServer
	loginServ services.ILoginService
}

func NewLoginGrpcService(loginServ services.ILoginService) *LoginService {
	return &LoginService{loginServ: loginServ}
}

func (loginServ *LoginService) SignIn(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	// TODO: Criar um service para Log in, e esse service vai chamar todas essas 3 funcoes. O controller deve só chamar o service.
	token, err := loginServ.loginServ.AuthenticateUser(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during user auth")
	}
	resp := &pb.LoginResponse{
		Token: token,
	}
	return resp, nil
}
