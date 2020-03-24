package config_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// Setting data to test with, valid and invalid path locations
const (
	portFile       = "../config/test_user_config_main.ini"
	socketFile     = "../config/test_config_nodes.ini"
	portFileFail   = "../config/test_config_main_fail.ini"
	socketFileFail = "../config/test_config_nodes_fail.ini"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	os.Chdir("../")
	// Set the Logger that will be used by the API through all the packages
	lgr.SetLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

// Testing the loading of the default port configuration file
func TestLoadPortConfig_Success_1(t *testing.T) {
	portConf := config.LoadPortConfiguration()
	if portConf == nil {
		t.Errorf("Failed to load the port file from path.")
	}
}

// Testing the loading of the default socket configuration file
func TestLoadSocketConfig_Success_1(t *testing.T) {
	socketConf := config.LoadSocketConfiguration()
	if socketConf == nil {
		t.Errorf("Failed to load the socket file from path.")
	}
}

// Testing the successful retrieval of ports once the configuration is loaded
func TestLoadPortConfig_Success_2(t *testing.T) {
	portConf := config.LoadSocketConfiguration()
	if portConf == nil {
		t.Errorf("Failed to load the port file from path.")
	}
	portConf_1 := config.GetPort()
	if portConf_1 == nil {
		t.Errorf("Failed to load the port file from path.")
	}
}

// Testing the successful retrieval of sockets once the configuration is loaded
func TestLoadSocketConfig_Success_2(t *testing.T) {
	socketConf := config.LoadSocketConfiguration()
	if socketConf == nil {
		t.Errorf("Failed to load the socket file from path.")
	}
	socketConf_1 := config.GetSockets()
	if socketConf_1 == nil {
		t.Errorf("Failed to load the socket file from path.")
	}
}

// Testing the failed loading of the another port configuration file
func TestLoadPortConfig_Failure_1(t *testing.T) {
	config.SetPortFile(portFileFail)
	portConf := config.LoadPortConfiguration()
	if portConf != nil {
		t.Errorf("Failed to not load the port file from path.")
	}
}

// Testing the failed loading of another socket configuration file
func TestLoadSocketConfig_Failure_1(t *testing.T) {
	config.SetSocketFile(socketFileFail)
	socketConf := config.LoadSocketConfiguration()
	if socketConf != nil {
		t.Errorf("Failed to not load the socket file from path.")
	}
}
