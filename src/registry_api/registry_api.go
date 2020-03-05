package registry_api

import (
	"context"
	"time"
	"fmt"
	
	registry "github.com/oasislabs/oasis-core/go/registry/api"
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