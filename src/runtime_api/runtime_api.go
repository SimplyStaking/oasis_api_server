package runtime_api

import (
	"context"
	"fmt"

	runtime "github.com/oasislabs/oasis-core/go/runtime/client/api"
)

//Object that holds the necessary information with regards to the runtime client
type RuntimeObject struct{
	clients map[string]runtime.RuntimeClient
	ctx context.Context
}

//Set the context of the runtime Objects that need to be used
func (c *RuntimeObject) SetContext(ctx context.Context){
	c.ctx = ctx
}

//Adding a client that will be used
//First check if there is a client in the Map, if not then create a map and assign it to the client
func (c *RuntimeObject) AddClient(name string, client runtime.RuntimeClient){
	fmt.Println("Adding node to runtime object : ", name)
	if len(c.clients) == 0{
		c.clients = make(map[string]runtime.RuntimeClient)
		c.clients[name] = client
	}else{
		c.clients[name] = client
	}
}

//Returns the clients map of the Registry object
func (c RuntimeObject) GetClients() (map[string]runtime.RuntimeClient){
	return c.clients
}

// //Parse the JSON to read the specific client and then respond with a pong message
// func (c RuntimeObject) Pong(w http.ResponseWriter, r *http.Request) {
// 	//Adding a header so that the receiver knows they are receiving a JSON structure
// 	w.Header().Add("Content-Type", "application/json")
// 	//Retrieving the name of the ndoe from the query request
// 	nodeName := r.URL.Query().Get("name")
// 	//Check if the node being pinged exists
// 	if clientVal, ok := c.clients[nodeName]; ok{
// 		//Check if the nodeName is online by attempting to retreive the height of the heighest block
// 		blk, err := clientVal.GetBlock(c.ctx, consensus.HeightLatest)
// 		if err != nil || blk == nil{
// 			json.NewEncoder(w).Encode(responses.Response_error{"No reply from node"})
// 			fmt.Println("Received request for /api/pingNode for node : " + nodeName + " but failed as node is offline!")
// 		}else{
// 			fmt.Println("Received request for /api/pingNode for node : " + nodeName + ". Current Block Height : " + strconv.FormatInt(int64(blk.Height), 10))
// 			json.NewEncoder(w).Encode(responses.Responded_pong)
// 		}
// 	}else{
// 		json.NewEncoder(w).Encode(responses.Response_pong{"An API for " + nodeName + " needs to be setup before it can be queried"})
// 		fmt.Println("Received request for /api/pingNode for node : " + nodeName + " but the node does not exist.")
// 	}
// }