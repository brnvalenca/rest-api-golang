package tests

import (
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/controllers"
	"testing"
)

func TestGetUsersById(t *testing.T) {
	req, err := http.NewRequest("GET", "users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controllers.GetUsersById)

	handler.ServeHTTP(rr, req)

	got := rr.Code
	want := http.StatusOK

	if got != want {
		t.Errorf("handler returned wrong status code: got %v want %v", got, want)
	}

}
