package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	common_namespace "github.com/oasislabs/oasis-core/go/common"
	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	registry "github.com/oasislabs/oasis-core/go/registry/api"
)

// loadRegistryClient loads the registry client and returns it
func loadRegistryClient(socket string) (*grpc.ClientConn, registry.Backend) {
	// Attempt to load a connection with the registry client
	connection, registryClient, err := rpc.RegistryClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the registry client : ", err)
		return nil, nil
	}
	return connection, registryClient
}

// GetEntities returns all the registered entities
func GetEntities(w http.ResponseWriter, r *http.Request) {
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

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if ro == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieve the entities at a specific block height
	entities, err := ro.GetEntities(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Entities!"})
		lgr.Error.Println("Request at /api/GetEntities/ Failed to retrieve the entities : ", err)
		return
	}

	// Responding with the retrieved entities
	lgr.Info.Println("Request at /api/GetEntities/ responding with a Entities!")
	json.NewEncoder(w).Encode(responses.EntitiesResponse{Entities: entities})
}

// GetNodes returns all the registered nodes at a specific block height
func GetNodes(w http.ResponseWriter, r *http.Request) {
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

	// Retrieving the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if ro == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieve the nodes from the Registry object at a specific height
	nodes, err := ro.GetNodes(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Nodes!"})
		lgr.Error.Println("Request at /api/GetNodes/ Failed to retrieve the nodes : ", err)
		return
	}

	// Respond with all the nodes retrieved above
	lgr.Info.Println("Request at /api/GetNodes/ responding with a Nodes!")
	json.NewEncoder(w).Encode(responses.NodesResponse{Nodes: nodes})
}

// GetRuntimes returns all the runtimes at a specific block height
func GetRuntimes(w http.ResponseWriter, r *http.Request) {
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

	// Retrieving the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if ro == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieving the runtimes at a specific block height from the registry client
	runtimes, err := ro.GetRuntimes(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Runtimes!"})
		lgr.Error.Println("Request at /api/GetRuntimes/ Failed to retrieve the runtimes : ", err)
		return
	}

	// Responding with the runtimes returned above
	lgr.Info.Println("Request at /api/GetRuntimes/ responding with a Runtimes!")
	json.NewEncoder(w).Encode(responses.RuntimesResponse{Runtimes: runtimes})
}

// GetRegistryStateToGenesis returns the StateToGenesis at the specified block height for Registry.
func GetRegistryStateToGenesis(w http.ResponseWriter, r *http.Request) {
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

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if ro == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieving the genesis state of the registry object
	genesisRegistry, err := ro.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Registry Genesis!"})
		lgr.Error.Println("Request at /api/GetRegistryStateToGenesis/ Failed to retrieve Registry Genesis : ", err)
		return
	}

	// Responding with the genesis state retrieved above
	lgr.Info.Println("Request at /api/GetRegistryStateToGenesis/ responding with a Registry Genesis!")
	json.NewEncoder(w).Encode(responses.RegistryGenesisResponse{GenesisRegistry: genesisRegistry})
}

// GetEntity returns the information with regards to a single entity
func GetEntity(w http.ResponseWriter, r *http.Request) {
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

	// Create a public key object and retrieve the entity from the query
	var pubKey common_signature.PublicKey
	entityID := r.URL.Query().Get("entity")
	if len(entityID) == 0 {
		// Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetEntity/ failed, EntityID can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "EntityID can't be empty!"})
		return
	}

	// Unmarshal the text into a public key
	err := pubKey.UnmarshalText([]byte(entityID))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if ro == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Creating a query to be used to retrieve Entity Information
	query := registry.IDQuery{Height: height, ID: pubKey}

	// Retrive the Entity and it's information from the Registry client using the above query.
	registryEntity, err := ro.GetEntity(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Registry Entity!"})
		lgr.Error.Println("Request at /api/GetEntity/ Failed to retrieve Registry Entity : ", err)
		return
	}

	// Responding with the Entity object retrieved above
	lgr.Info.Println("Request at /api/GetEntity/ responding with a Registry Entity!")
	json.NewEncoder(w).Encode(responses.RegistryEntityResponse{Entity: registryEntity})
}

// GetNode returns the information with regards to a single entity
func GetNode(w http.ResponseWriter, r *http.Request) {
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

	// Retrieving the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that the private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	nodeID := r.URL.Query().Get("nodeID")
	if len(nodeID) == 0 {
		// Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetNode/ failed, NodeID can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "NodeID can't be empty!"})
		return
	}

	// Unmarshal the recieved text into a public key object
	err := pubKey.UnmarshalText([]byte(nodeID))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if ro == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Creating a query that will be used to retrieved a Node by it's ID
	query := registry.IDQuery{Height: height, ID: pubKey}

	// Retriveing the node object using the above query
	registryNode, err := ro.GetNode(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Registry Node!"})
		lgr.Error.Println("Request at /api/GetNode/ Failed to retrieve Registry Node : ", err)
		return
	}

	// Responding with the retrieved node object
	lgr.Info.Println("Request at /api/GetNode/ responding with a Registry Node!")
	json.NewEncoder(w).Encode(responses.RegistryNodeResponse{Node: registryNode})
}

// GetRuntime returns the information with regards to a single entity
func GetRuntime(w http.ResponseWriter, r *http.Request) {
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

	// Retrieve the height from the query
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Note Make sure that the private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {
		// Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetRuntime/ failed, namespace can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "namespace can't be empty!"})
		return
	}

	// Unmarshal the received text into a namespace object
	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to UnmarshalText into Namespace."})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if ro == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Creating a query that will be used to return a runtime by it's namespace
	query := registry.NamespaceQuery{Height: height, ID: nameSpace}

	// Retrieving the runtime object using the above query
	registryRuntime, err := ro.GetRuntime(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Registry Runtime!"})
		lgr.Error.Println("Request at /api/GetRuntime/ Failed to retrieve Registry Runtime : ", err)
		return
	}

	// Responding with the runtime object retrieved above
	lgr.Info.Println("Request at /api/GetRuntime/ responding with a Registry Runtime!")
	json.NewEncoder(w).Encode(responses.RuntimeResponse{Runtime: registryRuntime})
}
