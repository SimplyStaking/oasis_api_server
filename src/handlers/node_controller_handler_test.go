package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
)

func Test_GetIsSynced_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/synced", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetIsSynced)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Node name requested doesn't exist"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetIsSynced_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/synced", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetIsSynced)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedTrue := `{"result":true}`
	expectedFalse := `{"result":false}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expectedTrue) &&
		strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expectedFalse){
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}
