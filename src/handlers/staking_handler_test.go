package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	common_quantity "github.com/oasisprotocol/oasis-core/go/common/quantity"
	staking_api "github.com/oasisprotocol/oasis-core/go/staking/api"
)

func Test_GetTotalSupply_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/totalsupply", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetTotalSupply)
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

func Test_GetTotalSupply_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/totalsupply", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetTotalSupply)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetTotalSupply_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/totalsupply", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetTotalSupply)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	quantity := &responses.QuantityResponse{
		Quantity: &common_quantity.Quantity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), quantity)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetCommonPool_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/commonpool", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetCommonPool)
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

func Test_GetCommonPool_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/commonpool", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetCommonPool)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetCommonPool_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/commonpool", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetCommonPool)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	quantity := &responses.QuantityResponse{
		Quantity: &common_quantity.Quantity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), quantity)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetStakingStateToGenesis_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/genesis", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetStakingStateToGenesis)
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

func Test_GetStakingStateToGenesis_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/genesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetStakingStateToGenesis)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetStakingStateToGenesis_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/genesis", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetStakingStateToGenesis)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	stakingGenesis := &responses.StakingGenesisResponse{
		GenesisStaking: &staking_api.Genesis{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), stakingGenesis)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetThreshold_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/threshold", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetThreshold)
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

func Test_GetThreshold_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/threshold", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetThreshold)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetThreshold_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/threshold", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("kind", "1")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetThreshold)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	quantity := &responses.QuantityResponse{
		Quantity: &common_quantity.Quantity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), quantity)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetAddresses_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/addresses", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAddresses)
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

func Test_GetAddresses_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/addresses", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAddresses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetAddresses_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/addresses", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAddresses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	AllAddresses := &responses.AllAddressesResponse{
		AllAddresses: []staking_api.Address{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), AllAddresses)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetAccount_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/account", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccount)
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

func Test_GetAccount_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/account", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccount)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetAccount_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/account", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("address", "oasis1qqqf342r78nz05dq2pa3wzh0w54k3ea49u6rqdhv")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccount)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	account := &responses.AccountResponse{
		Account: &staking_api.Account{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), account)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetDelegations_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/delegations", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDelegations)
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

func Test_GetDelegations_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/delegations", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDelegations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetDelegations_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/delegations", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("address", "oasis1qqqf342r78nz05dq2pa3wzh0w54k3ea49u6rqdhv")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDelegations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	delegations := &responses.DelegationsResponse{
		Delegations: map[staking_api.Address]*staking_api.Delegation{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), delegations)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetDebondingDelegations_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/debondingdelegations", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDebondingDelegations)
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

func Test_GetDebondingDelegations_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/debondingdelegations", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDebondingDelegations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetDebondingDelegations_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/debondingdelegations", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("address", "oasis1qqqf342r78nz05dq2pa3wzh0w54k3ea49u6rqdhv")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDebondingDelegations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	debondingDelegations := &responses.DebondingDelegationsResponse{
		DebondingDelegations: map[staking_api.Address][]*staking_api.DebondingDelegation{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), debondingDelegations)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func Test_GetEvents_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/events", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEvents)
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

func Test_GetEvents_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/events", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEvents)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Unexpected value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetEvents_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/staking/events", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetEvents)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "result"

	events := &responses.StakingEvents{
		StakingEvents: []staking_api.Event{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), events)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if strings.Contains(strings.TrimSpace(rr.Body.String()), expected) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}
