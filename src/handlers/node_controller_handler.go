package handlers


import (
	"context"
	"net/http"
	"encoding/json"
	"google.golang.org/grpc"

	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	control "github.com/oasislabs/oasis-core/go/control/api"
	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

//loadNodeControllerClient loads the node controller client and returns it
func loadNodeControllerClient(socket string) (*grpc.ClientConn, control.NodeController){
	//Attempt to load a connection with the staking client
	connection, nodeControllerClient, err := rpc.NodeControllerClient(socket)
	if err != nil {
		lgr.Error.Println("Failed to establish connection to the NodeController client : ", err)
		return nil, nil
	}
	return connection, nodeControllerClient
}

//GetIsSynced checks whether the node has finished syncing.
func GetIsSynced(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, socket := checkNodeName(nodeName)
	if confirmation == false {
		//Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Node name requested doesn't exist"})
		return
	}

	//Attempt to load a connection with the staking client
	connection, nc := loadNodeControllerClient(socket)

	//Wait for the code underneath it to execute and then close the connection
	defer connection.Close()

	//If a null object was retrieved send response
	if nc == nil{
		//Stop the code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
		return
	}

	synced, err := nc.IsSynced(context.Background())
	if err != nil{
		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get IsSynced!"})
		lgr.Error.Println("Request at /api/GetIsSynced/ Failed to retrieve the IsSynced : " , err)
		return
	}

	lgr.Info.Println("Request at /api/GetIsSynced/ responding with the IsSynced State!")
	json.NewEncoder(w).Encode(responses.IsSyncedResponse{synced})
}