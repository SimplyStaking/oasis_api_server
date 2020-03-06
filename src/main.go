package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/claudetech/ini"
	
	rpc "github.com/SimplyVC/oasis_api_server/src/rpc"
	consensus_api "github.com/SimplyVC/oasis_api_server/src/consensus_api"
	registry_api "github.com/SimplyVC/oasis_api_server/src/registry_api"
	staking_api "github.com/SimplyVC/oasis_api_server/src/staking_api"
	response "github.com/SimplyVC/oasis_api_server/src/responses"
)

//Main Function handles all the possible API routes.
func main() {
	err := startServer()
	if err != nil{
		fmt.Println("Server Stopped")
	}
}

//Configuration loader to read the .ini file and return maps of it.
func loadConfig(PortFile string, SocketFile string) (map [string]map[string]string, map [string]map[string]string){
	//First open and parse the .ini file
	var portConf ini.Config
	var socketConf ini.Config

	//Decode and read the file containing the port information
	if err := ini.DecodeFile(PortFile, &portConf); err != nil {
		fmt.Println(err)
		return nil,nil
	}

	//Decode and read the file containing the internal socket information
	if err := ini.DecodeFile(SocketFile, &socketConf); err != nil {
		fmt.Println(err)
		return nil,nil
	}
	return portConf, socketConf
}

func startServer() error {
	PortFile := "../config/user_config_main.ini"
	SocketFile := "../config/user_config_nodes.ini"
	//Load the configuration and start create the necessary objects
	portConf, socketConf := loadConfig(PortFile , SocketFile)
	
	//Assign the API port retrieved from the .ini file
	api_port := portConf["api_server"]["port"]
	fmt.Println("Hosting Port at : ", api_port)

	//Return the created Oasis Objects
	co, _, _ := loadOasisAPIs(socketConf)

	// //Router object to handle the requests
	router := mux.NewRouter().StrictSlash(true)

	//Router Function Handlers to run general API calls
	router.HandleFunc("/api/pingApi", Pong).Methods("Get")

	//Router Handlers to handle the Consensus Calls
	// router.Path("/consensus/").Queries("name","{name}").HandlerFunc(co.Pong)
	router.HandleFunc("/api/pingNode/",co.Pong).Queries("name","{name}").Methods("Get")
	router.HandleFunc("/api/getConnectionsList", co.GetConnectionslist).Methods("Get")

	// //Router Function Handlers to handle the Registry Calls
	//router.HandleFunc("/api/registry/pingNode/", ro.Pong).Queries("name","{name}").Methods("Get")

	// //Router Function Handlers to handle the Staking Calls
	//router.HandleFunc("/api/staking/pingNode/", so.Pong).Queries("name","{name}").Methods("Get")

	log.Fatal(http.ListenAndServe(":"+api_port, router))

	return nil
}

//Load the Objects needed to start a server
func loadOasisAPIs(socketConf map [string]map[string]string)(consensus_api.ConsensusObject,  registry_api.RegistryObject, staking_api.StakingObject){
	//Create context to be used by the Oasis API
	ctx := context.Background()

	//Creating the Consensus/Registry/Staking Object which will run all the consensus API methods
	co := consensus_api.ConsensusObject{}
	ro := registry_api.RegistryObject{}
	so := staking_api.StakingObject{}

	//Setting Context only once for each api object
	co.SetContext(ctx)
	ro.SetContext(ctx)
	so.SetContext(ctx)

	//Retrieve all the possible local internal.sock the api needs to run with.
	//Create a client for each internal socket
	for i, socket := range socketConf{
		fmt.Println(socket)
		fmt.Println(socketConf[i])
		fmt.Println(socketConf[i]["node_name"])
		fmt.Println(socketConf[i]["ws_url"])
		//Creating a Consensus Client Object
		_, consensusClient, err := rpc.ConsensusClient("unix:"+socketConf[i]["ws_url"])
		if err != nil {
			panic(err)
		}
		co.AddClient(socketConf[i]["node_name"],consensusClient)

		//Creating a Registry Client Object
		_, registryClient, err := rpc.RegistryClient("unix:"+socketConf[i]["ws_url"])
		if err != nil {
			panic(err)
		}
		ro.AddClient(socketConf[i]["node_name"],registryClient)
		
		//Creating a Staking Client Object
		_, stakingClient, err := rpc.StakingClient("unix:"+socketConf[i]["ws_url"])
		if err != nil {
			panic(err)
		}
		so.AddClient(socketConf[i]["node_name"],stakingClient)
	}

	return co,ro,so
}


//General ping function to ping if the entire API is online
func Pong(w http.ResponseWriter, r *http.Request) {
	//Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")
	fmt.Println("Received request for /api/pingApi")
	json.NewEncoder(w).Encode(response.Responded_pong)
}