package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	common_entity "github.com/oasislabs/oasis-core/go/common/entity"
	common_node "github.com/oasislabs/oasis-core/go/common/node"
	registry_api "github.com/oasislabs/oasis-core/go/registry/api"
)

func Test_GetEntities_BadNode(t *testing.T) {
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

func Test_GetEntities_InvalidHeight(t *testing.T) {
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

func Test_GetEntities_Height3(t *testing.T) {
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
	allEntities := &responses.EntitiesResponse{
		Entities: []*common_entity.Entity{},
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

func Test_GetNodes_BadNode(t *testing.T) {
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

func Test_GetNodes_InvalidHeight(t *testing.T) {
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

func Test_GetNodes_Height3(t *testing.T) {
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
	allNodes := &responses.NodesResponse{
		Nodes: []*common_node.Node{},
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

func Test_GetRuntimes_BadNode(t *testing.T) {
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

func Test_GetRuntimes_InvalidHeight(t *testing.T) {
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

func Test_GetRuntimes_Height3(t *testing.T) {
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
	allRuntimes := &responses.RuntimesResponse{
		Runtimes: []*registry_api.Runtime{},
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

func Test_GetRegistryStateToGenesis_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetRegistryStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRegistryStateToGenesis)
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

func Test_GetRegistryStateToGenesis_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetRegistryStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRegistryStateToGenesis)
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

func Test_GetRegistryStateToGenesis_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetRegistryStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRegistryStateToGenesis)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting 87 Entities to be found at Height 3 in Genesis State
	expected := 87

	//Responding with all the Runtimes
	registryGenesis := &responses.RegistryGenesisResponse{
		GenesisRegistry: &registry_api.Genesis{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), registryGenesis)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if len(registryGenesis.GenesisRegistry.Entities) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(registryGenesis.GenesisRegistry.Entities), expected)
	}
}

func Test_GetEntity_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetEntity", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEntity)
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

func Test_GetEntity_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetEntity", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEntity)
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

func Test_GetEntity_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetEntity", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("entity", "CVzqFIADD2Ed0khGBNf4Rvh7vSNtrL1ULTkWYQszDpc=")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEntity)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting the same entity ID to be retrieved
	expected := "CVzqFIADD2Ed0khGBNf4Rvh7vSNtrL1ULTkWYQszDpc="

	//Responding with all the Runtimes
	registryEntity := &responses.RegistryEntityResponse{
		Entity: &common_entity.Entity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), registryEntity)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if registryEntity.Entity.ID.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			registryEntity.Entity.ID.String(), expected)
	}
}

func Test_GetNode_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetNode", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNode)
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

func Test_GetNode_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetNode", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNode)
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

func Test_GetNode_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetNode", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("nodeID", "A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto=")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNode)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Expecting the same entity ID to be retrieved
	expected := "A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto="

	registryNode := &responses.RegistryNodeResponse{
		Node: &common_node.Node{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), registryNode)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if registryNode.Node.ID.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			registryNode.Node.ID.String(), expected)
	}
}

func Test_GetRuntime_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetRuntime", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRuntime)
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

func Test_GetRuntime_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetRuntime", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetRuntime)
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
