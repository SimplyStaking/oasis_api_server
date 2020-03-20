package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hdl "github.com/SimplyVC/oasis_api_server/src/handlers"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	common_quantity "github.com/oasislabs/oasis-core/go/common/quantity"
	staking_api "github.com/oasislabs/oasis-core/go/staking/api"
)

func Test_GetTotalSupply_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetTotalSupply", nil)
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
	req, _ := http.NewRequest("GET", "/api/GetTotalSupply", nil)
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

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetTotalSupply_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetTotalSupply", nil)
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

	// Expecting 10000000000000000000 Total Supply
	expected := "10000000000000000000"

	quantity := &responses.QuantityResponse{
		Quantity: &common_quantity.Quantity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), quantity)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if quantity.Quantity.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			quantity.Quantity.String(), expected)
	}
}

func Test_GetCommonPool_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetCommonPool", nil)
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
	req, _ := http.NewRequest("GET", "/api/GetCommonPool", nil)
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

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetCommonPool_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetCommonPool", nil)
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

	expected := "7999142984653504864"

	quantity := &responses.QuantityResponse{
		Quantity: &common_quantity.Quantity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), quantity)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if quantity.Quantity.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			quantity.Quantity.String(), expected)
	}
}

func Test_GetStakingStateToGenesis_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetStakingStateToGenesis", nil)
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
	req, _ := http.NewRequest("GET", "/api/GetStakingStateToGenesis", nil)
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

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetStakingStateToGenesis_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetStakingStateToGenesis", nil)
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

	expected := "7999142984653504864"

	stakingGenesis := &responses.StakingGenesisResponse{
		GenesisStaking: &staking_api.Genesis{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), stakingGenesis)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if stakingGenesis.GenesisStaking.CommonPool.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			stakingGenesis.GenesisStaking.CommonPool.String(), expected)
	}
}

func Test_GetThreshold_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetThreshold", nil)
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
	req, _ := http.NewRequest("GET", "/api/GetThreshold", nil)
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

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetThreshold_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetThreshold", nil)
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

	expected := "100000000000"

	quantity := &responses.QuantityResponse{
		Quantity: &common_quantity.Quantity{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), quantity)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if quantity.Quantity.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			quantity.Quantity.String(), expected)
	}
}

func Test_GetAccounts_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetAccounts", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccounts)
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

func Test_GetAccounts_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetAccounts", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccounts)
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

func Test_GetAccounts_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetAccounts", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccounts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := 260

	allAccounts := &responses.AllAccountsResponse{
		AllAccounts: []common_signature.PublicKey{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), allAccounts)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if len(allAccounts.AllAccounts) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(allAccounts.AllAccounts), expected)
	}
}

func Test_GetAccountInfo_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetAccountInfo", nil)
	q := req.URL.Query()
	q.Add("name", "Unicorn")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccountInfo)
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

func Test_GetAccountInfo_InvalidHeight(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetAccountInfo", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "Unicorn")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccountInfo)
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

func Test_GetAccountInfo_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetAccountInfo", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("ownerKey", "AbMv7E+H4MWxfvwzSEx/BmOOwwk11P3JnJVEVVKK/ZA=")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetAccountInfo)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "0"

	account := &responses.AccountResponse{
		AccountInfo: &staking_api.Account{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), account)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if account.AccountInfo.General.Balance.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			account.AccountInfo.General.Balance.String(), expected)
	}
}

func Test_GetDelegations_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetDelegations", nil)
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
	req, _ := http.NewRequest("GET", "/api/GetDelegations", nil)
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

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetDelegations_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetDelegations", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("ownerKey", "AbMv7E+H4MWxfvwzSEx/BmOOwwk11P3JnJVEVVKK/ZA=")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDelegations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := 0

	delegations := &responses.DelegationsResponse{
		Delegations: map[common_signature.PublicKey]*staking_api.Delegation{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), delegations)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}

	if len(delegations.Delegations) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(delegations.Delegations), expected)
	}
}

func Test_GetDebondingDelegations_BadNode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetDebondingDelegations", nil)
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
	req, _ := http.NewRequest("GET", "/api/GetDebondingDelegations", nil)
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

	expected := `{"error":"Unexepcted value found, height needs to be string of int!"}`

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func Test_GetDebondingDelegations_Height3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/GetDebondingDelegations", nil)
	q := req.URL.Query()
	q.Add("name", "Oasis_Local")
	q.Add("height", "3")
	q.Add("ownerKey", "AbMv7E+H4MWxfvwzSEx/BmOOwwk11P3JnJVEVVKK/ZA=")

	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetDebondingDelegations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := 0

	debondingDelegations := &responses.DebondingDelegationsResponse{
		DebondingDelegations: map[common_signature.PublicKey][]*staking_api.DebondingDelegation{},
	}

	err := json.Unmarshal([]byte(rr.Body.String()), debondingDelegations)
	if err != nil {
		t.Errorf("Failed to unmarshall data")
	}
	if len(debondingDelegations.DebondingDelegations) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			len(debondingDelegations.DebondingDelegations), expected)
	}
}
