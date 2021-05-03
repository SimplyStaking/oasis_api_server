package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	staking "github.com/oasisprotocol/oasis-core/go/staking/api"
	common_signature "github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
)

// loadStakingClient loads staking client and returns it
func loadStakingClient(socket string) (*grpc.ClientConn, staking.Backend) {

	// Attempt to load connection with staking client
	connection, stakingClient, err := rpc.StakingClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to staking client : ",
			err)
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
	if !confirmation  {
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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
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
			"Request at /api/staking/totalsupply failed to retrieve "+
				"totalsupply : ", err)
		return
	}

	lgr.Info.Println("Request at /api/staking/totalsupply responding with " +
		"TotalSupply!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: totalSupply})
}

// GetCommonPool returns common pool balance at block height
func GetCommonPool(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
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
			"Request at /api/staking/commonpool failed to retrieve common "+
				"pool : ", err)
		return
	}

	lgr.Info.Println("Request at /api/staking/commonpool responding with " +
		"Common Pool!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: commonPool})
}


// GetLastBlockFees returns the collected fees for previous block.
func GetLastBlockFees(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
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

	// Return LastBlockFees at specific block height
	lastestBlockFees, err := so.LastBlockFees(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get last block fees!"})

		lgr.Error.Println(
			"Request at /api/staking/lastblockfees failed to retrieve " +
				"last block fees : ", err)
		return
	}

	lgr.Info.Println("Request at /api/staking/lastblockfees responding with" +
		" latest block fees!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: 
		lastestBlockFees})
}

// GetStakingStateToGenesis returns state of genesis file of staking client
func GetStakingStateToGenesis(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
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
			"Request at /api/staking/genesis failed to retrieve Staking "+
				"Genesis State : ", err)
		return
	}

	lgr.Info.Println(
		"Request at /api/staking/genesis responding with Staking " +
			"Genesis State!")
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
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
		return
	}

	// Retrieving kind from query request
	recvKind := r.URL.Query().Get("kind")
	kind := checkKind(recvKind)
	if kind == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, kind needs to be a string representing an int!"})
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
			"Request at /api/staking/threshold failed to retrieve "+
				"Threshold : ", err)
		return
	}

	// Responding with threshold quantity retrieved
	lgr.Info.Println(
		"Request at /api/staking/threshold responding with Threshold!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: threshold})
}

// GetAddresses returns IDs of all accounts with non-zero general balance
func GetAddresses(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
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

	// Return addresses from staking client
	addresses, err := so.Addresses(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Addresses!"})
		lgr.Error.Println(
			"Request at /api/staking/addresses failed to retrieve Addresses : ",
			err)
		return
	}

	// Respond with array of all accounts
	lgr.Info.Println("Request at /api/staking/addresses responding with " +
		"Addresses!")
	json.NewEncoder(w).Encode(responses.AllAddressesResponse{AllAddresses: 
		addresses})
}

// GetAddressFromPublicKey returns a staking address from a given public key
func GetAddressFromPublicKey(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Get the public key from the query
	var pubKey common_signature.PublicKey
	publicKey := r.URL.Query().Get("pubKey")
	if len(publicKey) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println(
			"Request at /api/staking/publickeytoaddress failed, pubKey " +
				"can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "pubKey can't be empty!"})
		return
	}

	// Unmarshall text into public key object
	err := pubKey.UnmarshalText([]byte(publicKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into PublicKey", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into PublicKey."})
		return
	}

	// Return the address corresponding to the public key
	address := staking.NewAddress(pubKey)

	// Respond with  the address of the public key
	lgr.Info.Println("Request at /api/staking/publickeytoaddress responding " +
		"with Address!")
	json.NewEncoder(w).Encode(responses.AddressResponse{
		Address: address})
}


// GetConsensusParameters returns the staking consensus parameters.
func GetConsensusParameters(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
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

	// Return the staking consensus parameters
	consensusParameters, err := so.ConsensusParameters(context.Background(), 
		height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Addresses!"})
		lgr.Error.Println(
			"Request at /api/staking/consensusparameters failed to retrieve " +
			"Addresses : ",err)
		return
	}

	// Respond with array of all accounts
	lgr.Info.Println("Request at /api/staking/consensusparameters responding " +
		"with Addresses!")
	json.NewEncoder(w).Encode(responses.ConsensusParametersResponse{
		ConsensusParameters: consensusParameters})
}

// GetAccount returns the account descriptor for the given account.
func GetAccount(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
		return
	}

	var address staking.Address
	addressQuery := r.URL.Query().Get("address")
	if len(addressQuery) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println(
			"Request at /api/staking/account failed, address can't be " +
				"empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "address can't be empty!"})
		return
	}

	// Unmarshall text into public key object
	err := address.UnmarshalText([]byte(addressQuery))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Address", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Address."})
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
	query := staking.OwnerQuery{Height: height, Owner: address}

	// Retrieve account information using created query
	account, err := so.Account(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Account!"})
		lgr.Error.Println(
			"Request at /api/staking/account failed to retrieve Account: "+
				"", err)
		return
	}

	// Return account information for created query
	lgr.Info.Println("Request at /api/staking/account responding with " +
		"Account!")
	json.NewEncoder(w).Encode(responses.AccountResponse{Account: account})
}

// GetDelegations returns list of delegations for given owner
func GetDelegations(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
		return
	}

	var address staking.Address
	addressQuery := r.URL.Query().Get("address")
	if len(addressQuery) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println(
			"Request at /api/staking/delegations failed, address can't be " +
				"empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "address can't be empty!"})
		return
	}

	// Unmarshal text into public key object
	err := address.UnmarshalText([]byte(addressQuery))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Address", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Address."})
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
	query := staking.OwnerQuery{Height: height, Owner: address}

	// Return delegations for given account query
	delegations, err := so.DelegationsTo(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Delegations!"})

		lgr.Error.Println(
			"Request at /api/staking/delegations failed to retrieve "+
				"Delegations : ", err)
		return
	}

	// Respond with delegations for given account query
	lgr.Info.Println("Request at /api/staking/delegations responding with " +
		"delegations!")
	json.NewEncoder(w).Encode(responses.DelegationsResponse{Delegations:
		delegations})
}

// GetDebondingDelegations returns list of debonding delegations
// for given owner (delegator).
func GetDebondingDelegations(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {
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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
		return
	}

	var address staking.Address
	addressQuery := r.URL.Query().Get("address")
	if len(addressQuery) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println(
			"Request at /api/staking/account failed, address can't be " +
				"empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "address can't be empty!"})
		return
	}

	// Unmarshal text into public key object
	err := address.UnmarshalText([]byte(addressQuery))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Address", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Address."})
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
	query := staking.OwnerQuery{Height: height, Owner: address}

	// Retrieving debonding delegations for an account using above query
	debondingDelegations, err := so.DebondingDelegationsTo(context.Background(),
		&query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Debonding Delegations!"})
		lgr.Error.Println(
			"Request at /api/staking/debondingdelegations failed to retrieve"+
				" Debonding Delegations : ", err)
		return
	}

	// Responding with debonding delegations for given accounts
	lgr.Info.Println(
		"Request at /api/staking/debondingdelegations responding with " +
			"Debonding Delegations!")
	json.NewEncoder(w).Encode(responses.DebondingDelegationsResponse{
		DebondingDelegations: debondingDelegations})
}

// GetEvents returns events at a specific height.
func GetEvents(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation  {

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
			Error: "Unexpected value found, height needs to be a string representing an int!"})
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
			"Request at /api/staking/events failed to retrieve Events : ", err)
		return
	}

	// Respond with array of all accounts
	lgr.Info.Println("Request at /api/staking/events responding with" +
		" Events!")
	json.NewEncoder(w).Encode(responses.StakingEvents{StakingEvents: events})
}
