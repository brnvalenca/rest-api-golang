package authentication

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//var mySigningKey = []byte("secret_key") // TODO: Melhorar esse trabalho com o mysigningkey. Variável de Ambiente.
var key = os.Getenv("KEY")
var mySigningKey = []byte(key)

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
		tokenStr := extractToken(r)
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

func extractToken(r *http.Request) string {
	//Bearer sfdhskjdhfskj
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}
