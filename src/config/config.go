package config

import (
	"github.com/claudetech/ini"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// Variables to be exported and used by application, set with default values
var (
	confMain       ini.Config
	confSockets    ini.Config
	confPrometheus ini.Config
	mainConfigFile = "../config/user_config_main.ini"
	socketFile     = "../config/user_config_nodes.ini"
	prometheusFile = "../config/prometheus_config_main.ini"
)

// SetPortFile sets file location containing Port
func SetPortFile(newFile string) {
	mainConfigFile = newFile
}

// SetSocketFile sets file location containing Sockets
func SetSocketFile(newFile string) {
	socketFile = newFile
}

// SetPrometheusFile containing prometheus configuration
func SetPrometheusFile(newFile string) {
	prometheusFile = newFile
}

// GetPort returns Port configuration
func GetPort() map[string]map[string]string {
	return confMain
}

// GetSockets returns Socket configuration
func GetSockets() map[string]map[string]string {
	return confSockets
}

// GetPrometheusFile File Configuration Details
func GetPrometheusFile() map[string]map[string]string {
	return confPrometheus
}

// LoadMainConfiguration loads port configuration file from config folder
func LoadMainConfiguration() map[string]map[string]string {

	// Decode and read file containing port information
	if err := ini.DecodeFile(mainConfigFile, &confMain); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confMain
}

// LoadSocketConfiguration loads socket configuration file from config folder
func LoadSocketConfiguration() map[string]map[string]string {

	// Decode and read file containing port information
	if err := ini.DecodeFile(socketFile, &confSockets); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confSockets
}

// LoadPrometheusConfiguration loads prometheus configuration so that it can be queried
func LoadPrometheusConfiguration() map[string]map[string]string {

	// Decode and read file containing port information
	if err := ini.DecodeFile(prometheusFile, &confPrometheus); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confPrometheus
}
