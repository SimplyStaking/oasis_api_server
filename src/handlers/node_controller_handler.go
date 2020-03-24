package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	control "github.com/oasislabs/oasis-core/go/control/api"
)

// loadNodeControllerClient loads node controller client and returns it
func loadNodeControllerClient(socket string) (*grpc.ClientConn, control.NodeController) {

	// Attempt to load connection with staking client
	connection, nodeControllerClient, err := rpc.NodeControllerClient(socket)
	if err != nil {
		lgr.Error.Println(
			"Failed to establish connection to NodeController client : ", err)
		return nil, nil
	}
	return connection, nodeControllerClient
}

// GetIsSynced checks whether node has finished syncing.
func GetIsSynced(w http.ResponseWriter, r *http.Request) {

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

	// Attempt to load connection with staking client
	connection, nc := loadNodeControllerClient(socket)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if nc == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using socket : " + socket})
		return
	}

	// Retrieving synchronized state from node controller client
	synced, err := nc.IsSynced(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get IsSynced!"})
		lgr.Error.Println(
			"Request at /api/nodecontroller/synced/ Failed to retrieve IsSynced : ", err)
		return
	}

	// Responding with retrieved synchronizatio state above
	lgr.Info.Println(
		"Request at /api/nodecontroller/synced/ responding with IsSynced State!")
	json.NewEncoder(w).Encode(responses.IsSyncedResponse{Synced: synced})
}
