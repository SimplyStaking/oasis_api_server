package handlers

import (
	"sync"
	"encoding/json"
	"net/http"

	"github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
)

// Pong responds with ping if entire API is online
func Pong(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")
	lgr.Info.Println("Received request for /api/pingapi")
	json.NewEncoder(w).Encode(responses.SuccessResponsed)
}

// GetConnections retrieves the node names that are configured in the API
func GetConnections(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")
	lgr.Info.Println("Received request for /api/getconnectionslist")

	mutex := &sync.RWMutex{}
	mutex.Lock()

	// Create new empty Slice of strings where connections will be stored
	connectionsResponse := []string{}
	allSockets := config.GetNodes()

	lgr.Info.Println("Iterating through all socket connections.")
	for _, socket := range allSockets {
		lgr.Info.Printf("Node: %s has socket %s \n",socket["node_name"], 
			socket["isocket_path"])
		connectionsResponse = append(connectionsResponse, socket["node_name"])
	}
	// Encode object and send it using predefind response
	json.NewEncoder(w).Encode(responses.ConnectionsResponse{
		Results: connectionsResponse})
	mutex.Unlock()
}
