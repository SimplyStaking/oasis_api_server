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

//loadRegistryClient loads the registry client and returns it
func loadRegistryClient(socket string) (*grpc.ClientConn, registry.Backend) {
	//Attempt to load a connection with the registry client
	connection, registryClient, err := rpc.RegistryClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the registry client : ", err)
		return nil, nil
	}
	return connection, registryClient
}

//GetEntities returns all the registered entities
func GetEntities(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
		return
	}

	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Unexepcted value found, height needs to be string of int!"})
		return
	}

	//Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	entities, err := ro.GetEntities(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Entities!"})
		lgr.Error.Println("Request at /api/GetEntities/ Failed to retrieve the entities : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetEntities/ responding with a Entities!")
	json.NewEncoder(w).Encode(responses.EntitiesResponse{entities})
}

//GetNodes returns all the registered nodes
func GetNodes(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	nodes, err := ro.GetNodes(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Nodes!"})
		lgr.Error.Println("Request at /api/GetNodes/ Failed to retrieve the nodes : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetNodes/ responding with a Nodes!")
	json.NewEncoder(w).Encode(responses.NodesResponse{nodes})
}

//GetRuntimes returns all the registered entities
func GetRuntimes(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	runtimes, err := ro.GetRuntimes(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Runtimes!"})
		lgr.Error.Println("Request at /api/GetRuntimes/ Failed to retrieve the runtimes : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetRuntimes/ responding with a Runtimes!")
	json.NewEncoder(w).Encode(responses.RuntimesResponse{runtimes})
}

// GetNodeList returns the NodeList at the specified block height.
func GetNodeList(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	nodeList, err := ro.GetNodeList(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get NodeList!"})
		lgr.Error.Println("Request at /api/GetNodeList/ Failed to retrieve the nodelist : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetNodeList/ responding with a NodeList!")
	json.NewEncoder(w).Encode(responses.NodelistResponse{nodeList})
}

// GetRegistryStateToGenesis returns the StateToGenesis at the specified block height for Registry.
func GetRegistryStateToGenesis(w http.ResponseWriter, r *http.Request) {
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

	//Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	genesisRegistry, err := ro.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Registry Genesis!"})
		lgr.Error.Println("Request at /api/GetRegistryStateToGenesis/ Failed to retrieve Registry Genesis : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetRegistryStateToGenesis/ responding with a Registry Genesis!")
	json.NewEncoder(w).Encode(responses.RegistryGenesisResponse{genesisRegistry})
}

// GetEntity returns the information with regards to a single entity
func GetEntity(w http.ResponseWriter, r *http.Request) {
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
	var pubKey common_signature.PublicKey
	entityID := r.URL.Query().Get("entity")

	if len(entityID) == 0 {
		//Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetEntity/ failed, EntityID can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{"EntityID can't be empty!"})
		return
	}

	err := pubKey.UnmarshalText([]byte(entityID))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}
	//Creating a query
	query := registry.IDQuery{height, pubKey}

	registryEntity, err := ro.GetEntity(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Registry Entity!"})
		lgr.Error.Println("Request at /api/GetEntity/ Failed to retrieve Registry Entity : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetEntity/ responding with a Registry Entity!")
	json.NewEncoder(w).Encode(responses.RegistryEntityResponse{registryEntity})
}

// GetNode returns the information with regards to a single entity
func GetNode(w http.ResponseWriter, r *http.Request) {
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
	//Note Make sure that the private key that is being sent is coded properly
	//Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var pubKey common_signature.PublicKey
	nodeID := r.URL.Query().Get("nodeID")
	lgr.Warning.Println("Received : ", nodeID)

	if len(nodeID) == 0 {
		//Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetNode/ failed, NodeID can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{"NodeID can't be empty!"})
		return
	}

	err := pubKey.UnmarshalText([]byte(nodeID))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Public Key", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to UnmarshalText into Public Key."})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}
	//Creating a query
	query := registry.IDQuery{height, pubKey}

	registryNode, err := ro.GetNode(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Registry Node!"})
		lgr.Error.Println("Request at /api/GetNode/ Failed to retrieve Registry Node : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetNode/ responding with a Registry Node!")
	json.NewEncoder(w).Encode(responses.RegistryNodeResponse{registryNode})
}

// GetRuntime returns the information with regards to a single entity
func GetRuntime(w http.ResponseWriter, r *http.Request) {
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
	//Note Make sure that the private key that is being sent is coded properly
	//Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")

	if len(nmspace) == 0 {
		//Stop the code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/GetRuntime/ failed, namespace can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{"namespace can't be empty!"})
		return
	}

	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to UnmarshalText into Namespace."})
		return
	}

	// Attempt to load a connection with the registry client
	connection, ro := loadRegistryClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if ro == nil {
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}
	//Creating a query
	query := registry.NamespaceQuery{height, nameSpace}

	registryRuntime, err := ro.GetRuntime(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Registry Runtime!"})
		lgr.Error.Println("Request at /api/GetRuntime/ Failed to retrieve Registry Runtime : ", err)
		return
	}

	lgr.Info.Println("Request at /api/GetRuntime/ responding with a Registry Runtime!")
	json.NewEncoder(w).Encode(responses.RuntimeResponse{registryRuntime})
}
