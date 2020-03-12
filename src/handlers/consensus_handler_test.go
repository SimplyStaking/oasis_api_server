package handlers_test

import (
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"
	"os"
	"io/ioutil"
	"encoding/json"

	gen_api "github.com/oasislabs/oasis-core/go/genesis/api"
	epoch_api "github.com/oasislabs/oasis-core/go/epochtime/api"
	consensus_api "github.com/oasislabs/oasis-core/go/consensus/api"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
	conf "github.com/SimplyVC/oasis_api_server/src/config"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	mint_types "github.com/tendermint/tendermint/types"
)

//Setting data to test with, valid and invalid path locations
const (
	socketFileFail = "../config/user_config_nodes_test_fail.ini"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup(){
	//Set the Logger that will be used by the API through all the packages
	//And Load all the configuration that need to be used by the router
	os.Chdir("../")
	lgr.SetLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	conf.LoadPortConfiguration()
	conf.LoadSocketConfiguration()
}

func Test_GetConsensusStateToGenesis_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetConsensusStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetConsensusStateToGenesis)
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

func Test_GetConsensusStateToGenesis_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetConsensusStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetConsensusStateToGenesis)
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

func Test_GetConsensusStateToGenesis_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetConsensusStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetConsensusStateToGenesis)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "questnet-2020-03-05-1583427600"
	
	//Responding with a Genesis File
	geneisState := &responses.Response_consensusgenesis {
		GenJSON : &gen_api.Document{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), geneisState)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if geneisState.GenJSON.ChainID != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//By Sending a negative 2 value I'm making sure that the client can't retrieve a state
func Test_GetConsensusStateToGenesis_Heightn2(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetConsensusStateToGenesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "-2")
	
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetConsensusStateToGenesis)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Failed to get Genesis file of Block!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetEpoch_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetEpoch", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEpoch)
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

func Test_GetEpoch_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetEpoch", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEpoch)
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


func Test_GetEpoch_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetEpoch", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEpoch)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := 3492
	
	epochTime := &responses.Response_epoch{}

	err := json.Unmarshal([]byte(rr.Body.String()), epochTime)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if epochTime.Ep != epoch_api.EpochTime(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			epochTime.Ep, expected)
	}
}

func Test_GetBlock_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlock", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlock)
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

func Test_GetBlock_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlock", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlock)
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


func Test_GetBlock_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlock", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlock)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected int64
	expected =  3
	
	block := &responses.Response_block {
		Blk : &consensus_api.Block{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), block)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if block.Blk.Height != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			block.Blk.Height, expected)
	}
}


func Test_GetBlockHeader_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlockHeader", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlockHeader)
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

func Test_GetBlockHeader_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlockHeader", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlockHeader)
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

func Test_GetBlockHeader_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlockHeader", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlockHeader)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected int64
	expected =  3
	
	blockHeader := &responses.Response_blockHeader {
		BlkHeader : &mint_types.Header{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), blockHeader)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if blockHeader.BlkHeader.Height != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			blockHeader.BlkHeader.Height, expected)
	}
}

func Test_GetBlockLastCommit_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlockLastCommit", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlockLastCommit)
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

func Test_GetBlockLastCommit_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlockLastCommit", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlockLastCommit)
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

func Test_GetBlockLastCommit_Height3(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetBlockLastCommit", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetBlockLastCommit)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected int64
	expected =  2
	
	blkLastCommit := &responses.Response_blockLastCommit {
		BlkLastCommit : &mint_types.Commit{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), blkLastCommit)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	newCommitObject := mint_types.Commit(*blkLastCommit.BlkLastCommit)
	if newCommitObject.Height() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
		newCommitObject.Height(), expected)
	}
}

func Test_GetTransactions_BadNode(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetTransactions", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetTransactions)
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

func Test_GetTransactions_InvalidHeight(t *testing.T){
	req, _ := http.NewRequest("GET", "/api/GetTransactions", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetTransactions)
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