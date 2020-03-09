package registry_api

import (
	"context"
	"time"
	"fmt"
	"net/http"
	"encoding/json"

	responses "github.com/SimplyVC/oasis_api_server/src/responses"
	registry "github.com/oasislabs/oasis-core/go/registry/api"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
	// committee "github.com/oasislabs/oasis-core/go/runtime/committee"
)

//Connection time out and block information possible use in the future
const (
	recvTimeout = 5 * time.Second
	numWaitedBlocks = 3
)

//Object that holds the necessary information with regards to the 
//registry client
type RegistryObject struct{
	clients map[string]registry.Backend
	ctx context.Context
}

//Set the context of the Registry Objects that need to be used
func (c *RegistryObject) SetContext(ctx context.Context){
	c.ctx = ctx
}

//Adding a client that will be used
//First check if there is a client in the Map, if not then create a map and assign it to the client
func (c *RegistryObject) AddClient(name string, client registry.Backend){
	fmt.Println("Adding node to registry object : ", name)
	if len(c.clients) == 0{
		c.clients = make(map[string]registry.Backend)
		c.clients[name] = client
	}else{
		c.clients[name] = client
	}
}

//Returns the clients map of the Registry object
func (c RegistryObject) GetClients() (map[string]registry.Backend){
	return c.clients
}

func (c RegistryObject) GetRunTimes(w http.ResponseWriter, r *http.Request){
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	if clientVal, ok := c.clients[nodeName]; ok{
		//bnyN1L2h03x1WOR4BTw/z4FdYxAJPOGIZ1SoNbM/V40=
		//Check if the nodeName is online by attempting to retreive the height of the heighest block
		runtime, err := clientVal.GetRuntimes(c.ctx, consensus.HeightLatest)
		nodes, err := clientVal.GetNodes(c.ctx, consensus.HeightLatest)
		
		fmt.Println("Nodes : ", len(nodes[0].Runtimes))
		
		for k:= range nodes{
			fmt.Println("ID of Node :", nodes[k].ID)
			fmt.Println("Entity ID of Node :", nodes[k].EntityID)
			fmt.Println("Expirate of Node  :", nodes[k].Expiration)
			fmt.Println("Committee info of Node :", nodes[k].Committee)
			fmt.Println("P2P Info :", nodes[k].P2P)
			fmt.Println("Consensus :", nodes[k].Consensus)
			fmt.Println("Runtimes :", nodes[k].Runtimes)
			fmt.Println("Roles :", nodes[k].Roles)


		}
		// nw : = committee.NewNodeDescriptorWatcher(c.ctx, clientVal)

		fmt.Println("Value of Runtime : ", len(runtime))
		
		if err != nil || runtime == nil{
			json.NewEncoder(w).Encode(responses.Response_error{"No reply from node"})
			fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + " but failed as node is offline!")
		}else{
			fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + ". Chain ID is : ")
			json.NewEncoder(w).Encode(responses.Response_pong{"Chicken"})
		}
	}else{
		json.NewEncoder(w).Encode(responses.Response_pong{"An API for " + nodeName + " needs to be setup before it can be queried"})
		fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + " but the node does not exist.")
	}
}
//Parse the JSON to read the specific client and then respond with a pong message
func (c RegistryObject) GetEntity(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	//Retrieving the name of the ndoe from the query request
	nodeName := r.URL.Query().Get("name")
	if clientVal, ok := c.clients[nodeName]; ok{
		//bnyN1L2h03x1WOR4BTw/z4FdYxAJPOGIZ1SoNbM/V40=
		//Check if the nodeName is online by attempting to retreive the height of the heighest block
		entities, err := clientVal.GetEntities(c.ctx, consensus.HeightLatest)		
		nodes, err := clientVal.GetNodes(c.ctx, consensus.HeightLatest)

		for k:= range nodes{
			fmt.Println("Value of Node :", nodes[k])
		}
		for k := range entities {
			fmt.Println("Value of K : ", entities[k])
		}

		if err != nil || entities == nil{
			json.NewEncoder(w).Encode(responses.Response_error{"No reply from node"})
			fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + " but failed as node is offline!")
		}else{
			fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + ". Chain ID is : ")
			json.NewEncoder(w).Encode(responses.Response_pong{"Chicken"})
		}
	}else{
		json.NewEncoder(w).Encode(responses.Response_pong{"An API for " + nodeName + " needs to be setup before it can be queried"})
		fmt.Println("Received request for /api/rpc/system/chain for node : " + nodeName + " but the node does not exist.")
	}
}