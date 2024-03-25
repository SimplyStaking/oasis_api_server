package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// Setting data to test with, valid and invalid path locations
const (
	mainConfigFile = "config/user_config_main.ini"
	nodesFile      = "config/user_config_nodes.ini"
	mainFileFail   = "test_config_main_fail.ini"
	nodesFileFail  = "test_config_nodes_fail.ini"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	os.Chdir("../")

	// Set Logger that will be used by API through all packages
	lgr.SetLogger(os.Stdout, os.Stdout, os.Stderr)
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestLoadMainConfiguration_Success_1(t *testing.T) {
	mainConf, _ := config.LoadMainConfiguration("../config")
	if mainConf == nil {
		t.Errorf("Failed to load main config file from path.")
	}
}

func TestLoadNodesConfig_Success_1(t *testing.T) {
	nodesConf, _ := config.LoadNodesConfiguration("../config")
	if nodesConf == nil {
		t.Errorf("Failed to load node config file from path.")
	}
}
func TestLoadMainConfiguration_Success_2(t *testing.T) {
	mainConf, _ := config.LoadNodesConfiguration("../config")
	if mainConf == nil {
		t.Errorf("Failed to load main config file from path.")
	}
	mainConf_1 := config.GetMain()
	if mainConf_1 == nil {
		t.Errorf("Failed to load main config file from path.")
	}
}

func TestLoadNodesConfig_Success_2(t *testing.T) {
	nodesConf, _ := config.LoadNodesConfiguration("../config")
	if nodesConf == nil {
		t.Errorf("Failed to load node config file from path.")
	}
	nodesConf_1 := config.GetNodes()
	if nodesConf_1 == nil {
		t.Errorf("Failed to load node config file from path.")
	}
}

func TestLoadMainConfiguration_Failure_1(t *testing.T) {
	config.SetMainFile(mainFileFail)
	mainConf, _ := config.LoadMainConfiguration("../config")
	if mainConf != nil {
		t.Errorf("Failed to not load main config file from path.")
	}
}

func TestLoadNodesConfig_Failure_1(t *testing.T) {
	config.SetNodesFile(nodesFileFail)
	nodesConf, _ := config.LoadNodesConfiguration("../config")
	if nodesConf != nil {
		t.Errorf("Failed to not load nodes file from path.")
	}
}
