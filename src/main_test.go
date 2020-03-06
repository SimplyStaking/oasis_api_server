package main

import (
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"
	"reflect"

	consensus_api "github.com/SimplyVC/oasis_api_server/src/consensus_api"
	registry_api "github.com/SimplyVC/oasis_api_server/src/registry_api"
	staking_api "github.com/SimplyVC/oasis_api_server/src/staking_api"
)

//Global Configuration files set
const (
	portFile = "../config/user_config_main_test.ini"
	socketFile = "../config/user_config_nodes_test.ini"

	portFileFail = "../config/user_config_main_test1.ini"
	socketFileFail = "../config/user_config_nodes_test1.ini"
)

//Testing the pinging of the API itself.
func TestPingApi(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/pingApi", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Pong)
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

//Testing the loading on the configuration file
func TestLoadConfig_Success(t *testing.T){
	portConf, socketConf := loadConfig(portFile , socketFile)
	if portConf == nil{
		t.Errorf("Failed to load Port file from path : %v",portFile)
	}else if socketConf == nil{
		t.Errorf("Failed to load Socket file from path : %v",socketFile)
	}
}

//Testing to fail the loading of the configuration file
func TestLoadConfig_Fail(t *testing.T){
	portConf, socketConf := loadConfig(portFileFail , socketFileFail)
	if portConf != nil{
		t.Errorf("Failed to not return a nil when loading the Port file from path : %v",portFileFail)
	}else if socketConf != nil{
		t.Errorf("Failed to not return a nil when loading the Socket file from path : %v",socketFileFail)
	}
}

//Testing the creation of the Oasis API Objects
func TestOasisObjects(t *testing.T){
	//Load the configuration and start create the necessary objects
	_, socketConf := loadConfig(portFile , socketFile)

	//Checks are made to make sure that the Structs aren't empty
	co, ro, so := loadOasisAPIs(socketConf)
	if true == reflect.DeepEqual(co, (consensus_api.ConsensusObject{})){
		t.Errorf("Failed returned a nil when getting Consensus Object from path : %v",socketFile)
	}
	if true == reflect.DeepEqual(ro, (registry_api.RegistryObject{})){
		t.Errorf("Failed returned a nil when getting Registry Object from path : %v",socketFile)
	}
	if true == reflect.DeepEqual(so, (staking_api.StakingObject{})){
		t.Errorf("Failed returned a nil when getting Staking Object from path : %v",socketFile)
	}
}