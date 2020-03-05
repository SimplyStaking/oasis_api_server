package staking_api

import (
	"context"
	"fmt"

	staking "github.com/oasislabs/oasis-core/go/staking/api"
)

//Object that holds the necessary information with regards to the staking client
type StakingObject struct{
	clients map[string]staking.Backend
	ctx context.Context
}

//Set the context of the Staking Objects that need to be used
func (c *StakingObject) SetContext(ctx context.Context){
	c.ctx = ctx
}

//Adding a client that will be used
//First check if there is a client in the Map, if not then create a map and assign it to the client
func (c *StakingObject) AddClient(name string, client staking.Backend){
	fmt.Println("Adding node to staking object : ", name)
	if len(c.clients) == 0{
		c.clients = make(map[string]staking.Backend)
		c.clients[name] = client
	}else{
		c.clients[name] = client
	}
}

//Returns the clients map of the Registry object
func (c StakingObject) GetClients() (map[string]staking.Backend){
	return c.clients
}