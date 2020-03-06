package consensus_api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"

	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
)

//Object that holds the necessary information with regards to the Consensus Clients
type ConsensusObject struct{
	clients map[string]consensus.ClientBackend
	ctx context.Context
}

//Set the context of the Consensus Objects that need to be used
func (c *ConsensusObject) SetContext(ctx context.Context){
	c.ctx = ctx
}

//Returns the clients map of the Consensus object
func (c ConsensusObject) GetClients() (map[string]consensus.ClientBackend){
	return c.clients
}

//Adding a client that will be used
//First check if there is a client in the Map, if not then create a map and assign it to the client
func (c *ConsensusObject) AddClient(name string, client consensus.ClientBackend){
	fmt.Println("Adding node to consensus object : ", name)
	if len(c.clients) == 0{
		c.clients = make(map[string]consensus.ClientBackend)
		c.clients[name] = client
	}else{
		c.clients[name] = client
	}
}

//Parse the JSON to read the specific client and then respond with a pong message
func (c ConsensusObject) Pong(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	//Check if the node being pinged exists
	if clientVal, ok := c.clients[nodeName]; ok{
		//Check if the nodeName is online by attempting to retreive the height of the heighest block
		blk, err := clientVal.GetBlock(c.ctx, consensus.HeightLatest)
		if err != nil || blk == nil{
			json.NewEncoder(w).Encode(responses.Response_error{"No reply from node"})
			fmt.Println("Received request for /api/pingNode for node : " + nodeName + " but failed as node is offline!")
		}else{
			fmt.Println("Received request for /api/pingNode for node : " + nodeName + ". Current Block Height : " + strconv.FormatInt(int64(blk.Height), 10))
			json.NewEncoder(w).Encode(responses.Responded_pong)
		}
	}else{
		json.NewEncoder(w).Encode(responses.Response_pong{"An API for " + nodeName + " needs to be setup before it can be queried"})
		fmt.Println("Received request for /api/pingNode for node : " + nodeName + " but the node does not exist.")
	}
}

//Parse the JSON to read the specific client and then respond with a pong message
func (c ConsensusObject) GetChainID(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	if clientVal, ok := c.clients[nodeName]; ok{
		//Check if the nodeName is online by attempting to retreive the height of the heighest block
		consensusGenesis, err := clientVal.StateToGenesis(c.ctx, consensus.HeightLatest)		
		if err != nil || consensusGenesis == nil{
			json.NewEncoder(w).Encode(responses.Response_error{"No reply from node"})
			fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + " but failed as node is offline!")
		}else{
			fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + ". Chain ID is : " + consensusGenesis.ChainID)
			json.NewEncoder(w).Encode(responses.Response_pong{consensusGenesis.ChainID})
		}
	}else{
		json.NewEncoder(w).Encode(responses.Response_pong{"An API for " + nodeName + " needs to be setup before it can be queried"})
		fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + " but the node does not exist.")
	}
}

//Get all the possible connections
func (c ConsensusObject) GetConnectionslist(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	fmt.Println("Received request for /api/getConnectionsList")

	//Create new empty Slice of strings where the connections will be stored
	connectionsResponse := []string{} 

	//Iterate through all the clients and retrieve their Keys without retrieving the Values
	//Note Iteration is random
	for k, _ := range c.clients {
		fmt.Println("Iterating through Consensus Object Clients : ",k)
		connectionsResponse = append(connectionsResponse, k)
	}
	//Encode the object and send it using a predefind response
	json.NewEncoder(w).Encode(responses.Response_Conns{connectionsResponse})
}