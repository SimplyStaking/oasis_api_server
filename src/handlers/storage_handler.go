package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"google.golang.org/grpc"

// 	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
// 	responses "github.com/SimplyVC/oasis_api_server/src/responses"
// 	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
// 	strg "github.com/oasislabs/oasis-core/go/storage"
// 	api "github.com/oasislabs/oasis-core/go/storage/api"
// )

// //loadStorageClient loads the storage client and returns it
// func loadStorageClient(socket string) (*grpc.ClientConn, api.Backend) {
// 	//Attempt to load a connection with the storage client
// 	connection, storageClient, err := rpc.StorageClient(socket)
// 	if err != nil {
// 		lgr.Error.Println("Failed to establish connection to the storage client : ", err)
// 		return nil, nil
// 	}
// 	return connection, storageClient
// }

// //GetConnectionMetrics returns all the connections to the node
// func GetConnectionMetrics(w http.ResponseWriter, r *http.Request) {
// 	//Adding a header so that the receiver knows they are receiving a JSON structure
// 	w.Header().Add("Content-Type", "application/json")
// 	//Retrieving the name of the ndoe from the query request
// 	nodeName := r.URL.Query().Get("name")
// 	confirmation, socket := checkNodeName(nodeName)
// 	if confirmation == false {
// 		//Stop the code here no need to establish connection and reply
// 		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Node name requested doesn't exist"})
// 		return
// 	}

// 	connection, st := loadStorageClient(socket)

// 	//Wait for the code underneath it to execute and then close the connection
// 	defer connection.Close()

// 	//If a null object was retrieved send response
// 	if st == nil {
// 		//Stop the code here faild to establish connection and reply
// 		json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to establish a connection using the socket : " + socket})
// 		return
// 	}

// 	//Create a new MetricsWrapper to be used
// 	metricsBackend := strg.newMetricsWrapper(st)
// 	nodesReturned := metricsBackend.GetConnectedNodes()

// 	lgr.Warning.Println(len(nodesReturned))
// 	// if err != nil {
// 	// 	json.NewEncoder(w).Encode(responses.ErrorResponse{"Failed to get Entities!"})
// 	// 	lgr.Error.Println("Request at /api/GetEntities/ Failed to retrieve the entities : ", err)
// 	// 	return
// 	// }

// 	lgr.Info.Println("Request at /api/GetEntities/ responding with a Entities!")
// 	// json.NewEncoder(w).Encode(responses.EntitiesResponse{entities})
// }
