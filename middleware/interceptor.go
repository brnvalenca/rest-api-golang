package middleware

import (
	"context"
	"rest-api/golang/exercise/authentication"

	"google.golang.org/grpc"
)

const signUpMethod = "/apiservice.UserService/SignUp"
const loginMethod = "/apiservice.LoginService/SignIn"

func Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if info.FullMethod == signUpMethod || info.FullMethod == loginMethod {
			return handler(ctx, req)
		} else {
			err := authentication.Authorize(ctx, info.FullMethod)
			if err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
	}
}
