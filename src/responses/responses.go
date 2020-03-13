package responses

import (
	common_signature "github.com/oasislabs/oasis-core/go/common/crypto/signature"
	common_entity "github.com/oasislabs/oasis-core/go/common/entity"
	common_node "github.com/oasislabs/oasis-core/go/common/node"
	common_quantity "github.com/oasislabs/oasis-core/go/common/quantity"
	consensus_api "github.com/oasislabs/oasis-core/go/consensus/api"
	epoch_api "github.com/oasislabs/oasis-core/go/epochtime/api"
	gen_api "github.com/oasislabs/oasis-core/go/genesis/api"
	registry_api "github.com/oasislabs/oasis-core/go/registry/api"
	scheduler_api "github.com/oasislabs/oasis-core/go/scheduler/api"
	staking_api "github.com/oasislabs/oasis-core/go/staking/api"
	mint_types "github.com/tendermint/tendermint/types"
)

//SchedulerGenesisState responds with the scheduler genesis state
type SchedulerGenesisState struct {
	SchedulerGenesisState *scheduler_api.Genesis `json:"gensis"`
}

//CommitteesResponse responds with Committees
type CommitteesResponse struct {
	Committee []*scheduler_api.Committee `json:"commitees"`
}

//ValidatorsResponse responds with Validators and their voting power
type ValidatorsResponse struct {
	Validators []*scheduler_api.Validator `json:"validators"`
}

//IsSyncedResponse responds with the IsSynced State
type IsSyncedResponse struct {
	Synced bool `json:"issynced"`
}

//DebondingDelegationsResponse responds with debonding delegations for a public key
type DebondingDelegationsResponse struct {
	DebondingDelegations map[common_signature.PublicKey][]*staking_api.DebondingDelegation `json:"debondig_delegations"`
}

//DelegationsResponse responds with delegations for a public key
type DelegationsResponse struct {
	Delegations map[common_signature.PublicKey]*staking_api.Delegation `json:"delegations"`
}

//AccountResponse responds with an account
type AccountResponse struct {
	AccountInfo *staking_api.Account `json:"account"`
}

//AllAccountsResponse responds with a list of Accounts
type AllAccountsResponse struct {
	AllAccounts []common_signature.PublicKey `json:"accounts"`
}

//StakingGenesisResponse responds with a Staking Genesis File
type StakingGenesisResponse struct {
	GenesisStaking *staking_api.Genesis `json:"genesis"`
}

//QuantityResponse responds with a quantity
type QuantityResponse struct {
	Quantity *common_quantity.Quantity `json:"quantity"`
}

//RegistryEntityResponse responds with the details of a single Entity
type RegistryEntityResponse struct {
	Entity *common_entity.Entity `json:"entity"`
}

//RegistryNodeResponse responds with the details of a single Node
type RegistryNodeResponse struct {
	Node *common_node.Node `json:"node"`
}

//RegistryGenesisResponse responds with the genesis state of the registry
type RegistryGenesisResponse struct {
	GenesisRegistry *registry_api.Genesis `json:"genesis"`
}

//NodelistResponse responds with a NodeList
type NodelistResponse struct {
	NodeList *registry_api.NodeList `json:"nodelist"`
}

//RuntimeResponse responds with a single Runtime
type RuntimeResponse struct {
	Runtime *registry_api.Runtime `json:"runtime"`
}

//RuntimesResponse responds with Multiple Runtimes
type RuntimesResponse struct {
	Runtimes []*registry_api.Runtime `json:"runtimes"`
}

//NodesResponse responding with Multiple Nodes
type NodesResponse struct {
	Nodes []*common_node.Node `json:"nodes"`
}

//EntitiesResponse responding with Multiple Entities
type EntitiesResponse struct {
	Entities []*common_entity.Entity `json:"entities"`
}

//TransactionsResponse responds with all the transactions in a block
type TransactionsResponse struct {
	Transactions [][]byte `json:"transactions"`
}

//BlockHeaderResponse responds with a Tendermint Header Type
type BlockHeaderResponse struct {
	BlkHeader *mint_types.Header `json:"block_header"`
}

//BlockLastCommitResponse responds with a Tendermint Last Commit Type
type BlockLastCommitResponse struct {
	BlkLastCommit *mint_types.Commit `json:"block_last_commit"`
}

//BlockResponse responds with a custom Block Response with an unmarshalled message
type BlockResponse struct {
	Blk *consensus_api.Block `json:"block"`
}

//EpochResponse responds with epcoh time
type EpochResponse struct {
	Ep epoch_api.EpochTime `json:"epoch"`
}

//ConsensusGenesisResponse  with the consensus Genesis Document
type ConsensusGenesisResponse struct {
	GenJSON *gen_api.Document `json:"genesis"`
}

//PongResponse responding to pong requests
type PongResponse struct {
	Result string `json:"result"`
}

//ErrorResponse repsonds with an error message that will be set
type ErrorResponse struct {
	Error string `json:"error"`
}

//ConnectionsResponse responds with all the connections configured
type ConnectionsResponse struct {
	Results []string `json:"result"`
}

//PongResponsed Assinging Variable Responses that do not need to be changed.
var PongResponsed = PongResponse{Result: "pong"}
