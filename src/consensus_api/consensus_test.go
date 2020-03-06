package consensus_api_test

import (
	"context"
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"os"

	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	consensus_api "github.com/SimplyVC/oasis_api_server/src/consensus_api"
)

var (
	ctx context.Context
	co  consensus_api.ConsensusObject
	ws_url = "unix:/home/vvol/serverdir/node/internal.sock"
	ws_url_Invalid = "unix:/home/vvol/serverdir/nodes/internal.sock"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	ctx = context.Background()
	co = consensus_api.ConsensusObject{}

	co.SetContext(ctx)

	_, consensusClient, err := rpc.ConsensusClient(ws_url)
	if err != nil {
		panic(err)
	}
	
	co.AddClient("Oasis_Local",consensusClient)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}
 
func teardown() {

	ctx = context.Background()
	co = consensus_api.ConsensusObject{}

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

//Testing if the setup indeed added one client
func TestConsensusObject_1(t *testing.T){
	expected := 1
	retrieved := len(co.GetClients())

	if retrieved != expected {
		t.Errorf("Unexpected number of Consensus Clients : got %v want %v",
		retrieved, expected)
	}
}

//Adding another client and testing that there are two clients in the map
func TestConsensusObject_2(t *testing.T){
	nodeName := "Oasis_Local_1"

	_, consensusClient, err := rpc.ConsensusClient(ws_url_Invalid)
	if err != nil {
		panic(err)
	}
	
	co.AddClient(nodeName,consensusClient)

	expected := 2
	retrieved := len(co.GetClients())

	//Check if the map contains two clients
	if retrieved != expected {
		t.Errorf("Unexpected number of Consensus Clients : got %v want %v",
		retrieved, expected)
	}
	clients := co.GetClients()

	if _, ok := clients[nodeName]; ok{
		
	}else{
		t.Errorf("Expected to find node with name %v but didn't find it.",
		nodeName)
	}
	
}

//Testing the pinging of a Node API
func TestPingNode_1(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/pingNode", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(co.Pong)
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

//Testing the pinging of a Node that is registered but the client didnt find a connection
func TestPingNode_2(t *testing.T) {
	
	nodeName := "Oasis_Local_Failure"

	_, consensusClient, err := rpc.ConsensusClient(ws_url_Invalid)
	if err != nil {
		panic(err)
	}
	
	co.AddClient(nodeName,consensusClient)


	req, err := http.NewRequest("GET", "/api/pingNode", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("name", nodeName)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(co.Pong)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"error":"No reply from node"}`
	//Compare the strings after trimming white spaces
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//Testing the pinging of a Node that is not registered
func TestPingNode_3(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/pingNode", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(co.Pong)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"result":"An API for Unicorn needs to be setup before it can be queried"}`
	//Compare the strings after trimming white spaces
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}