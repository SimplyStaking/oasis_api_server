package api_router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	conf "github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	handler "github.com/SimplyVC/oasis_api_server/src/handlers"
)

func StartServer() error {
	//Load the port configurations
	portConf := conf.LoadPortConfiguration()
	if portConf == nil{
		lgr.Error.Println("Loading of Port configuration has Failed!")
		//Abort Program no Port configured to run the API on
		os.Exit(0)
	}
	//Load the socket configuration but do not use them
	socketConf:= conf.LoadSocketConfiguration()
	if socketConf == nil{
		lgr.Error.Println("Loading of Socket configuration has Failed!")
		//Abort Program no Sockets configured to run the API on
		os.Exit(0)
	}

	apiPort := portConf["api_server"]["port"]
	lgr.Info.Println("Loaded port : ", apiPort)

	//Router object to handle the requests
	router := mux.NewRouter().StrictSlash(true)

	//Router Handlers to handle the General API Calls
	router.HandleFunc("/api/ping/", handler.Pong).Queries("name","{name}").Methods("Get")
	router.HandleFunc("/api/getConnectionsList", handler.GetConnections).Methods("Get")

	//Router Handlers to handle the Consensus API Calls
	router.HandleFunc("/api/GetConsensusStateToGenesis/", handler.GetConsensusStateToGenesis).Methods("Get")
	router.HandleFunc("/api/GetEpoch/", handler.GetEpoch).Methods("Get")
	router.HandleFunc("/api/GetBlock/", handler.GetBlock).Methods("Get")
	router.HandleFunc("/api/GetBlockHeader/", handler.GetBlockHeader).Methods("Get")
	router.HandleFunc("/api/GetBlockLastCommit/", handler.GetBlockLastCommit).Methods("Get")
	router.HandleFunc("/api/GetTransactions/", handler.GetTransactions).Methods("Get")

	//Router Handlers to handle the Registry API Calls
	router.HandleFunc("/api/GetEntities/", handler.GetEntities).Methods("Get")
	router.HandleFunc("/api/GetNodes/", handler.GetNodes).Methods("Get")
	router.HandleFunc("/api/GetRuntimes/", handler.GetRuntimes).Methods("Get")
	router.HandleFunc("/api/GetNodeList/", handler.GetNodeList).Methods("Get")
	router.HandleFunc("/api/GetRegistryStateToGenesis/", handler.GetRegistryStateToGenesis).Methods("Get")
	router.HandleFunc("/api/GetEntity/", handler.GetEntity).Methods("Get")
	router.HandleFunc("/api/GetNode/", handler.GetNode).Methods("Get")
	router.HandleFunc("/api/GetRuntime/", handler.GetRuntime).Methods("Get")

	//Router Handlers to handle the Staking API Calls
	router.HandleFunc("/api/GetTotalSupply/", handler.GetTotalSupply).Methods("Get")
	router.HandleFunc("/api/GetCommonPool/", handler.GetCommonPool).Methods("Get")
	router.HandleFunc("/api/GetStakingStateToGenesis/", handler.GetStakingStateToGenesis).Methods("Get")
	router.HandleFunc("/api/GetThreshold/", handler.GetThreshold).Methods("Get")
	router.HandleFunc("/api/GetAccounts/", handler.GetAccounts).Methods("Get")
	router.HandleFunc("/api/GetAccountInfo/", handler.GetAccountInfo).Methods("Get")
	router.HandleFunc("/api/GetDelegations/", handler.GetDelegations).Methods("Get")
	router.HandleFunc("/api/GetDebondingDelegations/", handler.GetDebondingDelegations).Methods("Get")

	//Router Handlers to handle the NodeController API Calls
	router.HandleFunc("/api/GetIsSynced/", handler.GetIsSynced).Methods("Get")

	//Router Handlers to handle the Scheduler API Calls
	router.HandleFunc("/api/GetValidators/", handler.GetValidators).Methods("Get")
	router.HandleFunc("/api/GetCommittees/", handler.GetCommittees).Methods("Get")
	router.HandleFunc("/api/GetSchedulerStateToGenesis/", handler.GetSchedulerStateToGenesis).Methods("Get")

	log.Fatal(http.ListenAndServe(":"+apiPort, router))
	return nil
}