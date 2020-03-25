package config

import (
	"github.com/claudetech/ini"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// Variables to be exported and used by application, set with default values
var (
	confMain       ini.Config
	confNodes      ini.Config
	confPrometheus ini.Config
	confExtractor  ini.Config
	mainConfigFile = "../config/user_config_main.ini"
	nodesFile      = "../config/user_config_nodes.ini"
	prometheusFile = "../config/prometheus_config_main.ini"
	extractorFile  = "../config/extractor_config_main.ini"
)

// SetMainFile sets file location containing API configuration
func SetMainFile(newFile string) {
	mainConfigFile = newFile
}

// SetNodesFile sets file location containing Node configuration
func SetNodesFile(newFile string) {
	nodesFile = newFile
}

// SetPrometheusFile containing prometheus configuration
func SetPrometheusFile(newFile string) {
	prometheusFile = newFile
}

// SetExtractorFile containing the Node_Extractor configuration
func SetExtractorFile(newFile string) {
	extractorFile = newFile
}

// GetMain returns Main API configuration
func GetMain() map[string]map[string]string {
	return confMain
}

// GetNodes returns Nodes configuration
func GetNodes() map[string]map[string]string {
	return confNodes
}

// GetPrometheusFile returns Prometheus configuration
func GetPrometheusFile() map[string]map[string]string {
	return confPrometheus
}

// GetExtractorFile returns Node_Extrasctor configuration
func GetExtractorFile() map[string]map[string]string {
	return confExtractor
}

// LoadMainConfiguration loads main configuration file from config folder
func LoadMainConfiguration() map[string]map[string]string {

	// Decode and read file containing Main API information
	if err := ini.DecodeFile(mainConfigFile, &confMain); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confMain
}

// LoadNodesConfiguration loads node configuration file from config folder
func LoadNodesConfiguration() map[string]map[string]string {

	// Decode and read file containing Node information
	if err := ini.DecodeFile(nodesFile, &confNodes); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confNodes
}

// LoadPrometheusConfiguration loads prometheus configuration
func LoadPrometheusConfiguration() map[string]map[string]string {

	// Decode and read file containing prometheus information
	if err := ini.DecodeFile(prometheusFile, &confPrometheus); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confPrometheus
}

// LoadExtractorConfiguration loads Node_Extractor configuration
func LoadExtractorConfiguration() map[string]map[string]string {

	// Decode and read the file containing the Node_Extractor information
	if err := ini.DecodeFile(extractorFile, &confExtractor); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confExtractor
}
