package handlers

import (
	"context"
	"net/http"
	"encoding/json"
	"google.golang.org/grpc"

	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	staking "github.com/oasislabs/oasis-core/go/staking/api"
	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

//loadStakingClient loads the staking client and returns it
func loadStakingClient(socket string) (*grpc.ClientConn, staking.Backend){
	//Attempt to load a connection with the staking client
	connection, stakingClient, err := rpc.StakingClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the staking client : ", err)
		return nil, nil
	}
	return connection, stakingClient
}

//GetTotalSupply returns the total supply
func GetTotalSupply(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	totalSupply, err := so.TotalSupply(context.Background(), height)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get TotalSupply!"})
		lgr.Error.Println("Request at /api/GetTotalSupply/ Failed to retrieve the totalsupply : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetTotalSupply/ responding with the TotalSupply!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{totalSupply})
}

//GetCommonPool returns the common pool balance
func GetCommonPool(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	commonPool, err := so.CommonPool(context.Background(), height)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Common Pool!"})
		lgr.Error.Println("Request at /api/GetCommonPool/ Failed to retrieve the common pool : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetCommonPool/ responding with the Common Pool!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{commonPool})
}

//GetStakingStateToGenesis returns the common pool balance
func GetStakingStateToGenesis(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	genesisStaking, err := so.StateToGenesis(context.Background(), height)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Staking Genesis State!"})
		lgr.Error.Println("Request at /api/GetStakingStateToGenesis/ Failed to retrieve the Staking Genesis State : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetStakingStateToGenesis/ responding with the Staking Genesis State!")
	json.NewEncoder(w).Encode(responses.StakingGenesisResponse{genesisStaking})
}

//GetThreshold returns the specific staking threshold by kind.
func GetThreshold(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	recvKind := r.URL.Query().Get("kind")
	kind := checkKind(recvKind)
	if kind == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, kind needs to be string of int!"})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}
	//Create the ThresholdQuery
	query := staking.ThresholdQuery{height, staking.ThresholdKind(kind)}

	threshold, err := so.Threshold(context.Background(), &query)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Threshold!"})
		lgr.Error.Println("Request at /api/GetThreshold/ Failed to retrieve the Threshold : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetThreshold/ responding with the Threshold!")
	json.NewEncoder(w).Encode(responses.QuantityResponse{threshold})
}

//GetAccounts returns the IDs of all accounts with a non-zero general.
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}
	accounts, err := so.Accounts(context.Background(), height)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Accounts!"})
		lgr.Error.Println("Request at /api/GetAccounts/ Failed to retrieve Accounts : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetAccounts/ responding with Accounts!")
	json.NewEncoder(w).Encode(responses.AllAccountsResponse{accounts})
}

//GetAccountInfo returns the IDs of all accounts with a non-zero general.
func GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Note Make sure that the public key that is being sent is coded properly
	//Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	lgr.Warning.Println("Received : ", ownerKey)

	if len(ownerKey) == 0 {
		//Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetAccountInfo/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{"ownerKey can't be empty!"})
		return
	}

	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key",err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to UnmarshalText into Public Key."})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	query := staking.OwnerQuery{height, pubKey}

	account, err := so.AccountInfo(context.Background(), &query)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Account!"})
		lgr.Error.Println("Request at /api/GetAccountInfo/ Failed to retrieve Account : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetAccountInfo/ responding with Account!")
	json.NewEncoder(w).Encode(responses.AccountResponse{account})
}

//GetDelegations returns the list of delegations for the given owner
func GetDelegations(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Note Make sure that the public key that is being sent is coded properly
	//Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	lgr.Warning.Println("Received : ", ownerKey)

	if len(ownerKey) == 0 {
		//Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetDelegations/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{"ownerKey can't be empty!"})
		return
	}

	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key",err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to UnmarshalText into Public Key."})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	query := staking.OwnerQuery{height, pubKey}

	delegations, err := so.Delegations(context.Background(), &query)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Delegations!"})
		lgr.Error.Println("Request at /api/GetDelegations/ Failed to retrieve Delegations : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetDelegations/ responding with delegations!")
	json.NewEncoder(w).Encode(responses.DelegationsResponse{delegations})
}

//GetDebondingDelegations returns the list of debonding delegations for the given owner (delegator).
func GetDebondingDelegations(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}
	
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Note Make sure that the public key that is being sent is coded properly
	//Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	ownerKey := r.URL.Query().Get("ownerKey")
	lgr.Warning.Println("Received : ", ownerKey)

	if len(ownerKey) == 0 {
		//Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetAccountInfo/ failed, ownerKey can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{"ownerKey can't be empty!"})
		return
	}

	err := pubKey.UnmarshalText([]byte(ownerKey))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key",err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to UnmarshalText into Public Key."})
		return
	}

	//Attempt to load a connection with the staking client
	connection, so := loadStakingClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if so == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	query := staking.OwnerQuery{height, pubKey}

	debondingDelegations, err := so.DebondingDelegations(context.Background(), &query)
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Debonding Delegations!"})
		lgr.Error.Println("Request at /api/GetDebondingDelegations/ Failed to retrieve Debonding Delegations : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetDebondingDelegations/ responding with Debonding Delegations!")
	json.NewEncoder(w).Encode(responses.DebondingDelegationsResponse{debondingDelegations})
}