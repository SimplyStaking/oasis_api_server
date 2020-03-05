package registry_api_test

import (
	"context"
	"testing"
	"fmt"
	"os"

	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	registry_api "github.com/SimplyVC/oasis_api_server/src/registry_api"
)

var (
	ctx context.Context
	co  registry_api.RegistryObject
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
 
func setup() {
	ctx = context.Background()
	co = registry_api.RegistryObject{}

	co.SetContext(ctx)

	_, registryClient, err := rpc.RegistryClient("unix:/home/vvol/serverdir/node/internal.sock")
	if err != nil {
		panic(err)
	}
	
	co.AddClient("Oasis_Local",registryClient)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}
 
func teardown() {

	ctx = context.Background()
	co = registry_api.RegistryObject{}

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

//Testing if the setup indeed added one client
func TestRegistryObject_1(t *testing.T){
	expected := 1
	retrieved := len(co.GetClients())

	if retrieved != expected {
		t.Errorf("Unexpected number of registry Clients : got %v want %v",
		retrieved, expected)
	}
}

//Adding another client and testing that there are two clients in the map
func TestRegistryObject_2(t *testing.T){
	nodeName := "Oasis_Local_1"

	_, registryClient, err := rpc.RegistryClient("unix:/home/vvol/serverdir/node2/internal.sock")
	if err != nil {
		panic(err)
	}
	
	co.AddClient(nodeName,registryClient)

	expected := 2
	retrieved := len(co.GetClients())

	//Check if the map contains two clients
	if retrieved != expected {
		t.Errorf("Unexpected number of registry Clients : got %v want %v",
		retrieved, expected)
	}
	clients := co.GetClients()

	if _, ok := clients[nodeName]; ok{
		
	}else{
		t.Errorf("Expected to find node with name %v but didn't find it.",
		nodeName)
	}
	
}