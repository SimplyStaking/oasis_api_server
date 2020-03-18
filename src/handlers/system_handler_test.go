package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
)

func Test_GetMemory(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/system/GetMemory", nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetMemory)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "memory"
	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetDisk(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/system/GetDisk", nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDisk)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "disk"
	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetCPU(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/system/GetCPU", nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetCPU)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "cpu"
	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetNetwork(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/system/GetNetwork", nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNetwork)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "network"
	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
