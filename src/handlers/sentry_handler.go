package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/SimplyVC/oasis_api_server/src/rpc"
	sentry "github.com/oasislabs/oasis-core/go/sentry/api"
)

// loadSentryClient loads sentry client and returns it
func loadSentryClient(socket string, tls string) (*grpc.ClientConn,
	sentry.Backend) {

	// Attempt to load connection with sentry client
	connection, sentryClient, err := rpc.SentryClient(socket, tls)
	if err != nil {
		lgr.Error.Println(
			"Failed to establish connection to sentry client : ", err)
		return nil, nil
	}
	return connection, sentryClient
}

// GetSentryAddresses returns list of consensus and committee addresses of
// sentry node.
func GetSentryAddresses(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of sentry from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, extURL, tlsPath := checkSentryData(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Sentry name requested doesn't exist"})
		return
	}

	// Attempt to load connection with sentry client
	connection, sy := loadSentryClient(extURL, tlsPath)

	// Close connection once code underneath executes
	defer connection.Close()

	// If null object was retrieved send response
	if sy == nil {

		// Stop code here faild to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to establish connection using url : " + extURL})
		return
	}

	// Retrieve addresses connected to sentry
	sentryAddresses, err := sy.GetAddresses(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to get Sentry AddressesS!"})
		lgr.Error.Println(
			"Request at /api/sentry/addresses/ Failed to get addresses : ", err)
		return
	}

	// Responding with addresses connected to a sentry
	// retrieved from sentry client
	lgr.Info.Println(
		"Request at /api/sentry/addresses/ responding with Sentry Addresses!")
	json.NewEncoder(w).Encode(responses.SentryResponse{
		SentryAddresses: sentryAddresses})
}
