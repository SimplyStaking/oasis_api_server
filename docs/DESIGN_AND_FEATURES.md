# Design and Features of the Oasis API Server

This page will present the inner workings of the API Server as well as the features that one is able to interact with and how. The following points will be presented and discussed:

- [**Design**](#design)
- [**Complete List of Endpoints**](#complete-list-of-endpoints)
- [**Using the API**](#using-the-api)

## Design

The components involved in the API Server are the following:
- The **Oasis Nodes** from which the API retrieves information
- This **API Server** uses the multiple Oasis API Clients to retrieve data from the Oasis Nodes. 
- The **User/Program** which sends `GET Requests` to the `API Server` as per the defined Endpoints, and receives JSON formatted responses from the `API Server`

The diagram below gives an idea of the various components at play when the API Server is running, and how they interact with each other and the user/program:

<img src="SERVER.png" alt="design"/>

The API Server works as follows:
- The API Server loads the configuration containing the internal socket information for each node from the `config/user_config_nodes.ini` file together with the port at which the API will expose its endpoints taken from the file `config/user_config_main.ini`.
- The API Server also loads the information containing Prometheus endpoints from the file `config/prometheus_config_main.ini`.
- By communicating through this port, the API Server receives the endpoints specified in the `Complete List of Endpoints` section below, and requests information from the nodes it is connected to accordingly.
- Once a request is received for an endpoint the server will read the query which should contain the name of the node that will be queried, it then attempts to establish a connection to the node and request data from it. This data is then foramtted into JSON and returned.
- The server interacts with the protocol API through these clients :
1 : [Consensus Client](https://godoc.org/github.com/oasislabs/oasis-core/go/consensus/api#ClientBackend)
2 : [Registry Backend](https://godoc.org/github.com/oasislabs/oasis-core/go/registry/api#Backend)
3 : [Staking Backend](https://godoc.org/github.com/oasislabs/oasis-core/go/staking/api#Backend)
4 : [Scheduler Backend](https://godoc.org/github.com/oasislabs/oasis-core/go/scheduler/api#Backend)
5 : [NodeController](https://godoc.org/github.com/oasislabs/oasis-core/go/control/api#NodeController)

## Complete List of Endpoints
| API Endpoint                     | Required Inputs                 | Optional Inputs | Output                    | Description                                                                         |
|----------------------------------|---------------------------------|-----------------|---------------------------|-------------------------------------------------------------------------------------|
| /api/ping/                       | none                            | none            | Pong                      | Returns Pong if the API is online.                                                  |
| /api/getConnectionsList          | none                            | none            | List of Connections       | Returns a list of connections specified in the configuration files.                |
| /api/GetConsensusStateToGenesis/ | Node Name                       | Height          | Consensus Genesis State   | Returns the Genesis State of Consensus at a specific block height for a given node. |
| /api/GetEpoch/                   | Node Name                       | Height          | Epoch                     | Returns Epoch at a given block height                                               |
| /api/GetBlock/                   | Node Name                       | Height          | Block Object              | Returns Block object containing height and meta data.                               |
| /api/GetBlockHeader/             | Node Name                       | Height          | Block Header Object       | Returns the Block Header and all of it's data.                                      |
| /api/GetBlockLastCommit/         | Node Name                       | Height          | Block Last Commit Object  | Returns the Block Last Commit and all of it's data.                                 |
| /api/GetTransactions/            | Node Name                       | Height          | List of Transactions      | Returns the list of all transactions in a given block.                              |
| /api/pingNode/                   | Node Name                       | None            | Pong                      | Returns Pong is a connection with a given node can be established.                  |
| /api/GetEntities/                | Node Name                       | Height          | List of Entities          | Returns a list of entities at a given block height and their data.                  |
| /api/GetNodes/                   | Node Name                       | Height          | List of Nodes             | Returns a list of nodes at a given block height.                                    |
| /api/GetRuntimes/                | Node Name                       | Height          | List of RunTimes          | Returns a list of runtimes at a given block height.                                 |
| /api/GetRegistryStateToGenesis/  | Node Name                       | Height          | Genesis State of Registry | Returns the genesis state of the registry at a given block height.                  |
| /api/GetEntity/                  | Node Name, Entity Public Key    | Height          | Entity                    | Returns an Entity Object by it's public key a given block height.                   |
| /api/GetNode/                    |  Node Name, Node Public Key     | Height          | Node                      | Returns a Node Object by it's public key at a given block height.                   |
| /api/GetRuntime/                 |  Node Name, Runtime Namespace   | Height          | Runtime                   | Returns a Runtime object by it's namespace at a given block height.                 |
| /api/GetTotalSupply/             | Node Name                       | Height          | Total Supply              | Returns the total supply of tokens at a given block height.                         |
| /api/GetCommonPool/              | Node Name                       | Height          | Common Pool               | Returns the common pool of tokens at  a given block height.                         |
| /api/GetStakingStateToGenesis/   | Node Name                       | Height          | Staking Genesis State     | Returns the Staking Genesis state at a given block height.                          |
| /api/GetThreshold/               |  Node Name, kind                | Height          | Threshold                 |  Returns the Threshold of a specific kind at a given block height.                  |
| /api/GetAccounts/                | Node Name                       | Height          | List of accounts          | Returns the list of accounts and all their details.                                 |
| /api/GetAccountInfo/             |  Node Name, Account Public Key  | Height          | Account information       | Returns the account details using it's public key.                                  |
| /api/GetDelegations/             |  Node Name, Account Public Key  | Height          | Delegations               | Returns the delegations of a given account at a specific block height.              |
| /api/GetDebondingDelegations/    |  Node Name, Account Public Key  | Height          | DebondingDelegations      | Returns the debonding delegations of an account at a given block height.            |
| /api/GetIsSynced/                | Node Name                       | None            | Synchronized State        | Returns True/False depending on if the node is synchronized or not.                 |
| /api/GetValidators/              | Node Name                       | Height          | List of Validators        | Returns a list of validators and their voting power a specific block height.        |
| /api/GetCommittees/              | Node Name, Namespace            | Height          | Committees                | Returns committees under a namespace at a given block height.                       |
| /api/GetSchedulerStateToGenesis/ | Node Name                       | Height          | Scheduler Genesis State   | Returns a scheduler genesis state at a given block height.                          |
| /api/prometheus/gauge/           |  Node Name, gauge name          | none            | Gauge Value               | Returns a gauge value for a given gauge name.                                       |
| /api/prometheus/counter/         |  Node Name, counter name        | none            | Counter Value             | Returns a counter value for a given counter name.                                   |

## Using the API

For example, the endpoint `/api/GetIsSynced` can be called as follows: `http://localhost:8880/api/GetIsSynced?name=Oasis_Local`.
If successful, this will return:
```json
{
    "issynced": true
}
```

If an API connection for the node specified in the `websocket` field is not online, this will return:
```json
{
    "error": "Failed to get IsSynced!"
}
```

[Back to API front page](../README.md)
