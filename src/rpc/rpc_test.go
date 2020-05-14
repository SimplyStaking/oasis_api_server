package rpc_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/SimplyVC/oasis_api_server/src/rpc"
)

var (
	ctx                  context.Context
	isocket_path         = "unix:/home/vvol/serverdir/node/internal.sock"
	isocket_path_Invalid = "unix:/home/vvol/serverdir/nodes/internal.sock"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	ctx = context.Background()
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	ctx = context.Background()

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

// Testing if Scheduler Client Connects
func TestSchedulerClient_Success(t *testing.T) {
	_, _, err := rpc.SchedulerClient(isocket_path)
	if err != nil {
		t.Errorf("Failed to create SchedulerClient for socket %v got %v",
			isocket_path, err)
	}
}

// Testing if Node Controller Client Connects
func TestNodeControllerClient_Success(t *testing.T) {
	_, _, err := rpc.NodeControllerClient(isocket_path)
	if err != nil {
		t.Errorf("Failed to create SchedulerClient for socket %v got %v",
			isocket_path, err)
	}
}

// Testing if Registry Client Connects
func TestRegistryClient_Success(t *testing.T) {
	_, _, err := rpc.RegistryClient(isocket_path)
	if err != nil {
		t.Errorf("Failed to create RegistryClient for socket %v got %v",
			isocket_path, err)
	}
}

// Testing if Registry Client Connects
func TestStakingClient_Success(t *testing.T) {
	_, _, err := rpc.StakingClient(isocket_path)
	if err != nil {
		t.Errorf("Failed to create StakingClient for socket %v got %v",
			isocket_path, err)
	}
}

// Testing if Registry Client Connects
func TestConsensusClient_Success(t *testing.T) {
	_, _, err := rpc.ConsensusClient(isocket_path)
	if err != nil {
		t.Errorf("Failed to create ConsensusClient for socket %v got %v",
			isocket_path, err)
	}
}

// Testing connection function
func TestConnect_Success(t *testing.T) {
	_, err := rpc.Connect(isocket_path)
	if err != nil {
		t.Errorf("Failed to create connection for socket %v got %v",
			isocket_path, err)
	}
}
