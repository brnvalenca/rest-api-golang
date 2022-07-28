package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/controllers"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {

	body := strings.NewReader(`{
		"id" : "1",
		"name" : "bruno",
		"email" : "brn@gmail.com",
		"password" : "321"
	}`)

	req, err := http.NewRequest("POST", "/users/create", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controllers.CreateUser)

	handler.ServeHTTP(rr, req)
	x := rr.Body
	fmt.Printf("%v", x)
	got := rr.Code
	want := http.StatusOK

	if got != want {
		t.Errorf("Handler returned wrong status code: got %v want %v", got, want)
	}
}
