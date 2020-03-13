package rpc

import (
	"fmt"

	"google.golang.org/grpc"

	cmnGrpc "github.com/oasislabs/oasis-core/go/common/grpc"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
	control "github.com/oasislabs/oasis-core/go/control/api"
	registry "github.com/oasislabs/oasis-core/go/registry/api"
	scheduler "github.com/oasislabs/oasis-core/go/scheduler/api"
	staking "github.com/oasislabs/oasis-core/go/staking/api"
	storage "github.com/oasislabs/oasis-core/go/storage/api"
)

//StorageClient - initiate a new storage client
func StorageClient(address string) (*grpc.ClientConn, storage.Backend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to establish Storage Client Connection with node %s", address)
	}

	client := storage.NewStorageClient(conn)
	return conn, client, nil
}

//SchedulerClient - initiate a new scheduler client
func SchedulerClient(address string) (*grpc.ClientConn, scheduler.Backend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to establish Scheduler Client Connection with node %s", address)
	}

	client := scheduler.NewSchedulerClient(conn)
	return conn, client, nil
}

//NodeControllerClient - initiate a new registry client
func NodeControllerClient(address string) (*grpc.ClientConn, control.NodeController, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to establish NodeController Client Connection with node %s", address)
	}

	client := control.NewNodeControllerClient(conn)
	return conn, client, nil
}

//RegistryClient - initiate a new registry client
func RegistryClient(address string) (*grpc.ClientConn, registry.Backend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to establish Registry Client Connection with node %s", address)
	}

	client := registry.NewRegistryClient(conn)
	return conn, client, nil
}

// ConsensusClient - initiate a new consensus client
func ConsensusClient(address string) (*grpc.ClientConn, consensus.ClientBackend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to establish connection with node %s", address)
	}

	client := consensus.NewConsensusClient(conn)
	return conn, client, nil
}

// StakingClient - initiate a new staking client
func StakingClient(address string) (*grpc.ClientConn, staking.Backend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to establish connection with node %s", address)
	}

	client := staking.NewStakingClient(conn)
	return conn, client, nil
}

// Connect - connect to grpc
//Add a grpc.WithBlock() and grpc.WithTimeout() to have the dial to constantly try and establish a connection
func Connect(address string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.WaitForReady(false)))

	conn, err := cmnGrpc.Dial(
		address,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
