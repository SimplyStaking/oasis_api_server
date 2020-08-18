package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	common_namespace "github.com/oasisprotocol/oasis-core/go/common"
	common_signature "github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	registry "github.com/oasisprotocol/oasis-core/go/registry/api"
)

// loadRegistryClient loads registry client and returns it
func loadRegistryClient(socket string) (*grpc.ClientConn, registry.Backend) {

	// Attempt to load connection with registry client
	connection, registryClient, err := rpc.RegistryClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to registry"+
			" client: ", err)
		return nil, nil
	}
	return connection, registryClient
}

// GetEntities returns all registered entities
func GetEntities(w http.ResponseWriter, r *http.Request) {

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
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve entities at specific block height
	entities, err := ro.GetEntities(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get entities!"})
		lgr.Error.Println("Request at /api/registry/entities failed "+
			"to retrieve entities : ", err)
		return
	}

	// Responding with retrieved entities
	lgr.Info.Println("Request at /api/registry/entities responding with" +
		" entities!")
	json.NewEncoder(w).Encode(responses.EntitiesResponse{
		Entities: entities})
}

// GetNodes returns all registered nodes at specific block height
func GetNodes(w http.ResponseWriter, r *http.Request) {

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

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve nodes from Registry object at specific height
	nodes, err := ro.GetNodes(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Nodes!"})
		lgr.Error.Println(
			"Request at /api/registry/nodes failed to retrieve "+
				"nodes : ", err)
		return
	}

	// Respond with all nodes retrieved above
	lgr.Info.Println(
		"Request at /api/registry/nodes responding with Nodes!")
	json.NewEncoder(w).Encode(responses.NodesResponse{Nodes: nodes})
}

// GetRegistryEvents returns the events at specified block height.
func GetRegistryEvents(w http.ResponseWriter, r *http.Request) {

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

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve the events at specified block height.
	events, err := ro.GetEvents(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Events!"})
		lgr.Error.Println(
			"Request at /api/registry/events failed to retrieve "+
				"events : ", err)
		return
	}

	// Respond with events retrieved at height
	lgr.Info.Println(
		"Request at /api/registry/events responding with Events!")
	json.NewEncoder(w).Encode(responses.RegistryEventsResponse{Events: events})
}

// GetRuntimes returns all runtimes at specific block height
func GetRuntimes(w http.ResponseWriter, r *http.Request) {

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

	suspendedBool, err := strconv.ParseBool(r.URL.Query().Get("suspended"))
	if err != nil {
		suspendedBool = false
	}

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	query := registry.GetRuntimesQuery{Height: height, 
		IncludeSuspended: suspendedBool}

	// Retrieving runtimes at specific block height from registry client
	runtimes, err := ro.GetRuntimes(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get runtimes!"})
		lgr.Error.Println(
			"Request at /api/registry/runtimes failed to "+
				"retrieve runtimes : ", err)
		return
	}

	// Responding with runtimes returned above
	lgr.Info.Println("Request at /api/registry/runtimes responding " +
		"with runtimes!")
	json.NewEncoder(w).Encode(responses.RuntimesResponse{
		Runtimes: runtimes})
}

// GetRegistryStateToGenesis returns StateToGenesis at the specified
// block height for Registry.
func GetRegistryStateToGenesis(w http.ResponseWriter, r *http.Request) {

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
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieving genesis state of registry object
	genesisRegistry, err := ro.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Registry Genesis!"})
		lgr.Error.Println(
			"Request at /api/registry/genesis failed to retrieve"+
				" Registry Genesis : ", err)
		return
	}

	// Responding with genesis state retrieved above
	lgr.Info.Println(
		"Request at /api/registry/genesis responding with Registry" +
			" Genesis!")
	json.NewEncoder(w).Encode(responses.RegistryGenesisResponse{
		GenesisRegistry: genesisRegistry})
}

// GetEntity returns information with regards to single entity
func GetEntity(w http.ResponseWriter, r *http.Request) {

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
			Error: "Unexpected value found, height needs to be " +
				"a string representing an int!"})
		return
	}

	// Create public key object and retrieve entity from query
	var pubKey common_signature.PublicKey
	entityID := r.URL.Query().Get("entity")
	if len(entityID) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/registry/entity failed," +
			" EntityID can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "EntityID can't be empty!"})
		return
	}

	// Unmarshal text into public key
	err := pubKey.UnmarshalText([]byte(entityID))
	if err != nil {
		lgr.Error.Println(
			"Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Creating query to be used to retrieve Entity Information
	query := registry.IDQuery{Height: height, ID: pubKey}

	// Retrieve Entity and it's information from Registry
	// client using above query.
	registryEntity, err := ro.GetEntity(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Registry Entity!"})
		lgr.Error.Println("Request at /api/registry/entity failed to"+
			" retrieve Registry Entity : ", err)
		return
	}

	// Responding with Entity object retrieved above
	lgr.Info.Println("Request at /api/registry/entity responding with" +
		" Registry Entity!")
	json.NewEncoder(w).Encode(responses.RegistryEntityResponse{
		Entity: registryEntity})
}

// GetNode returns information with regards to single entity
func GetNode(w http.ResponseWriter, r *http.Request) {

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

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	nodeID := r.URL.Query().Get("nodeID")
	if len(nodeID) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/registry/node failed, " +
			"NodeID can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "NodeID can't be empty!"})
		return
	}

	// Unmarshal received text into public key object
	err := pubKey.UnmarshalText([]byte(nodeID))
	if err != nil {
		lgr.Error.Println(
			"Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Creating query that will be used to retrieved Node by it's ID
	query := registry.IDQuery{Height: height, ID: pubKey}

	// Retriveing node object using above query
	registryNode, err := ro.GetNode(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Registry Node!"})
		lgr.Error.Println("Request at /api/registry/node failed to "+
			"retrieve Registry Node : ", err)
		return
	}

	// Responding with retrieved node object
	lgr.Info.Println("Request at /api/registry/node responding with " +
		"Registry Node!")
	json.NewEncoder(w).Encode(responses.RegistryNodeResponse{
		Node: registryNode})
}

// GetNodeStatus returns eturns a node's status.
func GetNodeStatus(w http.ResponseWriter, r *http.Request) {

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

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	nodeID := r.URL.Query().Get("nodeID")
	if len(nodeID) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/registry/node failed, " +
			"NodeID can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "NodeID can't be empty!"})
		return
	}

	// Unmarshal received text into public key object
	err := pubKey.UnmarshalText([]byte(nodeID))
	if err != nil {
		lgr.Error.Println(
			"Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Creating query that will be used to retrieved Node by it's ID
	query := registry.IDQuery{Height: height, ID: pubKey}

	// Retriveing a node's status.
	nodeStatus, err := ro.GetNodeStatus(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Node Status!"})
		lgr.Error.Println("Request at /api/registry/nodestatus failed to "+
			"retrieve Node Status: ", err)
		return
	}

	// Responding with retrieved node object
	lgr.Info.Println("Request at /api/registry/nodestatus responding with " +
		"Node Status!")
	json.NewEncoder(w).Encode(responses.NodeStatusResponse{
		NodeStatus: nodeStatus})
}

// GetRuntime returns information with regards to single entity
func GetRuntime(w http.ResponseWriter, r *http.Request) {

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

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {
		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/registry/runtime failed" +
			", namespace can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "namespace can't be empty!"})
		return
	}

	// Unmarshal received text into namespace object
	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Namespace."})
		return
	}

	// Attempt to load connection with registry client
	connection, ro := loadRegistryClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if ro == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Creating query that will be used to return runtime by it's namespace
	query := registry.NamespaceQuery{Height: height, ID: nameSpace}

	// Retrieving runtime object using above query
	registryRuntime, err := ro.GetRuntime(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Registry Runtime!"})
		lgr.Error.Println("Request at /api/registry/runtime failed "+
			"to retrieve Registry Runtime : ", err)
		return
	}

	// Responding with runtime object retrieved above
	lgr.Info.Println("Request at /api/registry/runtime responding with " +
		"Registry Runtime!")
	json.NewEncoder(w).Encode(responses.RuntimeResponse{
		Runtime: registryRuntime})
}
