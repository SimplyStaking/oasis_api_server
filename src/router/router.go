package router

import (
	"log"
	"os"

	"github.com/gorilla/mux"

	conf "github.com/SimplyVC/oasis_api_server/src/config"
	handler "github.com/SimplyVC/oasis_api_server/src/handlers"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/zenazn/goji/graceful"
)

// StartServer starts server by setting router and all endpoints
func StartServer() error {

	// Load port configurations
	mainConf, err2 := conf.LoadMainConfiguration()
	if err2 != nil {
		lgr.Error.Println("Loading of Port configuration has failed!")
		// Abort Program no Port configured to run API on
		os.Exit(0)
	}

	// Load socket configuration but do not use them
	_, err3 := conf.LoadNodesConfiguration()
	if err3 != nil {
		lgr.Error.Println("Loading of Socket configuration has failed!")
		// Abort Program no Sockets configured to run API on
		os.Exit(0)
	}

	// Load sentry configuration
	_, err4 := conf.LoadSentryConfiguration()
	if err4 != nil {
		lgr.Error.Println("Loading of Sentry configuration has failed!")
	}

	apiPort := mainConf["api_server"]["port"]
	lgr.Info.Println("Loaded port : ", apiPort)

	// Router object to handle requests
	router := mux.NewRouter().StrictSlash(true)

	// Router Handlers to handle General API Calls
	router.HandleFunc("/api/ping/", handler.Pong).Methods("Get")
	router.HandleFunc("/api/getconnectionslist",
		handler.GetConnections).Methods("Get")

	// Router Handlers to handle Consensus API Calls
	router.HandleFunc("/api/consensus/genesis/",
		handler.GetConsensusStateToGenesis).Methods("Get")
	router.HandleFunc("/api/consensus/epoch/",
		handler.GetEpoch).Methods("Get")
	router.HandleFunc("/api/consensus/block/",
		handler.GetBlock).Methods("Get")
	router.HandleFunc("/api/consensus/blockheader/",
		handler.GetBlockHeader).Methods("Get")
	router.HandleFunc("/api/consensus/blocklastcommit/",
		handler.GetBlockLastCommit).Methods("Get")
	router.HandleFunc("/api/consensus/pubkeyaddress/",
		handler.PublicKeyToAddress).Methods("Get")
	router.HandleFunc("/api/consensus/transactions/",
		handler.GetTransactions).Methods("Get")
	router.HandleFunc("/api/pingnode/",
		handler.PingNode).Methods("Get")

	// Router Handlers to handle Registry API Calls
	router.HandleFunc("/api/registry/entities/",
		handler.GetEntities).Methods("Get")
	router.HandleFunc("/api/registry/nodes/",
		handler.GetNodes).Methods("Get")
	router.HandleFunc("/api/registry/runtimes/",
		handler.GetRuntimes).Methods("Get")
	router.HandleFunc("/api/registry/genesis/",
		handler.GetRegistryStateToGenesis).Methods("Get")
	router.HandleFunc("/api/registry/entity/",
		handler.GetEntity).Methods("Get")
	router.HandleFunc("/api/registry/node/",
		handler.GetNode).Methods("Get")
	router.HandleFunc("/api/registry/runtime/",
		handler.GetRuntime).Methods("Get")

	// Router Handlers to handle Staking API Calls
	router.HandleFunc("/api/staking/totalsupply/",
		handler.GetTotalSupply).Methods("Get")
	router.HandleFunc("/api/staking/commonpool/",
		handler.GetCommonPool).Methods("Get")
	router.HandleFunc("/api/staking/genesis/",
		handler.GetStakingStateToGenesis).Methods("Get")
	router.HandleFunc("/api/staking/threshold/",
		handler.GetThreshold).Methods("Get")
	router.HandleFunc("/api/staking/accounts/",
		handler.GetAccounts).Methods("Get")
	router.HandleFunc("/api/staking/accountinfo/",
		handler.GetAccountInfo).Methods("Get")
	router.HandleFunc("/api/staking/delegations/",
		handler.GetDelegations).Methods("Get")
	router.HandleFunc("/api/staking/debondingdelegations/",
		handler.GetDebondingDelegations).Methods("Get")
	router.HandleFunc("/api/staking/events/",
		handler.GetEvents).Methods("Get")

	// Router Handlers to handle NodeController API Calls
	router.HandleFunc("/api/nodecontroller/synced/",
		handler.GetIsSynced).Methods("Get")

	// Router Handlers to handle Scheduler API Calls
	router.HandleFunc("/api/scheduler/validators/",
		handler.GetValidators).Methods("Get")
	router.HandleFunc("/api/scheduler/committees/",
		handler.GetCommittees).Methods("Get")
	router.HandleFunc("/api/scheduler/genesis/",
		handler.GetSchedulerStateToGenesis).Methods("Get")

	// Router Handlers to handle Prometheus API Calls
	router.HandleFunc("/api/prometheus/gauge/",
		handler.PrometheusQueryGauge).Methods("Get")
	router.HandleFunc("/api/prometheus/counter/",
		handler.PrometheusQueryCounter).Methods("Get")

	// Router Handlers to handle the Node Exporter API Calls
	router.HandleFunc("/api/exporter/gauge/",
		handler.NodeExporterQueryGauge).Methods("Get")
	router.HandleFunc("/api/exporter/counter/",
		handler.NodeExporterQueryCounter).Methods("Get")

	// Router Handlers to handle Sentry API Calls
	router.HandleFunc("/api/sentry/addresses/",
		handler.GetSentryAddresses).Methods("Get")

	log.Fatal(graceful.ListenAndServe(":"+apiPort, router))
	return nil
}
