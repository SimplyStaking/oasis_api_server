package main

import (
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"
)

//Testing the pinging of the API itself.
func TestPingApi(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/pingApi", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Pong)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"result":"pong"}`
	//Compare the strings after trimming white spaces
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}