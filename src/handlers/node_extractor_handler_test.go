package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
)

func Test_NodeExtractorQueryGauge(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/extractor/gauge", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("gauge", "node_nf_conntrack_entries")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.NodeExtractorQueryGauge)
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

func Test_NodeExtractorQueryCounter(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/extractor/counter", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("counter", "node_timex_pps_calibration_total")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.NodeExtractorQueryCounter)
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
