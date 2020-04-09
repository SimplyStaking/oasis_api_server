package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	sentry_api "github.com/oasislabs/oasis-core/go/sentry/api"
)

func Test_GetSentryAddresses_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/sentry/addresses/", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetSentryAddresses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Sentry name requested doesn't exist"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetSentryAddresses_success(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/sentry/addresses/", nil)
	q := req.URL.Query()
	q.Add("name", "sentry_1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetSentryAddresses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Responding with Genesis File
	addresses := &responses.SentryResponse{
		SentryAddresses: &sentry_api.SentryAddresses{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), addresses)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	expected := "result"
	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
