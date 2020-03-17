package config

import (
	
	"github.com/claudetech/ini"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

//Variables to be exported and used by the application, set with default values
var (
	confPort ini.Config
	confSockets ini.Config
	confPrometheus ini.Config
	portFile = "../config/user_config_main.ini"
	socketFile = "../config/user_config_nodes.ini"
	prometheusFile = "../config/prometheus_config_main.ini"
)

//Set the file containing the Port
func SetPortFile(newFile string){
	portFile = newFile
}

//Set the file containing the Sockets
func SetSocketFile(newFile string){
	socketFile = newFile
}

//Set the file containing prometheus configuration
func SetPrometheusFile(newFile string){
	prometheusFile = newFile
}

//Return the Port configuration
func GetPort() (map [string]map[string]string) {
	return confPort
}

//Return the Socket configuration
func GetSockets() (map [string]map[string]string){
	return confSockets
}

//Get the Prometheus File Configuration Details
func GetPrometheusFile() (map [string]map[string]string){
	return confPrometheus
}

//Load the port configuration file from the config folder
func LoadPortConfiguration() (map [string]map[string]string){
	
	//Decode and read the file containing the port information
	if err := ini.DecodeFile(portFile, &confPort); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confPort
}

//Load the socket configuration file from the config folder
func LoadSocketConfiguration() (map [string]map[string]string){
	//Decode and read the file containing the port information
	if err := ini.DecodeFile(socketFile, &confSockets); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confSockets
}

//Load the Prometheus Configuration to be able to query it
func LoadPrometheusConfiguration() (map [string]map[string]string){
	//Decode and read the file containing the port information
	if err := ini.DecodeFile(prometheusFile, &confPrometheus); err != nil {
		lgr.Error.Println(err)
		return nil
	}
	return confPrometheus
}