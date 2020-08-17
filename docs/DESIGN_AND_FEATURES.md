# Design and Features of the Oasis API Server

This page will present the inner workings of the API Server as well as the features that one is able to interact with and how. The following points will be presented and discussed:

- [**Design**](#design)
- [**Complete List of Endpoints**](#complete-list-of-endpoints)
- [**Using the API**](#using-the-api)

## Design

The components involved in the API Server are the following:
- The **Oasis Nodes** from which the API retrieves information
- This **API Server** uses multiple Oasis API clients to retrieve data from the Oasis nodes.
- The **User/Program** which sends `GET Requests` to the `API Server` as per the defined endpoints, and receives JSON formatted responses from the `API Server`

The diagram below gives an idea of the various components at play when the API Server is running, and how they interact with each other and the user/program:

<img src="SERVER.png" alt="design"/>

The API Server works as follows:
- The API Server loads the configuration containing the internal socket information for each node from the `config/user_config_nodes.ini` file together with Prometheus endpoints that are used to query blockchain data.
- The API Server loads the API server configuration from the `config/user_config_main.ini` file together with the Node Exporter endpoint which will be used to query machine data.
- The API Server has an option to also retrieve the data of Sentries connected to the node through the External URl and tls certificate data of the Sentry. This data is set up in the `config/user_config_sentry` file.
- By communicating through this port, the API Server receives the endpoints specified in the `Complete List of Endpoints` section below, and requests information from the nodes it is connected to accordingly.
- Once a request is received for an endpoint the server will read the query which should contain the name of the node that will be queried, it then attempts to establish a connection to the node and request data from it. This data is then foramtted into JSON and returned.
- The server interacts with the protocol API through these clients :
    1. [Consensus Client](https://godoc.org/github.com/oasisprotocol/oasis-core/go/consensus/api#ClientBackend)
    2. [Registry Backend](https://godoc.org/github.com/oasisprotocol/oasis-core/go/registry/api#Backend)
    3. [Staking Backend](https://godoc.org/github.com/oasisprotocol/oasis-core/go/staking/api#Backend)
    4. [Scheduler Backend](https://godoc.org/github.com/oasisprotocol/oasis-core/go/scheduler/api#Backend)
    5. [NodeController](https://godoc.org/github.com/oasisprotocol/oasis-core/go/control/api#NodeController)
    6. [Sentry](https://godoc.org/github.com/oasisprotocol/oasis-core/go/sentry/api#Backend)

## Complete List of Endpoints

| API Endpoint                         | Required Inputs                 | Optional Inputs | Output                    | 
|--------------------------------------|---------------------------------|-----------------|---------------------------|
| /api/ping                            | none                            | none            | Pong                      | 
| /api/getconnectionslist              | none                            | none            | List of Connections       |
| /api/consensus/genesis               | Node Name                       | Height          | Consensus Genesis State   |
| /api/consensus/genesisdocument       | Node Name                       |                 | Original Genesis Document |
| /api/consensus/epoch                 | Node Name                       | Height          | Epoch                     |
| /api/consensus/status                | Node Name                       |                 | Node Status               | 
| /api/consensus/block                 | Node Name                       | Height          | Block Object              | 
| /api/consensus/blockheader           | Node Name                       | Height          | Block Header Object       | 
| /api/consensus/blocklastcommit       | Node Name                       | Height          | Block Last Commit Object  |
| /api/consensus/pubkeyaddress         | Consensus Public Key            | none            | Tendermint Key Address    |
| /api/consensus/transactions          | Node Name                       | Height          | List of Transactions      | 
| /api/pingnode                        | Node Name                       | None            | Pong                      | 
| /api/registry/entities               | Node Name                       | Height          | List of entities          | 
| /api/registry/nodes                  | Node Name                       | Height          | List of Nodes             | 
| /api/registry/runtimes               | Node Name                       | Height          | List of RunTimes          | 
| /api/registry/genesis                | Node Name                       | Height          | Genesis State of Registry | 
| /api/registry/entity                 | Node Name, Entity Public Key    | Height          | Entity                    | 
| /api/registry/node                   | Node Name, Node Public Key      | Height          | Node                      | 
| /api/registry/nodestatus             | Node Name, Node Public Key      | Height          | Node Status               | 
| /api/registry/events                 | Node Name                       | Height          | Registry Events           | 
| /api/registry/runtime                | Node Name, Runtime Namespace    | Height          | Runtime                   |
| /api/registry/runtimes               | Node Name, Suspended Boolean    | Height          | Runtimes                  | 
| /api/staking/totalsupply             | Node Name                       | Height          | Total Supply              | 
| /api/staking/commonpool              | Node Name                       | Height          | Common Pool               | 
| /api/staking/lastblockfees           | Node Name                       | Height          | Last Block Fees           |
| /api/staking/genesis                 | Node Name                       | Height          | Staking Genesis State     | 
| /api/staking/threshold               | Node Name, kind                 | Height          | Threshold                 | 
| /api/staking/addresses               | Node Name                       | Height          | List of accounts          |
| /api/staking/account                 | Node Name, Account Address      | Height          | Account information       | 
| /api/staking/delegations             | Node Name, Account Address      | Height          | Delegations               | 
| /api/staking/debondingdelegations    | Node Name, Account Address      | Height          | DebondingDelegations      |
| /api/staking/events                  | Node Name                       | Height          | List of Events            |
| /api/staking/publickeytoaddress      | Public Key                      |                 | Staking Address           |
| /api/nodecontroller/synced           | Node Name                       | None            | Synchronized State        | 
| /api/scheduler/validators            | Node Name                       | Height          | List of Validators        | 
| /api/scheduler/committees            | Node Name, Namespace            | Height          | Committees                | 
| /api/scheduler/genesis               | Node Name                       | Height          | Scheduler Genesis State   | 
| /api/prometheus/gauge                | Node Name, Gauge Name           | none            | Gauge Value               | 
| /api/prometheus/counter              | Node Name, Counter Name         | none            | Counter Value             | 
| /api/exporter/gauge                  | Gauge Name                      | none            | Gauge Value               | 
| /api/exporter/counter                | Counter Name                    | none            | Counter Value             | 
| /api/sentry/addresses                | Node Name                       | none            | Nodes Connected to Sentry |

## Example Queries

This is a list of example queries with the exact parameters and a custom node `Oasis_Main_Validator` being pinged from localhost `127.0.0.1`.

| API Endpoint                         | Query                                                                                                                                        |
|--------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------|
| /api/ping                            | 127.0.0.1:8686/api/ping                                                                                                                      | 
| /api/getconnectionslist              | 127.0.0.1:8686/api/getconnectionslist                                                                                                        |
| /api/consensus/genesis               | 127.0.0.1:8686/api/consensus/genesis?name=Oasis_Main_Validator&height=1000                                                                   |
| /api/consensus/genesisdocument       | 127.0.0.1:8686/api/consensus/genesisdocument?name=Oasis_Main_Validator&height=1000                                                           |
| /api/consensus/epoch                 | 127.0.0.1:8686/api/consensus/epoch?name=Oasis_Main_Validator&height=1000                                                                     |
| /api/consensus/status                | 127.0.0.1:8686/api/consensus/status?name=Oasis_Main_Validator&height=1000                                                                    | 
| /api/consensus/block                 | 127.0.0.1:8686/api/consensus/block?name=Oasis_Main_Validator&height=1000                                                                     |
| /api/consensus/blockheader           | 127.0.0.1:8686/api/consensus/blockheader?name=Oasis_Main_Validator&height=1000                                                               |
| /api/consensus/blocklastcommit       | 127.0.0.1:8686/api/consensus/blocklastcommit?name=Oasis_Main_Validator&height=1000                                                           |
| /api/consensus/pubkeyaddress         | 127.0.0.1:8686/api/consensus/pubkeyaddress?consensus_public_key=AzJTHgUZKYGYVPoN5F8WLtMyEPh7OKpM1uJGQVRiZek=                                 |
| /api/consensus/transactions          | 127.0.0.1:8686/api/consensus/transactions?name=Oasis_Main_Validator&height=1000                                                              |
| /api/pingnode                        | 127.0.0.1:8686/api/pingnode?name=Oasis_Main_Validator                                                                                        |
| /api/registry/entities               | 127.0.0.1:8686/api/registry/entities?name=Oasis_Main_Validator&height=1000                                                                   |
| /api/registry/nodes                  | 127.0.0.1:8686/api/registry/nodes?name=Oasis_Main_Validator&height=1000                                                                      |
| /api/registry/runtimes               | 127.0.0.1:8686/api/registry/runtimes?name=Oasis_Main_Validator&height=1000                                                                   |
| /api/registry/genesis                | 127.0.0.1:8686/api/registry/genesis?name=Oasis_Main_Validator&height=1000                                                                    |
| /api/registry/entity                 | 127.0.0.1:8686/api/registry/entity?name=Oasis_Main_Validator&height=1000&entity=gb8SHLeDc69Elk7OTfqhtVgE2sqxrBCDQI84xKR+Bjg=                 |
| /api/registry/node                   | 127.0.0.1:8686/api/registry/node?name=Oasis_Main_Validator&height=1000&nodeID=5RIMVgnsN1D/HdvNxXCpE+lWH5U/SGYUrYsvhsTMbyA=                   |
| /api/registry/nodestatus             | 127.0.0.1:8686/api/registry/nodestatus?name=Oasis_Main_Validator&height=1000&nodeID=5RIMVgnsN1D/HdvNxXCpE+lWH5U/SGYUrYsvhsTMbyA=             |
| /api/registry/events                 | 127.0.0.1:8686/api/registry/events?name=Oasis_Main_Validator&height=1000                                                                     |
| /api/registry/runtime                | 127.0.0.1:8686/api/registry/runtime?name=Oasis_Main_Validator&height=1000&namespace=6XJLXaerB2A/HdvNxXCpE+lWH5U/SGYUrXsvhsTMbyB=             |
| /api/registry/runtimes               | 127.0.0.1:8686/api/registry/runtimes?name=Oasis_Main_Validator&height=1000&suspended=true                                                    |
| /api/staking/totalsupply             | 127.0.0.1:8686/api/staking/totalsupply?name=Oasis_Main_Validator&height=1000                                                                 |
| /api/staking/commonpool              | 127.0.0.1:8686/api/staking/commonpool?name=Oasis_Main_Validator&height=1000                                                                  |
| /api/staking/lastblockfees           | 127.0.0.1:8686/api/staking/lastblockfees?name=Oasis_Main_Validator&height=1000                                                               |
| /api/staking/genesis                 | 127.0.0.1:8686/api/staking/genesis?name=Oasis_Main_Validator&height=1000                                                                     |
| /api/staking/threshold               | 127.0.0.1:8686/api/staking/threshold?name=Oasis_Main_Validator&height=1000&kind=1                                                            |
| /api/staking/addresses               | 127.0.0.1:8686/api/staking/addresses?name=Oasis_Main_Validator&height=1000                                                                   |
| /api/staking/account                 | 127.0.0.1:8686/api/staking/account?name=Oasis_Main_Validator&height=1000&address=oasis1qqqf342r78nz05dq2pa3wzh0w54k3ea49u6rqdhv              |
| /api/staking/delegations             | 127.0.0.1:8686/api/staking/delegations?name=Oasis_Main_Validator&height=1000&address=oasis1qqqf342r78nz05dq2pa3wzh0w54k3ea49u6rqdhv          |
| /api/staking/debondingdelegations    | 127.0.0.1:8686/api/staking/debondingdelegations?name=Oasis_Main_Validator&height=1000&address=oasis1qqqf342r78nz05dq2pa3wzh0w54k3ea49u6rqdhv |
| /api/staking/events                  | 127.0.0.1:8686/api/staking/events?name=Oasis_Main_Validator&height=1000                                                                      |
| /api/staking/publickeytoaddress      | 127.0.0.1:8686/api/staking/publickeytoaddress?pubKey=BKNMlGLov7tJZi4Gopeu0sXGxXWvg1uKDfY4wNY3WCM=                                            |
| /api/nodecontroller/synced           | 127.0.0.1:8686/api/nodecontroller/synced?name=Oasis_Main_Validator                                                                           |
| /api/scheduler/validators            | 127.0.0.1:8686/api/scheduler/validators?name=Oasis_Main_Validator&height=1000                                                                |
| /api/scheduler/committees            | 127.0.0.1:8686/api/scheduler/committees?name=Oasis_Main_Validator&height=1000&namespace=6XJLXaerB2A/HdvNxXCpE+lWH5U/SGYUrXsvhsTMbyB=         |
| /api/scheduler/genesis               | 127.0.0.1:8686/api/scheduler/genesis?name=Oasis_Main_Validator&height=1000                                                                   |
| /api/prometheus/gauge                | 127.0.0.1:8686/api/prometheus/gauge?name=Oasis_Main_Validator&gauge=go_goroutines                                                            |
| /api/prometheus/counter              | 127.0.0.1:8686/api/prometheus/counter?name=Oasis_Main_Validator&gauge=go_memstats_alloc_bytes_total                                          |
| /api/exporter/gauge                  | 127.0.0.1:8686/api/exporter/gauge?gauge=node_nf_conntrack_entries                                                                            |
| /api/exporter/counter                | 127.0.0.1:8686/api/exporter/counter?counter=node_timex_pps_calibration_total                                                                 |
| /api/sentry/addresses                | 127.0.0.1:8686/api/sentry/addresses?name=Oasis_Main_Validator                                                                                |

## Using the API

To use the API one can either go in the browser and type in the URL that has the IP address of your running server, for example : `http://127.0.0.1:8686/api/consensus/blockheader?name=Oasis_Main_Validator&height=1000` or in the command line they can use the `curl` command to query it, for example : `curl "127.0.0.1:8686/api/consensus/blockheader?name=Oasis_Main_Validator&height=1000"`.

[Back to API front page](../README.md)
