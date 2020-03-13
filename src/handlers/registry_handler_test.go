package handlers_test

import (
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"

	common_entity "github.com/oasislabs/oasis-core/go/common/entity"
	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
)

func Test_GetEntities_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetEntities", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEntities)
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


func Test_GetEntities_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetEntities", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEntities)
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

func Test_GetEntities_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetEntities", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEntities)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting 87 Entities to be found at Height 3
	expected := 87
	
	//Responding with a Genesis File
	allEntities := &responses.EntitiesResponse {
		Entities : []*common_entity.Entity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), allEntities)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if len(allEntities.Entities) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(allEntities.Entities), expected)
	}
}
