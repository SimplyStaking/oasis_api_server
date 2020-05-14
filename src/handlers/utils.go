package handlers

import (
	"strconv"
	"sync"

	"github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
)

// Function to verify and retrieve sentry data
func checkSentryData(nodeName string) (bool, string, string) {
	mutex := &sync.RWMutex{}
	mutex.Lock()
	// Get Sentry IP and localtion of TLS Cert
	allSentries := config.GetSentryData()
	for _, sentry := range allSentries {
		// If nodeName is in configuration reply with it's websocket
		if sentry["node_name"] == nodeName {
			lgr.Info.Printf("Requested sentry %s was found!",
				nodeName)
			return true, sentry["ext_url"], sentry["tls_path"]
		}
	}

	// If nodeName isn't in configuration produce Log and Reply with False
	lgr.Error.Printf("Requested sentry %s was not found, check if "+
		"configured!", nodeName)

	mutex.Unlock()
	return false, "", ""
}

// Function to check if node name is in configuration and return socket for it
func checkNodeName(nodeName string) (bool, string) {
	mutex := &sync.RWMutex{}
	mutex.Lock()
	// Check if nodeName is in configuration
	allSockets := config.GetNodes()
	for _, socket := range allSockets {

		// If nodeName is in configuration reply with it's websocket
		if socket["node_name"] == nodeName {
			lgr.Info.Printf("Requested node %s was found!",
				nodeName)
			return true, socket["isocket_path"]
		}
	}

	// If nodeName isn't in configuration produce Log and Reply with False
	lgr.Error.Printf(
		"Requested node %s was not found, check if configured!",
		nodeName)

	mutex.Unlock()
	return false, ""
}

// Function to check if node name has prometheus configuration for it
func checkNodeNamePrometheus(nodeName string) (bool, string) {
	mutex := &sync.RWMutex{}
	mutex.Lock()
	// Check if nodeName is in configuration
	nodeData := config.GetNodes()
	for _, node := range nodeData {

		// If nodeName is in configuration reply with it's prometheus
		// url
		if node["node_name"] == nodeName {
			lgr.Info.Printf("Requested node %s was found!",
				nodeName)
			return true, node["prometheus_url"]
		}
	}

	// If nodeName isn't in configuration produce Log and Reply with False
	lgr.Error.Printf(
		"Requested node %s was not found, check if configured!",
		nodeName)

	mutex.Unlock()
	return false, ""
}

// Function to check if height is valid or to set height to latest
func checkHeight(recvHeight string) int64 {
	mutex := &sync.RWMutex{}
	mutex.Lock()
	// Declare height here so that it can be set inside if statement
	var height int64

	// If string is empty meaning no optional
	//  parameter was passed use latest height
	if len(recvHeight) == 0 {
		height = consensus.HeightLatest
		lgr.Info.Println("No height specified getting latest height!")
	} else {

		// If height isn't empty attempt to parse it into int64
		_, err := (strconv.ParseInt(recvHeight, 10, 64))
		if err != nil {

			// If it fails it means that string given
			// wasn't number and return result for
			lgr.Error.Println("Unexpected value found, required "+
				"string of int but received ", recvHeight)
			return -1
		}

		// If succeeded then parse it again and set height.
		height, _ = (strconv.ParseInt(recvHeight, 10, 64))
	}
	mutex.Unlock()
	return height
}

// Function to check if Kind is valid
func checkKind(recvKind string) int64 {

	mutex := &sync.RWMutex{}
	mutex.Lock()

	// Declare kind here so that it can be set inside if statement
	var kind int64

	// If string is empty meaning no optional parameter
	// was passed use latest kind therefore set kind to 0
	if len(recvKind) == 0 {
		kind = 0
		lgr.Info.Println("No Kind is specified setting kind to 0!")
	} else {

		// If kind isn't empty attempt to parse it into int64
		_, err := (strconv.ParseInt(recvKind, 10, 64))
		if err != nil {

			// If it fails it means that string
			// given wasn't number and return result for
			lgr.Error.Println("Unexpected value found, required "+
				"string of int but received ", recvKind)
			return -1
		}

		// If succeeded then parse it again and set kind.
		kind, _ = (strconv.ParseInt(recvKind, 10, 64))
	}
	mutex.Unlock()
	return kind
}

// Function to check if amount is valid
func checkAmount(recvAmount string) int64 {
	mutex := &sync.RWMutex{}
	mutex.Lock()
	// Declare amount here so that it can be set inside if statement
	var amount int64
	var err error

	// If string is empty meaning no optional parameter
	// was passed use latest amount therefore set amount to 0
	if len(recvAmount) == 0 {
		amount = 0
		lgr.Info.Println("No amount is specified setting amount to 0!")
	} else {

		// If amount isn't empty attempt to parse it into int64
		amount, err = (strconv.ParseInt(recvAmount, 10, 64))
		if err != nil {
			// If it fails it means that string
			// given wasn't number and return result for
			lgr.Error.Println("Unexpected value found, required "+
				"string of int but received ", recvAmount)
			return -1
		}
	}
	mutex.Unlock()
	return amount
}

// Function to check if a Node Exporter URL exists
func getNodeExporter() (bool, string) {
	mutex := &sync.RWMutex{}
	mutex.Lock()
	// Check if nodeName is in the configuration
	mainInfo := config.GetMain()

	if mainInfo["api_server"]["metrics_url"] != "" {
		lgr.Info.Println("Requested node Node Exporter was found!")
		return true, mainInfo["api_server"]["metrics_url"]
	}

	// If the Node Exporter was not configured then reply with False
	lgr.Error.Println(
		"Requested node was not found, check if configured!")
	mutex.Unlock()
	return false, ""
}
