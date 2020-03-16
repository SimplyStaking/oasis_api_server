package handlers

import (
	"context"
	"net/http"
	"encoding/json"
	"google.golang.org/grpc"

	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	cbor "github.com/oasislabs/oasis-core/go/common/cbor"
	mint_api "github.com/oasislabs/oasis-core/go/consensus/tendermint/api"
)

//loadConsensusClient loads the consensus client and returns it
func loadConsensusClient(socket string) (*grpc.ClientConn, consensus.ClientBackend){
	//Attempt to load a connection with the consensus client
	connection, consensusClient, err := rpc.ConsensusClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the consensus client : ", err)
		return nil, nil
	}
	return connection, consensusClient
}

// GetConsensusStateToGenesis returns the genesis state at the specified block height for Consensus.
func GetConsensusStateToGenesis(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if co == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	consensusGenesis, err := co.StateToGenesis(context.Background(), height)		
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Genesis file of Block!"})
		lgr.Error.Println("Request at /api/GetStateToGenesis/ Failed to retrieve the genesis file : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetStateToGenesis/ responding with a genesis file!")
	json.NewEncoder(w).Encode(responses.ConsensusGenesisResponse{consensusGenesis})
}

//GetEpoch returns the current epoch.
func GetEpoch(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if co == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	epoch, err := co.GetEpoch(context.Background(), height)		
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to retrieve Epoch of Block!"})
		lgr.Error.Println("Request at /api/GetEpoch/ Failed to retrieve Epoch : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetEpoch/ responding with an Epoch!")
	json.NewEncoder(w).Encode(responses.EpochResponse{epoch})
}

// PingNode returns a consensus block at a specific height thus signifying that it was pinged.
func PingNode(w http.ResponseWriter, r *http.Request) {
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
	
	height := consensus.HeightLatest

	//Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if co == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	_, err := co.GetBlock(context.Background(), height)		
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to ping a node by retrieving heighest block height!"})
		lgr.Error.Println("Request at /api/pingNode/ Failed to ping node : " , err)
		return
	}

	lgr.Info.Println("Request at /api/pingNode/ responding with Pong!")
	json.NewEncoder(w).Encode(responses.PongResponsed)
}

// GetBlock returns a consensus block at a specific height.
func GetBlock(w http.ResponseWriter, r *http.Request) {
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
	
		//Attempt to load a connection with the consensus client
		connection, co := loadConsensusClient(socket)
	
		//Wait for the code underneath it to execute and then close the connection
		defer connection.Close()
	
		//If a null object was retrieved send response
		if co == nil{
			//Stop the code here faild to establish connection and reply
			json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
			return
		}
	
		blk, err := co.GetBlock(context.Background(), height)		
		if err != nil{
			json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to retrieve Block!"})
			lgr.Error.Println("Request at /api/GetBlock/ Failed to retrieve a Block : " , err)
			return
		}

		lgr.Info.Println("Request at /api/GetBlock/ responding with a Block!")
		json.NewEncoder(w).Encode(responses.BlockResponse{blk})
}


// GetBlockHeader returns a consensus block header at a specific height
func GetBlockHeader(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if co == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	blk, err := co.GetBlock(context.Background(), height)		
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to retrieve Block!"})
		lgr.Error.Println("Request at /api/GetBlockHeader/ Failed to retrieve a Block : " , err)
		return
	}

	var meta mint_api.BlockMeta
	if err := cbor.Unmarshal(blk.Meta, &meta); err != nil {
		lgr.Error.Println("Request at /api/GetBlockHeader/ Failed to Unmarshal Block Metadata : " , err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to Unmarshal Block Metadata!"})
		return
	}

	lgr.Info.Println("Request at /api/GetBlockHeader/ responding with a Block Header!")
	json.NewEncoder(w).Encode(responses.BlockHeaderResponse{meta.Header})
}


// GetBlockLastCommit returns a consensus block last commit at a specific height
func GetBlockLastCommit(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if co == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	blk, err := co.GetBlock(context.Background(), height)		
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to retrieve Block!"})
		lgr.Error.Println("Request at /api/GetBlockLastCommit/ Failed to retrieve a Block : " , err)
		return
	}
	var meta mint_api.BlockMeta
	if err := cbor.Unmarshal(blk.Meta, &meta); err != nil {
		lgr.Error.Println("Request at /api/GetBlockLastCommit/ Failed to Unmarshal Block Metadata : " , err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to Unmarshal Block Metadata!"})
		return
	}
	lgr.Info.Println("Request at /api/GetBlockLastCommit/ responding with a Block Last Commit!")
	json.NewEncoder(w).Encode(responses.BlockLastCommitResponse{meta.LastCommit})
}


// GetTransactions returns a consensus block header at a specific height
func GetTransactions(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the consensus client
	connection, co := loadConsensusClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if co == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	//Use the consensus client to retrieve transactions
	transactions, err := co.GetTransactions(context.Background(), height)		
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to retrieve Transactions!"})
		lgr.Error.Println("Request at /api/GetTransactions/ Failed to retrieve Transactions : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetTransactions/ responding with all the transactions in the specified Block!")
	json.NewEncoder(w).Encode(responses.TransactionsResponse{transactions})
}