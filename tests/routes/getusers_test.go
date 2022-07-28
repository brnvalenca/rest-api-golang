package tests

import (
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/controllers"
	"testing"
)

func TestGetUse(t *testing.T) {

	// Creates a HTTP GET Request
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Now creates a ResponseRecorder to satisfy the http.ResponseWriter parameter

	rr := httptest.NewRecorder()

	// HandlerFunc allows me to use ordinary functions as HTTP handlers if they have the correct signature

	handler := http.HandlerFunc(controllers.GetUsers)

	// Call a HTTP service on the handler with the correct arguments

	handler.ServeHTTP(rr, req)

	got := rr.Code

	want := http.StatusOK

	if got != want {
		t.Errorf("handler returned wrong status code: got %v want %v", got, want)
	}

}
