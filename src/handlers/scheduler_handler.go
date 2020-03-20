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
	scheduler "github.com/oasislabs/oasis-core/go/scheduler/api"
)

// loadSchedulerClient loads the scheduler client and returns it
func loadSchedulerClient(socket string) (*grpc.ClientConn, scheduler.Backend) {
	// Attempt to load a connection with the scheduler client
	connection, schedulerClient, err := rpc.SchedulerClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the scheduler client : ", err)
		return nil, nil
	}
	return connection, schedulerClient
}

// GetValidators returns the vector of consensus validators for a given epoch.
func GetValidators(w http.ResponseWriter, r *http.Request) {
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

	// Retrieving the height of the node from the query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {
		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Unexepcted value found, height needs to be string of int!"})
		return
	}

	// Attempt to load a connection with the scheduler client
	connection, sc := loadSchedulerClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if sc == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieve validators at a given block height
	validators, err := sc.GetValidators(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Validators!"})
		lgr.Error.Println("Request at /api/GetValidators/ Failed to retrieve the validators : ", err)
		return
	}

	// Responding with Validators retrieved from the scheduler client
	lgr.Info.Println("Request at /api/GetValidators/ responding with Validators!")
	json.NewEncoder(w).Encode(responses.ValidatorsResponse{Validators: validators})
}

// GetCommittees returns the vector of committees for a given
// runtime ID, at the specified block height, and optional callback
// for querying the beacon for a given epoch/block height.
func GetCommittees(w http.ResponseWriter, r *http.Request) {
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

	// Unmarshal text into namespace object to be used in a query
	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to UnmarshalText into Namespace."})
		return
	}

	// Attempt to load a connection with the scheduler client
	connection, sc := loadSchedulerClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if sc == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Create a query to be used to retrieve Committees
	query := scheduler.GetCommitteesRequest{Height: height, RuntimeID: nameSpace}

	// Retrieving the Committees using the query above
	committees, err := sc.GetCommittees(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Committees!"})
		lgr.Error.Println("Request at /api/GetCommittees/ Failed to retrieve the committees : ", err)
		return
	}

	// Responding with committees that were retrieved from the scheduler client
	lgr.Info.Println("Request at /api/GetCommittees/ responding with Committees!")
	json.NewEncoder(w).Encode(responses.CommitteesResponse{Committee: committees})
}

// GetSchedulerStateToGenesis returns the genesis state of the scheduler at specified block height.
func GetSchedulerStateToGenesis(w http.ResponseWriter, r *http.Request) {
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

	// Attempt to load a connection with the scheduler client
	connection, sc := loadSchedulerClient(socket)

	// Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	// If a null object was retrieved send response
	if sc == nil {
		// Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to establish a connection using the socket : " + socket})
		return
	}

	// Retrieve the genesis state of the scheduler at a specific block height
	gensis, err := sc.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to get Scheduler Genesis State!"})
		lgr.Error.Println("Request at /api/GetSchedulerStateToGenesis/ Failed to retrieve the Scheduler Genesis State : ", err)
		return
	}

	// Responding with the genesis state retrieved above
	lgr.Info.Println("Request at /api/GetSchedulerStateToGenesis/ responding with scheduler genesis state!")
	json.NewEncoder(w).Encode(responses.SchedulerGenesisState{SchedulerGenesisState: gensis})
}
