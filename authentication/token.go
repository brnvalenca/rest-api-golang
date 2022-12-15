package authentication

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

var key = os.Getenv("KEY")
var mySigningKey = []byte(key)

func VerifyToken(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return mySigningKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func Authorize(ctx context.Context, method string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization is not provided")
	}

	accessToken := extractkToken(values[0])

	_, err := VerifyToken(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return nil

}

func extractkToken(token string) string {
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// ABOVE THE FUNCTIONS BASED ON REST HTTP API //

func GenerateJWT(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("something went wrong: %s", err.Error())
	}
	return tokenString, nil
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := extract(r)
		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error during token parsing")
				}
				return mySigningKey, nil
			})

			if err != nil {
				w.Write([]byte(err.Error()))
			}

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				handler.ServeHTTP(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func extract(r *http.Request) string {
	//Bearer sfdhskjdhfskj
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}
