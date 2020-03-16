package handlers_test

import (
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"

	common_entity "github.com/oasislabs/oasis-core/go/common/entity"
	common_node "github.com/oasislabs/oasis-core/go/common/node"
	registry_api "github.com/oasislabs/oasis-core/go/registry/api"
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


func Test_GetNodes_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetNodes", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNodes)
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


func Test_GetNodes_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetNodes", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNodes)
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

func Test_GetNodes_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetNodes", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNodes)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting 52 Nodes to be found at Height 3
	expected := 52

	//Responding with all nodes
	allNodes := &responses.NodesResponse {
		Nodes :  []*common_node.Node{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), allNodes)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if len(allNodes.Nodes) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(allNodes.Nodes), expected)
	}
}

func Test_GetRuntimes_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetRuntimes", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRuntimes)
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


func Test_GetRuntimes_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetRuntimes", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRuntimes)
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

func Test_GetRuntimes_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetRuntimes", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRuntimes)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting 0 Runtimes to be found at Height 3
	expected := 0

	//Responding with all the Runtimes
	allRuntimes := &responses.RuntimesResponse {
		Runtimes : []*registry_api.Runtime{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), allRuntimes)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if len(allRuntimes.Runtimes) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(allRuntimes.Runtimes), expected)
	}
}

