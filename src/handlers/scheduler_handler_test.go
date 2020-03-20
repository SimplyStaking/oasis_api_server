package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	scheduler_api "github.com/oasislabs/oasis-core/go/scheduler/api"
)

func Test_GetValidators_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetValidators", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetValidators)
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

func Test_GetValidators_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetValidators", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetValidators)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetValidators_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetValidators", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetValidators)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting 52 Validators at Block Height 3
	expected := 52

	// Responding with a Genesis File
	validators := &responses.ValidatorsResponse{
		Validators: []*scheduler_api.Validator{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), validators)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if len(validators.Validators) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(validators.Validators), expected)
	}
}

func Test_GetCommittees_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetCommittees", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetCommittees)
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

func Test_GetCommittees_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetCommittees", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetCommittees)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetSchedulerStateToGenesis_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetSchedulerStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetSchedulerStateToGenesis)
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

func Test_GetSchedulerStateToGenesis_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetSchedulerStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetValidators)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetSchedulerStateToGenesis_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetSchedulerStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetSchedulerStateToGenesis)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting 10 Minimum Validators at Block Height 3
	expected := 10

	// Responding with a Genesis File
	schedulerGenesis := &responses.SchedulerGenesisState{
		SchedulerGenesisState: &scheduler_api.Genesis{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), schedulerGenesis)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if schedulerGenesis.SchedulerGenesisState.Parameters.MinValidators != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			schedulerGenesis.SchedulerGenesisState.Parameters.MinValidators, expected)
	}
}
