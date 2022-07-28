package tests

import (
	"net/http"
	"net/http/httptest"
	"rest-api/golang/exercise/controllers"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/delete/3213", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controllers.DeleteUser)

	handler.ServeHTTP(rr, req)

	got := rr.Code
	want := http.StatusOK

	if got != want {
		t.Errorf("Handler returned wrong status code: got %v want %v", got, want)
	}
}
