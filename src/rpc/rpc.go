package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	cmnGrpc "github.com/oasislabs/oasis-core/go/common/grpc"
	"github.com/oasislabs/oasis-core/go/common/identity"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
	control "github.com/oasislabs/oasis-core/go/control/api"
	registry "github.com/oasislabs/oasis-core/go/registry/api"
	scheduler "github.com/oasislabs/oasis-core/go/scheduler/api"
	sentry "github.com/oasislabs/oasis-core/go/sentry/api"
	staking "github.com/oasislabs/oasis-core/go/staking/api"
)

// SentryClient - initiate new sentry client
func SentryClient(address string, tlsPath string) (*grpc.ClientConn,
	sentry.Backend, error) {

	conn, err := ConnectTLS(address, tlsPath)
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf(
			"Failed to establish Sentry Connection with node %s", address)
	}

	client := sentry.NewSentryClient(conn)
	return conn, client, nil
}

// SchedulerClient - initiate new scheduler client
func SchedulerClient(address string) (*grpc.ClientConn, scheduler.Backend, 
		error) {

	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to establish Scheduler Client "+
		"Connection with node %s", address)
	}

	client := scheduler.NewSchedulerClient(conn)
	return conn, client, nil
}

// NodeControllerClient - initiate new registry client
func NodeControllerClient(address string) (*grpc.ClientConn, 
		control.NodeController, error) {

	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to establish NodeController Client"+
		" Connection with node %s", address)
	}

	client := control.NewNodeControllerClient(conn)
	return conn, client, nil
}

// RegistryClient - initiate new registry client
func RegistryClient(address string) (*grpc.ClientConn, 
		registry.Backend, error) {

	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to establish Registry Client "+
		"Connection with node %s", address)
	}

	client := registry.NewRegistryClient(conn)
	return conn, client, nil
}

// ConsensusClient - initiate new consensus client
func ConsensusClient(address string) (*grpc.ClientConn, 
		consensus.ClientBackend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to establish connection with "+
		"node %s", address)
	}

	client := consensus.NewConsensusClient(conn)
	return conn, client, nil
}

// StakingClient - initiate new staking client
func StakingClient(address string) (*grpc.ClientConn, staking.Backend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to establish connection with " +
		"node %s", address)
	}

	client := staking.NewStakingClient(conn)
	return conn, client, nil
}

// ConnectTLS connects to server using TLS Certificate
func ConnectTLS(address string, tlsPath string) (*grpc.ClientConn, error) {

	// Open and read tls file containing connection information
	b, err := ioutil.ReadFile(tlsPath)
	if err != nil {
		return nil, err
	}

	// Add Credentials to a certificate pool
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(b) {
		return nil, fmt.Errorf("credentials: failed to append certificates")
	}

	// Create new TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		RootCAs:    certPool,
		ServerName: identity.CommonName,
	})

	// Add Credentials to grpc options to be used for TLS Connection
	opts := grpc.WithTransportCredentials(creds)
	conn, err := cmnGrpc.Dial(
		address,
		opts,
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Connect - connect to grpc
// Add grpc.WithBlock() and grpc.WithTimeout()
// to have dial to constantly try and establish connection
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
