package staking_api_test

import (
	"context"
	"testing"
	"fmt"
	"os"

	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	staking_api "github.com/SimplyVC/oasis_api_server/src/staking_api"
)

var (
	ctx context.Context
	co  staking_api.StakingObject
	ws_url = "unix:/home/vvol/serverdir/node/internal.sock"
	ws_url_Invalid = "unix:/home/vvol/serverdir/nodes/internal.sock"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
 
func setup() {
	ctx = context.Background()
	co = staking_api.StakingObject{}

	co.SetContext(ctx)

	_, stakingClient, err := rpc.StakingClient(ws_url)
	if err != nil {
		panic(err)
	}
	
	co.AddClient("Oasis_Local",stakingClient)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}
 
func teardown() {

	ctx = context.Background()
	co = staking_api.StakingObject{}

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

//Testing if the setup indeed added one client
func TestStakingObject_1(t *testing.T){
	expected := 1
	retrieved := len(co.GetClients())

	if retrieved != expected {
		t.Errorf("Unexpected number of staking Clients : got %v want %v",
		retrieved, expected)
	}
}

//Adding another client and testing that there are two clients in the map
func TestStakingObject_2(t *testing.T){
	nodeName := "Oasis_Local_1"

	_, stakingClient, err := rpc.StakingClient(ws_url_Invalid)
	if err != nil {
		panic(err)
	}
	
	co.AddClient(nodeName,stakingClient)

	expected := 2
	retrieved := len(co.GetClients())

	//Check if the map contains two clients
	if retrieved != expected {
		t.Errorf("Unexpected number of staking Clients : got %v want %v",
		retrieved, expected)
	}
	clients := co.GetClients()

	if _, ok := clients[nodeName]; ok{
		
	}else{
		t.Errorf("Expected to find node with name %v but didn't find it.",
		nodeName)
	}
	
}