package handlers

import (
	"strconv"

	config "github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
)

// Function to check if node name is in the configuration and return the socket for it
func checkNodeName(nodeName string) (bool, string) {
	// Check if nodeName is in the configuration
	allSockets := config.GetSockets()
	for _, socket := range allSockets {
		// If the nodeName is in the configuration reply with it's websocket
		if socket["node_name"] == nodeName {
			lgr.Info.Println("Requested node ", nodeName, "was found!")
			return true, socket["ws_url"]
		}
	}
	// If the nodeName isn't in the configuration produce Log and Reply with False
	lgr.Error.Println("Requested node ", nodeName, "was not found, check if configured!")
	return false, ""
}

// Function to check if node name has a prometheus configuration for it
func checkNodeNamePrometheus(nodeName string) (bool, string) {
	// Check if nodeName is in the configuration
	prometheus := config.GetPrometheusFile()
	for _, conf := range prometheus {
		// If the nodeName is in the configuration reply with it's websocket
		if conf["node_name"] == nodeName {
			lgr.Info.Println("Requested node ", nodeName, "was found!")
			return true, conf["ws_url"]
		}
	}
	// If the nodeName isn't in the configuration produce Log and Reply with False
	lgr.Error.Println("Requested node ", nodeName, "was not found, check if configured!")
	return false, ""
}

// Function to check if height is valid or to set height to latest
func checkHeight(recvHeight string) int64 {
	// Declare height here so that it can be set inside the if statement
	var height int64
	// If the string is empty meaning no optional parameter was passed use the latest height
	if len(recvHeight) == 0 {
		height = consensus.HeightLatest
		lgr.Info.Println("No height specified getting latest height!")
	} else {
		// If the height isn't empty attempt to parse it into a int64
		_, err := (strconv.ParseInt(recvHeight, 10, 64))
		if err != nil {
			// If it fails it means that the string given wasn't a number and return result for
			lgr.Error.Println("Unexpected value found, required string of int but received ", recvHeight)
			return -1
		}
		// If succeeded then parse it again and set the height.
		height, _ = (strconv.ParseInt(recvHeight, 10, 64))
	}
	return height
}

// Function to check if Kind is valid
func checkKind(recvKind string) int64 {
	// Declare kind here so that it can be set inside the if statement
	var kind int64
	// If the string is empty meaning no optional parameter was passed use the latest kind therefore set the kind to 0
	if len(recvKind) == 0 {
		kind = 0
		lgr.Info.Println("No Kind is specified setting kind to 0!")
	} else {
		// If the kind isn't empty attempt to parse it into a int64
		_, err := (strconv.ParseInt(recvKind, 10, 64))
		if err != nil {
			// If it fails it means that the string given wasn't a number and return result for
			lgr.Error.Println("Unexpected value found, required string of int but received ", recvKind)
			return -1
		}
		// If succeeded then parse it again and set the kind.
		kind, _ = (strconv.ParseInt(recvKind, 10, 64))
	}
	return kind
}

// Function to check if amount is valid
func checkAmount(recvAmount string) int64 {
	// Declare amount here so that it can be set inside the if statement
	var amount int64
	var err error
	// If the string is empty meaning no optional parameter was passed use the latest amount therefore set the amount to 0
	if len(recvAmount) == 0 {
		amount = 0
		lgr.Info.Println("No amount is specified setting amount to 0!")
	} else {
		// If the amount isn't empty attempt to parse it into a int64
		amount, err = (strconv.ParseInt(recvAmount, 10, 64))
		if err != nil {
			// If it fails it means that the string given wasn't a number and return result for
			lgr.Error.Println("Unexpected value found, required string of int but received ", recvAmount)
			return -1
		}
	}
	return amount
}
