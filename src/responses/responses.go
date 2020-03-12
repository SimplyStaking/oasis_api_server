package responses

import (
	gen_api "github.com/oasislabs/oasis-core/go/genesis/api"
	epoch_api "github.com/oasislabs/oasis-core/go/epochtime/api"
	consensus_api "github.com/oasislabs/oasis-core/go/consensus/api"
	registry_api "github.com/oasislabs/oasis-core/go/registry/api"
	staking_api "github.com/oasislabs/oasis-core/go/staking/api"
	scheduler_api "github.com/oasislabs/oasis-core/go/scheduler/api"
	mint_types "github.com/tendermint/tendermint/types"
	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	common_entity "github.com/oasislabs/oasis-core/go/common/entity"
	common_node "github.com/oasislabs/oasis-core/go/common/node"
	common_quantity "github.com/oasislabs/oasis-core/go/common/quantity"
)

//Responding with the scheduler genesis state
type Response_SchedulerGenesisState struct{
	SchedulerGenesisState *scheduler_api.Genesis `json:"gensis"`
}

//Responding with Committees
type Response_Committees struct{
	Committee []*scheduler_api.Committee `json:"commitees"`
}

//Responding with Validators and their voting power
type Response_Validators struct{
	Validators []*scheduler_api.Validator `json:"validators"`
}

//Responding with the IsSynced State
type Response_IsSynced struct{
	Synced bool `json:"issynced"`
}

//Responding with debonding delegations for a public key
type Response_DebondingDelegations struct{
	DebondingDelegations map[common_signature.PublicKey][] *staking_api.DebondingDelegation `json:"debondig_delegations"`
}

//Responding with delegations for a public key
type Response_delegations struct{
	Delegations map[common_signature.PublicKey] *staking_api.Delegation `json:"delegations"`
}

//Responding with an account
type Response_account struct{
	AccountInfo *staking_api.Account `json:"account"`
}

//Responding with a list of Accounts
type Response_allAccounts struct{
	AllAccounts []common_signature.PublicKey `json:"accounts"`
}

//Responding with a Staking Genesis File
type Response_stakingGenesis struct{
	GenesisStaking *staking_api.Genesis `json:"genesis"`
}

//Responding with a quantity
type Response_quantity struct{
	Quantity *common_quantity.Quantity  `json:"quantity"`
}

//Respond with the details of a single Entity
type Response_registryEntity struct{
	Entity *common_entity.Entity  `json:"entity"`
}

//Respond with the details of a single Node
type Response_registryNode struct{
	Node *common_node.Node  `json:"node"`
}

//Respond with the genesis state of the registry
type Response_registryGenesis struct {
	GenesisRegistry *registry_api.Genesis `json:"genesis"`
}

//Respond with a NodeList
type Response_nodelist struct{
	NodeList *registry_api.NodeList  `json:"nodelist"`
}

//Respondig with a single Runtime
type Response_runtime struct{
	Runtime *registry_api.Runtime  `json:"runtime"`
}

//Respondig with Multiple Runtimes
type Response_runtimes struct{
	Runtimes []*registry_api.Runtime  `json:"runtimes"`
}

//Respondig with Multiple Nodes
type Response_nodes struct{
	Nodes []*common_node.Node  `json:"nodes"`
}

//Respondig with Multiple Entities
type Response_entities struct{
	Entities []*common_entity.Entity  `json:"entities"`
}

//Responding with all the transactions in a block
type Response_transactions struct{
	Transactions [][]byte  `json:"transactions"`
}

//Responding with a Tendermint Header Type
type Response_blockHeader struct{
	BlkHeader *mint_types.Header  `json:"block_header"`
}

//Responding with a Tendermint Last Commit Type
type Response_blockLastCommit struct{
	BlkLastCommit *mint_types.Commit  `json:"block_last_commit"`
}

//Responding with a custom Block Response with an unmarshalled message
type Response_block struct{
	Blk *consensus_api.Block `json:"block"`
}

//Responding with epcoh time
type Response_epoch struct{
	Ep epoch_api.EpochTime `json:"epoch"`
}

//Responding with a Genesis File
type Response_consensusgenesis struct{
	GenJSON *gen_api.Document `json:"genesis"`
}

//Responding to Pong Requests
type Response_pong struct{
	Result string `json:"result"`
}

//Responding to Pong Requests with an error
type Response_error struct{
	Error string `json:"error"`
}

type Response_Conns struct{
	Results []string `json:"result"`
}

//Assinging Variable Responses that do not need to be changed.
var Responded_pong = Response_pong{"pong"}