package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	staking "github.com/oasislabs/oasis-core/go/staking/api"
)

// loadStakingClient loads staking client and returns it
func loadStakingClient(socket string) (*grpc.ClientConn, staking.Backend) {

	// Attempt to load connection with staking client
	connection, stakingClient, err := rpc.StakingClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to staking client : ", err)
		return nil, nil
	}
	return connection, stakingClient
}

// GetTotalSupply returns total supply at block height
func GetTotalSupply(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Using Oasis API to return total supply of tokens at specific block height
	totalSupply, err := so.TotalSupply(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get TotalSupply!"})
		lgr.Error.Println(
			"Request at /api/staking/totalsupply/ Failed to retrieve totalsupply : ", err)
		return
	}

	lgr.Info.Println("Request at /api/staking/totalsupply/ responding with TotalSupply!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: totalSupply})
}

// GetCommonPool returns common pool balance at block height
func GetCommonPool(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Return common pool at specific block height
	commonPool, err := so.CommonPool(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Common Pool!"})

		lgr.Error.Println(
			"Request at /api/staking/commonpool/ Failed to retrieve common pool : ", err)
		return
	}

	lgr.Info.Println("Request at /api/staking/commonpool/ responding with Common Pool!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: commonPool})
}

// GetStakingStateToGenesis returns state of genesis file of staking client
func GetStakingStateToGenesis(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Returning state to genesis at specific height
	genesisStaking, err := so.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Staking Genesis State!"})
		lgr.Error.Println(
			"Request at /api/staking/genesis/ Failed to retrieve Staking Genesis State : ", err)
		return
	}

	lgr.Info.Println(
		"Request at /api/staking/genesis/ responding with Staking Genesis State!")
	json.NewEncoder(w).Encode(responses.StakingGenesisResponse{
		GenesisStaking: genesisStaking})
}

// GetThreshold returns specific staking threshold by kind.
func GetThreshold(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Retrieving kind from query request
	recvKind := r.URL.Query().Get("kind")
	kind := checkKind(recvKind)
	if kind == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, kind needs to be string of int!"})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Create ThresholdQuery that will be sued to retrieved threshold amount
	query := staking.ThresholdQuery{Height: height, Kind: staking.ThresholdKind(kind)}

	// Return threshold from staking client using created query
	threshold, err := so.Threshold(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Threshold!"})
		lgr.Error.Println(
			"Request at /api/staking/threshold/ Failed to retrieve Threshold : ", err)
		return
	}

	// Responding with threshold quantity retrieved
	lgr.Info.Println(
		"Request at /api/staking/threshold/ responding with Threshold!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: threshold})
}

// GetAccounts returns IDs of all accounts with non-zero general balance
func GetAccounts(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Return accounts from staking client
	accounts, err := so.Accounts(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Accounts!"})
		lgr.Error.Println(
			"Request at /api/staking/accounts/ Failed to retrieve Accounts : ", err)
		return
	}

	// Respond with array of all accounts
	lgr.Info.Println("Request at /api/staking/accounts/ responding with Accounts!")
	json.NewEncoder(w).Encode(responses.AllAccountsResponse{AllAccounts: accounts})
}

// GetAccountInfo returns IDs of all accounts with non-zero general.
func GetAccountInfo(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that public key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	if len(ownerKey) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println(
			"Request at /api/staking/accountinfo/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "ownerKey can't be empty!"})
		return
	}

	// Unmarshall text into public key object
	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Create an owner query to be able to retrieve data with regards to account
	query := staking.OwnerQuery{Height: height, Owner: pubKey}

	// Retrieve account information using created query
	account, err := so.AccountInfo(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Account!"})
		lgr.Error.Println(
			"Request at /api/staking/accountinfo/ Failed to retrieve Account Info : ", err)
		return
	}

	// Return account information for created query
	lgr.Info.Println("Request at /api/staking/accountinfo/ responding with Account!")
	json.NewEncoder(w).Encode(responses.AccountResponse{AccountInfo: account})
}

// GetDelegations returns list of delegations for given owner
func GetDelegations(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that public key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	if len(ownerKey) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println(
			"Request at /api/staking/delegations/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "ownerKey can't be empty!"})
		return
	}

	// Unmarshal text into public key object
	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Create an owner query to be able to retrieve data with regards to account
	query := staking.OwnerQuery{Height: height, Owner: pubKey}

	// Return delegations for given account query
	delegations, err := so.Delegations(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Delegations!"})

		lgr.Error.Println(
			"Request at /api/staking/delegations/ Failed to retrieve Delegations : ", err)
		return
	}

	// Respond with delegations for given account query
	lgr.Info.Println("Request at /api/staking/delegations/ responding with delegations!")
	json.NewEncoder(w).Encode(responses.DelegationsResponse{Delegations: delegations})
}

// GetDebondingDelegations returns list of debonding delegations for given owner (delegator).
func GetDebondingDelegations(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that public key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	if len(ownerKey) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println(
			"Request at /api/staking/accountinfo/ failed, ownerKey can't be empty!")

		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "ownerKey can't be empty!"})
		return
	}

	// Unmarshal text into public key object
	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Query created to retrieved Debonding Delegations for an account
	query := staking.OwnerQuery{Height: height, Owner: pubKey}

	// Retrieving debonding delegations for an account using above query
	debondingDelegations, err := so.DebondingDelegations(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Debonding Delegations!"})
		lgr.Error.Println(
			"Request at /api/staking/debondingdelegations/ Failed to retrieve Debonding Delegations : ", err)
		return
	}

	// Responding with debonding delegations for given accounts
	lgr.Info.Println(
		"Request at /api/staking/debondingdelegations/ responding with Debonding Delegations!")
	json.NewEncoder(w).Encode(responses.DebondingDelegationsResponse{
		DebondingDelegations: debondingDelegations})
}

// GetEvents 
func GetEvents(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load connection with staking client
	connection, so := loadStakingClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if so == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Return accounts from staking client
	events, err := so.GetEvents(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Events!"})
		lgr.Error.Println(
			"Request at /api/staking/events/ Failed to retrieve Events : ", err)
		return
	}


	// Respond with array of all accounts
	lgr.Info.Println("Request at /api/staking/events/ responding with" +
		"Events!")
	json.NewEncoder(w).Encode(responses.StakingEvents{StakingEvents: events})
}
