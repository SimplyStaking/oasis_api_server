package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	conf "github.com/SimplyVC/oasis_api_server/src/config"
	handler "github.com/SimplyVC/oasis_api_server/src/handlers"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// StartServer starts server by setting router and all endpoints
func StartServer() error {
	// Load prometheus configurations
	prometheusConf := conf.LoadPrometheusConfiguration()
	if prometheusConf == nil {
		lgr.Error.Println("Loading of Prometheus configuration has Failed!")
		// Abort Program no Port configured to run API on
		os.Exit(0)
	}

	// Load port configurations
	mainConf := conf.LoadMainConfiguration()
	if mainConf == nil {
		lgr.Error.Println("Loading of Port configuration has Failed!")
		// Abort Program no Port configured to run API on
		os.Exit(0)
	}
	// Load socket configuration but do not use them
	nodesConf := conf.LoadNodesConfiguration()
	if nodesConf == nil {
		lgr.Error.Println("Loading of Socket configuration has Failed!")
		// Abort Program no Sockets configured to run API on
		os.Exit(0)
	}

	apiPort := mainConf["api_server"]["port"]
	lgr.Info.Println("Loaded port : ", apiPort)

	// Router object to handle requests
	router := mux.NewRouter().StrictSlash(true)

	// Router Handlers to handle General API Calls
	router.HandleFunc("/api/ping/", handler.Pong).Queries("name", "{name}").Methods("Get")
	router.HandleFunc("/api/getconnectionslist", handler.GetConnections).Methods("Get")

	// Router Handlers to handle Consensus API Calls
	router.HandleFunc("/api/consensus/genesis/", handler.GetConsensusStateToGenesis).Methods("Get")
	router.HandleFunc("/api/consensus/epoch/", handler.GetEpoch).Methods("Get")
	router.HandleFunc("/api/consensus/block/", handler.GetBlock).Methods("Get")
	router.HandleFunc("/api/consensus/blockheader/", handler.GetBlockHeader).Methods("Get")
	router.HandleFunc("/api/consensus/blocklastcommit/", handler.GetBlockLastCommit).Methods("Get")
	router.HandleFunc("/api/consensus/transactions/", handler.GetTransactions).Methods("Get")
	router.HandleFunc("/api/pingnode/", handler.PingNode).Queries("name", "{name}").Methods("Get")

	// Router Handlers to handle Registry API Calls
	router.HandleFunc("/api/registry/entities/", handler.GetEntities).Methods("Get")
	router.HandleFunc("/api/registry/nodes/", handler.GetNodes).Methods("Get")
	router.HandleFunc("/api/registry/runtimes/", handler.GetRuntimes).Methods("Get")
	router.HandleFunc("/api/registry/genesis/", handler.GetRegistryStateToGenesis).Methods("Get")
	router.HandleFunc("/api/registry/entity/", handler.GetEntity).Methods("Get")
	router.HandleFunc("/api/registry/node/", handler.GetNode).Methods("Get")
	router.HandleFunc("/api/registry/runtime/", handler.GetRuntime).Methods("Get")

	// Router Handlers to handle Staking API Calls
	router.HandleFunc("/api/staking/totalsupply/", handler.GetTotalSupply).Methods("Get")
	router.HandleFunc("/api/staking/commonpool/", handler.GetCommonPool).Methods("Get")
	router.HandleFunc("/api/staking/genesis/", handler.GetStakingStateToGenesis).Methods("Get")
	router.HandleFunc("/api/staking/threshold/", handler.GetThreshold).Methods("Get")
	router.HandleFunc("/api/staking/accounts/", handler.GetAccounts).Methods("Get")
	router.HandleFunc("/api/staking/accountinfo/", handler.GetAccountInfo).Methods("Get")
	router.HandleFunc("/api/staking/delegations/", handler.GetDelegations).Methods("Get")
	router.HandleFunc("/api/staking/debondingdelegations/", handler.GetDebondingDelegations).Methods("Get")

	// Router Handlers to handle NodeController API Calls
	router.HandleFunc("/api/nodecontroller/synced/", handler.GetIsSynced).Methods("Get")

	// Router Handlers to handle Scheduler API Calls
	router.HandleFunc("/api/scheduler/validators/", handler.GetValidators).Methods("Get")
	router.HandleFunc("/api/scheduler/committees/", handler.GetCommittees).Methods("Get")
	router.HandleFunc("/api/scheduler/genesis/", handler.GetSchedulerStateToGenesis).Methods("Get")

	// Router Handlers to handle Prometheus API Calls
	router.HandleFunc("/api/prometheus/gauge/", handler.PrometheusQueryGauge).Methods("Get")
	router.HandleFunc("/api/prometheus/counter/", handler.PrometheusQueryCounter).Methods("Get")

	// Router Handlers to handle the Node_Extractor API Calls
	router.HandleFunc("/api/extractor/gauge/", handler.NodeExtractorQueryGauge).Methods("Get")
	router.HandleFunc("/api/extractor/counter/", handler.NodeExtractorQueryCounter).Methods("Get")

	// Router Handlers to handle System API Calls
	router.HandleFunc("/api/system/memory/", handler.GetMemory).Methods("Get")
	router.HandleFunc("/api/system/disk/", handler.GetDisk).Methods("Get")
	router.HandleFunc("/api/system/cpu/", handler.GetCPU).Methods("Get")
	router.HandleFunc("/api/system/network/", handler.GetNetwork).Methods("Get")

	log.Fatal(http.ListenAndServe(":"+apiPort, router))
	return nil
}
