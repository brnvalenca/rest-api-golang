package tests

import (
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/controllers"
	"strings"
	"testing"
)

func TestUpdateUser(t *testing.T) {

	body := strings.NewReader(`{
		"name" : "bruno",
		"email" : "brn@gmail.com",
		"password" : "321"
	}`)

	req, err := http.NewRequest("PUT", "/users/update", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controllers.UpdateUser)

	handler.ServeHTTP(rr, req)

	got := rr.Code
	want := http.StatusOK

	if got != want {
		t.Errorf("Handler returned wrong status code: got %v want %v", got, want)
	}
}
