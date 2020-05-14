package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	common_namespace "github.com/oasislabs/oasis-core/go/common"
	scheduler "github.com/oasislabs/oasis-core/go/scheduler/api"
)

// loadSchedulerClient loads scheduler client and returns it
func loadSchedulerClient(socket string) (*grpc.ClientConn, scheduler.Backend) {

	// Attempt to load connection with scheduler client
	connection, schedulerClient, err := rpc.SchedulerClient(socket)
	if err != nil {
		lgr.Error.Println(
			"Failed to establish connection to scheduler client : ",
			err)
		return nil, nil
	}
	return connection, schedulerClient
}

// GetValidators returns vector of consensus validators for given epoch.
func GetValidators(w http.ResponseWriter, r *http.Request) {

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

	// Retrieving height of node from query request
	recvHeight := r.URL.Query().Get("height")
	height := checkHeight(recvHeight)
	if height == -1 {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}

	// Attempt to load connection with scheduler client
	connection, sc := loadSchedulerClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if sc == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve validators at given block height
	validators, err := sc.GetValidators(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Validators!"})
		lgr.Error.Println("Request at /api/scheduler/validators/ "+
			"failed to retrieve validators : ", err)
		return
	}

	// Responding with Validators retrieved from scheduler client
	lgr.Info.Println("Request at /api/scheduler/validators/ responding " +
		"with Validators!")
	json.NewEncoder(w).Encode(responses.ValidatorsResponse{
		Validators: validators})
}

// GetCommittees returns vector of committees for given
// runtime ID, at specified block height, and optional callback
// for querying beacon for given epoch/block height.
func GetCommittees(w http.ResponseWriter, r *http.Request) {

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
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}

	// Note Make sure that private key that is being sent is coded properly
	// Example A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU+h+blS9pto= should be
	// A1X90rT/WK4AOTh/dJsUlOqNDV/nXM6ZU%2Bh%2BblS9pto=
	var nameSpace common_namespace.Namespace
	nmspace := r.URL.Query().Get("namespace")
	if len(nmspace) == 0 {

		// Stop code here no need to establish connection and reply
		lgr.Warning.Println("Request at /api/registry/runtime/ failed" +
			", namespace can't be empty!")
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "namespace can't be empty!"})
		return
	}

	// Unmarshal text into namespace object to be used in query
	err := nameSpace.UnmarshalText([]byte(nmspace))
	if err != nil {
		lgr.Error.Println("Failed to UnmarshalText into Namespace", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to UnmarshalText into Namespace."})
		return
	}

	// Attempt to load connection with scheduler client
	connection, sc := loadSchedulerClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if sc == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Create query to be used to retrieve Committees
	query := scheduler.GetCommitteesRequest{Height: height,
		RuntimeID: nameSpace}

	// Retrieving Committees using query above
	committees, err := sc.GetCommittees(context.Background(), &query)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Committees!"})
		lgr.Error.Println("Request at /api/scheduler/committees/ "+
			"failed to retrieve committees : ", err)
		return
	}

	// Responding with committees that were retrieved from scheduler client
	lgr.Info.Println("Request at /api/scheduler/committees/ responding " +
		"with Committees!")
	json.NewEncoder(w).Encode(responses.CommitteesResponse{
		Committee: committees})
}

// GetSchedulerStateToGenesis returns genesis state of scheduler at the
// specified block height.
func GetSchedulerStateToGenesis(w http.ResponseWriter, r *http.Request) {

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
			Error: "Unexepcted value found, height needs to be " +
				"string of int!"})
		return
	}

	// Attempt to load connection with scheduler client
	connection, sc := loadSchedulerClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if sc == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket: " +
				socket})
		return
	}

	// Retrieve genesis state of scheduler at specific block height
	gensis, err := sc.StateToGenesis(context.Background(), height)
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Scheduler Genesis State!"})
		lgr.Error.Println("Request at /api/scheduler/genesis/ failed "+
			"to retrieve Scheduler Genesis State : ", err)
		return
	}

	// Responding with genesis state retrieved above
	lgr.Info.Println("Request at /api/scheduler/genesis/ responding with " +
		"scheduler genesis state!")
	json.NewEncoder(w).Encode(responses.SchedulerGenesisState{
		SchedulerGenesisState: gensis})
}
