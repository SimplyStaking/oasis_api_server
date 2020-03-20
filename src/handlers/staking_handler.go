package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	staking "github.com/oasislabs/oasis-core/go/staking/api"
)

// loadStakingClient loads the staking client and returns it
func loadStakingClient(socket string) (*grpc.ClientConn, staking.Backend) {
	// Attempt to load a connection with the staking client
	connection, stakingClient, err := rpc.StakingClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the staking client : ", err)
		return nil, nil
	}
	return connection, stakingClient
}

// GetTotalSupply returns the total supply at a block height
func GetTotalSupply(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Using the Oasis API to return the total supply of tokens at a specific block height
	totalSupply, err := so.TotalSupply(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get TotalSupply!"})
		lgr.Error.Println("Request at /api/GetTotalSupply/ Failed to retrieve the totalsupply : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetTotalSupply/ responding with the TotalSupply!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: totalSupply})
}

// GetCommonPool returns the common pool balance at a block height
func GetCommonPool(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Return the common pool at a specific block height
	commonPool, err := so.CommonPool(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Common Pool!"})
		lgr.Error.Println("Request at /api/GetCommonPool/ Failed to retrieve the common pool : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetCommonPool/ responding with the Common Pool!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: commonPool})
}

// GetStakingStateToGenesis returns the state of the genesis file of the staking client
func GetStakingStateToGenesis(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Returning the state to genesis at a specific height
	genesisStaking, err := so.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Staking Genesis State!"})
		lgr.Error.Println("Request at /api/GetStakingStateToGenesis/ Failed to retrieve the Staking Genesis State : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetStakingStateToGenesis/ responding with the Staking Genesis State!")
	json.NewEncoder(w).Encode(responses.StakingGenesisResponse{GenesisStaking: genesisStaking})
}

// GetThreshold returns the specific staking threshold by kind.
func GetThreshold(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Retrieving the kind from the query request
	recvKind := r.URL.Query().Get("kind")
	kind := checkKind(recvKind)
	if kind == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, kind needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Create the ThresholdQuery that will be sued to retrieved the threshold amount
	query := staking.ThresholdQuery{Height: height, Kind: staking.ThresholdKind(kind)}

	// Return the threshold from the staking client using the created query
	threshold, err := so.Threshold(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Threshold!"})
		lgr.Error.Println("Request at /api/GetThreshold/ Failed to retrieve the Threshold : ", err)
		return
	}

	// Responding with the threshold quantity retrieved
	lgr.Info.Println("Request at /api/GetThreshold/ responding with the Threshold!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{Quantity: threshold})
}

// GetAccounts returns the IDs of all accounts with a non-zero general balance
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Return the accounts from the staking client
	accounts, err := so.Accounts(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Accounts!"})
		lgr.Error.Println("Request at /api/GetAccounts/ Failed to retrieve Accounts : ", err)
		return
	}

	// Respond with the array of all the accounts
	lgr.Info.Println("Request at /api/GetAccounts/ responding with Accounts!")
	json.NewEncoder(w).Encode(responses.AllAccountsResponse{AllAccounts: accounts})
}

// GetAccountInfo returns the IDs of all accounts with a non-zero general.
func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that the public key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	if len(ownerKey) == 0 {
		// Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetAccountInfo/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "ownerKey can't be empty!"})
		return
	}

	// Unmarshall the text into a public key object
	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Create an owner query to be able to retrieve data with regards to the account
	query := staking.OwnerQuery{Height: height, Owner: pubKey}

	// Retrieve the account information using the created query
	account, err := so.AccountInfo(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Account!"})
		lgr.Error.Println("Request at /api/GetAccountInfo/ Failed to retrieve Account Info : ", err)
		return
	}

	// Return the account information for the created query
	lgr.Info.Println("Request at /api/GetAccountInfo/ responding with Account!")
	json.NewEncoder(w).Encode(responses.AccountResponse{AccountInfo: account})
}

// GetDelegations returns the list of delegations for the given owner
func GetDelegations(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that the public key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	if len(ownerKey) == 0 {
		// Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetDelegations/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "ownerKey can't be empty!"})
		return
	}

	// Unmarshal the text into a public key object
	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Create an owner query to be able to retrieve data with regards to the account
	query := staking.OwnerQuery{Height: height, Owner: pubKey}

	// Return the delegations for a given account query
	delegations, err := so.Delegations(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Delegations!"})
		lgr.Error.Println("Request at /api/GetDelegations/ Failed to retrieve Delegations : ", err)
		return
	}

	// Respond with the delegations for a given account query
	lgr.Info.Println("Request at /api/GetDelegations/ responding with delegations!")
	json.NewEncoder(w).Encode(responses.DelegationsResponse{Delegations: delegations})
}

// GetDebondingDelegations returns the list of debonding delegations for the given owner (delegator).
func GetDebondingDelegations(w http.ResponseWriter, r *http.Request) {
	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving the height from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that the public key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	if len(ownerKey) == 0 {
		// Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetAccountInfo/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "ownerKey can't be empty!"})
		return
	}

	// Unmarshal text into a public key object
	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if so == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Query created to retrieved the Debonding Delegations for an account
	query := staking.OwnerQuery{Height: height, Owner: pubKey}

	// Retrieving the debonding delegations for an account using the above query
	debondingDelegations, err := so.DebondingDelegations(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Debonding Delegations!"})
		lgr.Error.Println("Request at /api/GetDebondingDelegations/ Failed to retrieve Debonding Delegations : ", err)
		return
	}

	// Responding with the debonding delegations for a given accounts
	lgr.Info.Println("Request at /api/GetDebondingDelegations/ responding with Debonding Delegations!")
	json.NewEncoder(w).Encode(responses.DebondingDelegationsResponse{DebondingDelegations: debondingDelegations})
}
