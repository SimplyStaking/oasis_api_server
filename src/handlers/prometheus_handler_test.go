package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
)

func Test_PrometheusQueryGauge(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/prometheus/gauge", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("gauge", "go_goroutines")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.PrometheusQueryGauge)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"
	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_PrometheusQueryCounter(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/prometheus/counter", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("counter", "go_memstats_alloc_bytes_total")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.PrometheusQueryCounter)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"
	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
