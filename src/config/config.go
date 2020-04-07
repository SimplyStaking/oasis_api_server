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
	confExporter   ini.Config
	confSentry     ini.Config
	mainConfigFile = "config/user_config_main.ini"
	nodesFile      = "config/user_config_nodes.ini"
	prometheusFile = "config/prometheus_config_main.ini"
	exporterFile   = "config/node_exporter_nodes.ini"
	sentryFile     = "config/user_config_sentry.ini"
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

// SetPrometheusFile containing prometheus configuration
func SetPrometheusFile(newFile string) {
	prometheusFile = newFile
}

// SetExporterFile containing the Node Exporter configuration
func SetExporterFile(newFile string) {
	exporterFile = newFile
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

// GetPrometheusFile returns Prometheus configuration
func GetPrometheusFile() map[string]map[string]string {
	return confPrometheus
}

// GetExporterFile returns Node_Extrasctor configuration
func GetExporterFile() map[string]map[string]string {
	return confExporter
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

// LoadExporterConfiguration loads Node Exporter configuration
func LoadExporterConfiguration() map[string]map[string]string {

	// Decode and read the file containing the Node Exporter information
	if err := ini.DecodeFile(exporterFile, &confExporter); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confExporter
}

// LoadSentryConfiguration loads sentry configuration details
func LoadSentryConfiguration() map[string]map[string]string {

	// Decode and read file containing sentry information
	if err := ini.DecodeFile(sentryFile, &confSentry); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confSentry
}
