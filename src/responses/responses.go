package responses

import (
	tmed "github.com/cometbft/cometbft/crypto"
	mint_types "github.com/cometbft/cometbft/types"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
	beacon_api "github.com/oasisprotocol/oasis-core/go/beacon/api"
	common_entity "github.com/oasisprotocol/oasis-core/go/common/entity"
	common_node "github.com/oasisprotocol/oasis-core/go/common/node"
	common_quantity "github.com/oasisprotocol/oasis-core/go/common/quantity"
	consensus_api "github.com/oasisprotocol/oasis-core/go/consensus/api"
	document_api "github.com/oasisprotocol/oasis-core/go/genesis/api"
	gen_api "github.com/oasisprotocol/oasis-core/go/genesis/api"
	registry_api "github.com/oasisprotocol/oasis-core/go/registry/api"
	scheduler_api "github.com/oasisprotocol/oasis-core/go/scheduler/api"
	sentry_api "github.com/oasisprotocol/oasis-core/go/sentry/api"
	staking_api "github.com/oasisprotocol/oasis-core/go/staking/api"
)

// StakingEvents responds with a list of events
type StakingEvents struct {
	StakingEvents []*staking_api.Event `json:"result"`
}

// TendermintAddress responds with a tendermint public key address
type TendermintAddress struct {
	TendermintAddress *tmed.Address `json:"result"`
}

// SentryResponse responds with network statistics of system
type SentryResponse struct {
	SentryAddresses *sentry_api.SentryAddresses `json:"result"`
}

// NetworkResponse responds with network statistics of system
type NetworkResponse struct {
	Network []network.Stats `json:"result"`
}

// CPUResponse responds with CPU statistics of system
type CPUResponse struct {
	CPU *cpu.Stats `json:"result"`
}

// DiskResponse responds with memory statistics of system
type DiskResponse struct {
	Disk []disk.Stats `json:"result"`
}

// MemoryResponse responds with memory statistics of system
type MemoryResponse struct {
	Memory *memory.Stats `json:"result"`
}

// SchedulerGenesisState responds with scheduler genesis state
type SchedulerGenesisState struct {
	SchedulerGenesisState *scheduler_api.Genesis `json:"result"`
}

// CommitteesResponse responds with Committees
type CommitteesResponse struct {
	Committee []*scheduler_api.Committee `json:"result"`
}

// ValidatorsResponse responds with Validators and their voting power
type ValidatorsResponse struct {
	Validators []*scheduler_api.Validator `json:"result"`
}

// IsSyncedResponse responds with boolean signifying synchronisation
// state of node
type IsSyncedResponse struct {
	Synced bool `json:"result"`
}

// DebondingDelegationsResponse responds with debonding delegations
// for specified public key
type DebondingDelegationsResponse struct {
	DebondingDelegations map[staking_api.Address][]*staking_api.DebondingDelegation `json:"result"`
}

// DelegationsResponse responds with delegations for public key
type DelegationsResponse struct {
	Delegations map[staking_api.Address]*staking_api.Delegation `json:"result"`
}

// AccountResponse responds with an account
type AccountResponse struct {
	Account *staking_api.Account `json:"result"`
}

// AllAddressesResponse responds with list of Accounts
type AllAddressesResponse struct {
	AllAddresses []staking_api.Address `json:"result"`
}

// AddressResponse responds with a staking address
type AddressResponse struct {
	Address staking_api.Address `json:"result"`
}

// StakingGenesisResponse responds with Staking Genesis File
type StakingGenesisResponse struct {
	GenesisStaking *staking_api.Genesis `json:"result"`
}

// QuantityResponse responds with quantity
type QuantityResponse struct {
	Quantity *common_quantity.Quantity `json:"result"`
}

// RegistryEntityResponse responds with details of single Entity
type RegistryEntityResponse struct {
	Entity *common_entity.Entity `json:"result"`
}

// RegistryNodeResponse responds with details of single Node
type RegistryNodeResponse struct {
	Node *common_node.Node `json:"result"`
}

// RegistryEventsResponse responds with events at specified block height.
type RegistryEventsResponse struct {
	Events []*registry_api.Event `json:"results"`
}

// NodeStatusResponse responds with a node's status.
type NodeStatusResponse struct {
	NodeStatus *registry_api.NodeStatus `json:"result"`
}

// RegistryGenesisResponse responds with genesis state of registry
type RegistryGenesisResponse struct {
	GenesisRegistry *registry_api.Genesis `json:"result"`
}

// NodelistResponse responds with NodeList
type NodelistResponse struct {
	NodeList *registry_api.NodeList `json:"result"`
}

// RuntimeResponse responds with single Runtime
type RuntimeResponse struct {
	Runtime *registry_api.Runtime `json:"result"`
}

// RuntimesResponse responds with Multiple runtimes
type RuntimesResponse struct {
	Runtimes []*registry_api.Runtime `json:"result"`
}

// NodesResponse responding with Multiple Nodes
type NodesResponse struct {
	Nodes []*common_node.Node `json:"result"`
}

// EntitiesResponse responding with Multiple entities
type EntitiesResponse struct {
	Entities []*common_entity.Entity `json:"result"`
}

// TransactionsResponse responds with all transactions in block
type TransactionsResponse struct {
	Transactions [][]byte `json:"result"`
}

// BlockHeaderResponse responds with Tendermint Header Type
type BlockHeaderResponse struct {
	BlkHeader *mint_types.Header `json:"result"`
}

// BlockLastCommitResponse responds with Tendermint Last Commit Type
type BlockLastCommitResponse struct {
	BlkLastCommit *mint_types.Commit `json:"result"`
}

// BlockResponse responds with custom Block response with an unmarshalled
// message
type BlockResponse struct {
	Blk *consensus_api.Block `json:"result"`
}

// EpochResponse responds with epcoh time
type EpochResponse struct {
	Ep beacon_api.EpochTime `json:"result"`
}

// ConsensusGenesisResponse with consensus Genesis Document
type ConsensusGenesisResponse struct {
	GenJSON *gen_api.Document `json:"result"`
}

// SuccessResponse with succeful result
type SuccessResponse struct {
	Result string `json:"result"`
}

// ErrorResponse responds with an error message that will be set
type ErrorResponse struct {
	Error string `json:"error"`
}

// ConnectionsResponse responds with all connections configured
type ConnectionsResponse struct {
	Results []string `json:"result"`
}

// ConsensusParametersResponse responds with the staking consensus parameters
type ConsensusParametersResponse struct {
	ConsensusParameters *staking_api.ConsensusParameters `json:"result"`
}

// StatusResponse responds with the current status overview
type StatusResponse struct {
	Status *consensus_api.Status `json:"result"`
}

// GenesisDocumentResponse reponds with the original genesis document.
type GenesisDocumentResponse struct {
	GenesisDocument *document_api.Document `json:"result"`
}

// SuccessResponsed Assinging Variable Responses that do not need to be changed.
var SuccessResponsed = SuccessResponse{Result: "pong"}
