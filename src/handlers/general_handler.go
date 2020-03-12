package handlers

import (
	"net/http"
	"encoding/json"
	
	config "github.com/SimplyVC/oasis_api_server/src/config"
	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

//General ping function to ping if the entire API is online
func Pong(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	lgr.Info.Println("Received request for /api/pingApi")
	json.NewEncoder(w).Encode(responses.Responded_pong)
}

//Get all the possible connections that have been loaded in the configuration file
func GetConnections(w http.ResponseWriter, r *http.Request){
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	lgr.Info.Println("Received request for /api/getConnectionsList")

	//Create new empty Slice of strings where the connections will be stored
	connectionsResponse := []string{} 
	allSockets := config.GetSockets()

	lgr.Info.Println("Iterating through all socket connections.")
	for _, socket := range allSockets{
		lgr.Info.Println("Node : ", socket["node_name"]," has socket at : ", socket["ws_url"])
		connectionsResponse = append(connectionsResponse, socket["ws_url"])
	}
	//Encode the object and send it using a predefind response
	json.NewEncoder(w).Encode(responses.Response_Conns{connectionsResponse})
}