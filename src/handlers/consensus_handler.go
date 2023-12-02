package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	"github.com/oasisprotocol/oasis-core/go/common/cbor"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	consensus "github.com/oasisprotocol/oasis-core/go/consensus/api"
	mint_api "github.com/oasisprotocol/oasis-core/go/consensus/cometbft/api"
	"github.com/oasisprotocol/oasis-core/go/consensus/cometbft/crypto"
)

// loadConsensusClient loads consensus client and returns it
func loadConsensusClient(socket string) (*grpc.ClientConn,
	consensus.ClientBackend) {

	// Attempt to load connection with consensus client
	connection, consensusClient, err := rpc.ConsensusClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to consensus"+
			" client : ", err)
		return nil, nil
	}
	return connection, consensusClient
}

// GetConsensusStateToGenesis returns genesis state
// at specified block height for Consensus.
func GetConsensusStateToGenesis(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieve height from query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieving genesis state of consensus object at specified height
	consensusGenesis, err := co.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Genesis file of Block!"})

		lgr.Error.Println("Request at /api/consensus/genesis failed "+
			"to retrieve genesis file : ", err)
		return
	}

	// Responding with consensus genesis state object, retrieved above.
	lgr.Info.Println("Request at /api/consensus/genesis responding with" +
		" genesis file!")
	json.NewEncoder(w).Encode(responses.ConsensusGenesisResponse{
		GenJSON: consensusGenesis})
}

// GetEpoch returns current epoch of given block height
func GetEpoch(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieve height from query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {
		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Return beacon backend object from consensus
	beacon := co.Beacon()

	// Return Epoch at current block height
	epoch, err := beacon.GetEpoch(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve Epoch of Block!"})

		lgr.Error.Println("Request at /api/consensus/epoch failed to"+
			" retrieve Epoch : ", err)
		return
	}

	// Respond with retrieved epoch above
	lgr.Info.Println("Request at /api/consensus/epoch responding" +
		" with an Epoch!")
	json.NewEncoder(w).Encode(responses.EpochResponse{Ep: epoch})
}

// PingNode returns consensus block at specific height
// thus signifying that it was pinged.
func PingNode(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")
	lgr.Info.Println("Received request for /api/pingnode")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {
		lgr.Info.Println("Node name requested doesn't exist")
		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Setting height to latest
	height := consensus.HeightLatest

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Making sure that the error being retrieved is nill, meaning that API
	// is pingable
	_, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to ping node by retrieving highest " +
				"block height!"})

		lgr.Error.Println("Request at /api/pingnode failed to ping"+
			" node : ", err)
		return
	}

	// Responding with Pong response
	lgr.Info.Println("Request at /api/pingnode responding with Pong!")
	json.NewEncoder(w).Encode(responses.SuccessResponsed)
}

// GetBlock returns consensus block at specific height.
func GetBlock(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieve height from query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve block at specific height from consensus client
	blk, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve Block!"})

		lgr.Error.Println("Request at /api/consensus/block failed "+
			"to retrieve Block : ", err)
		return
	}

	// Responding with retrieved block
	lgr.Info.Println("Request at /api/consensus/block responding with Block!")
	json.NewEncoder(w).Encode(responses.BlockResponse{Blk: blk})
}

// GetStatus returns the current status overview.
func GetStatus(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve the current status overview
	status, err := co.GetStatus(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve Status!"})

		lgr.Error.Println("Request at /api/consensus/status failed "+
			"to retrieve Status : ", err)
		return
	}

	// Responding with retrieved block
	lgr.Info.Println("Request at /api/consensus/status responding with Status!")
	json.NewEncoder(w).Encode(responses.StatusResponse{Status: status})
}

// GetGenesisDocument returns the original genesis document.
func GetGenesisDocument(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve the current status overview
	genesisDocument, err := co.GetGenesisDocument(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve Status!"})

		lgr.Error.Println("Request at /api/consensus/genesisdocument failed "+
			"to retrieve Genesis Document : ", err)
		return
	}

	// Responding with retrieved block
	lgr.Info.Println(
		"Request at /api/consensus/genesisdocument responding with Genesis " +
			"Document!")
	json.NewEncoder(w).Encode(responses.GenesisDocumentResponse{GenesisDocument: genesisDocument})
}

// GetBlockHeader returns consensus block header at specific height
func GetBlockHeader(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retriving Block at specific height using Consensus client
	blk, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve Block!"})

		lgr.Error.Println("Request at /api/consensus/blockheader "+
			"failed to retrieve Block : ", err)
		return
	}

	// Creating BlockMeta object
	var meta mint_api.BlockMeta
	if err := cbor.Unmarshal(blk.Meta, &meta); err != nil {
		lgr.Error.Println("Request at /api/consensus/blockheader "+
			"failed to Unmarshal Block Metadata : ", err)

		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to Unmarshal Block Metadata!"})
		return
	}

	// Responds with block header retrieved above
	lgr.Info.Println("Request at /api/consensus/blockheader responding " +
		"with Block Header!")
	json.NewEncoder(w).Encode(responses.BlockHeaderResponse{
		BlkHeader: meta.Header})
}

// GetBlockLastCommit returns consensus block last commit at specific height
func GetBlockLastCommit(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve block at specific height from consensus client
	blk, err := co.GetBlock(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve Block!"})

		lgr.Error.Println("Request at /api/consensus/blocklastcommit "+
			"failed to retrieve Block : ", err)
		return
	}

	// Creating BlockMeta object
	var meta mint_api.BlockMeta
	if err := cbor.Unmarshal(blk.Meta, &meta); err != nil {
		lgr.Error.Println("Request at /api/consensus/blocklastcommit "+
			"failed Unmarshal Block Metadata : ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to Unmarshal Block Metadata!"})
		return
	}
	// Responds with Block Last commit retrieved above
	lgr.Info.Println("Request at /api/consensus/blocklastcommit " +
		"responding with Block Last Commit!")
	json.NewEncoder(w).Encode(responses.BlockLastCommitResponse{
		BlkLastCommit: meta.LastCommit})
}

// PublicKeyToAddress accepts a Consensus Public Key and respond with
// crypto.address which is used to match consensus public keys with
// Tendermint Addresses
func PublicKeyToAddress(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving consensus public key from the query
	consensusKey := r.URL.Query().Get("consensus_public_key")
	if consensusKey == "" {
		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "No Consensus Key Provided"})
		return
	}
	consensusPublicKey := &signature.PublicKey{}

	err := consensusPublicKey.UnmarshalText([]byte(consensusKey))
	if err != nil {
		lgr.Error.Println("Request at /api/consensus/pubkeyaddress "+
			"failed to Unmarshal Consensus PublicKey : ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to Unmarshal Public Key!"})
		return
	}
	// Convert the consensusKey into a signature PublicKey
	tendermintKey := crypto.PublicKeyToCometBFT(consensusPublicKey)
	cryptoAddress := tendermintKey.Address()
	// Responds with transactions retrieved above
	lgr.Info.Println("Request at /api/consensus/pubkeyaddress responding " +
		"with Tendermint Public Key Address!")
	json.NewEncoder(w).Encode(responses.TendermintAddress{
		TendermintAddress: &cryptoAddress})
}

// GetTransactions returns consensus block header at specific height
func GetTransactions(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if !confirmation {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Retrieving height from query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with consensus client
	connection, co := loadConsensusClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if co == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Use consensus client to retrieve transactions at specific block
	// height
	transactions, err := co.GetTransactions(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve Transactions!"})

		lgr.Error.Println("Request at /api/consensus/transactions "+
			"failed to retrieve Transactions : ", err)
		return
	}

	// Responds with transactions retrieved above
	lgr.Info.Println("Request at /api/consensus/transactions responding" +
		"with all transactions in specified Block!")
	json.NewEncoder(w).Encode(responses.TransactionsResponse{
		Transactions: transactions})
}
