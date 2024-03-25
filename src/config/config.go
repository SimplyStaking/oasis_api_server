package config

import (
	"fmt"

	"github.com/claudetech/ini"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// Variables to be exported and used by application, set with default values
var (
	confMain       ini.Config
	confNodes      ini.Config
	confSentry     ini.Config
	mainConfigFile = "user_config_main.ini"
	nodesFile      = "user_config_nodes.ini"
	sentryFile     = "user_config_sentry.ini"
)

// SetSentryFile sets file location containing sentry data
func SetSentryFile(newFile string) {
	sentryFile = newFile
}

// SetMainFile sets file location containing API configuration
func SetMainFile(newFile string) {
	mainConfigFile = newFile
}

// SetNodesFile sets file location containing Node configuration
func SetNodesFile(newFile string) {
	nodesFile = newFile
}

// GetSentryData returns Sentry configuration
func GetSentryData() map[string]map[string]string {
	return confSentry
}

// GetMain returns Main API configuration
func GetMain() map[string]map[string]string {
	return confMain
}

// GetNodes returns Nodes configuration
func GetNodes() map[string]map[string]string {
	return confNodes
}

// LoadMainConfiguration loads main configuration file from config folder
func LoadMainConfiguration(baseDir string) (map[string]map[string]string, error) {

	// Decode and read file containing Main API information
	mainConfigFilePath := fmt.Sprintf("%s/%s", baseDir, mainConfigFile)
	lgr.Info.Println("Loading main config from: %s", mainConfigFilePath)
	if err := ini.DecodeFile(mainConfigFilePath, &confMain); err != nil {
		lgr.Error.Println(err)
		return nil, err
	}
	return confMain, nil
}

// LoadNodesConfiguration loads node configuration file from config folder
func LoadNodesConfiguration(baseDir string) (map[string]map[string]string, error) {

	// Decode and read file containing Node information
	nodesConfigFilePath := fmt.Sprintf("%s/%s", baseDir, nodesFile)
	lgr.Info.Println("Loading nodes config from: %s", nodesConfigFilePath)
	if err := ini.DecodeFile(nodesConfigFilePath, &confNodes); err != nil {
		lgr.Error.Println(err)
		return nil, err
	}
	return confNodes, nil
}

// LoadSentryConfiguration loads sentry configuration details
func LoadSentryConfiguration(baseDir string) (map[string]map[string]string, error) {

	// Decode and read file containing sentry information
	sentryConfigFilePath := fmt.Sprintf("%s/%s", baseDir, sentryFile)
	lgr.Info.Println("Loading sentry config from: %s", sentryConfigFilePath)
	if err := ini.DecodeFile(sentryConfigFilePath, &confSentry); err != nil {
		lgr.Error.Println(err)
		return nil, err
	}
	return confSentry, nil
}
