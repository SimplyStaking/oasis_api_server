package config

import (
	"github.com/claudetech/ini"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// Variables to be exported and used by application, set with default values
var (
	confPort       ini.Config
	confSockets    ini.Config
	confPrometheus ini.Config
	portFile       = "../config/user_config_main.ini"
	socketFile     = "../config/user_config_nodes.ini"
	prometheusFile = "../config/prometheus_config_main.ini"
)

// SetPortFile sets file location containing Port
func SetPortFile(newFile string) {
	portFile = newFile
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
	return confPort
}

// GetSockets returns Socket configuration
func GetSockets() map[string]map[string]string {
	return confSockets
}

// GetPrometheusFile File Configuration Details
func GetPrometheusFile() map[string]map[string]string {
	return confPrometheus
}

// LoadPortConfiguration loads port configuration file from config folder
func LoadPortConfiguration() map[string]map[string]string {

	// Decode and read file containing port information
	if err := ini.DecodeFile(portFile, &confPort); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confPort
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
