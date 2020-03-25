package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
)

// NodeExtractorQueryGauge to retreive extractor data.
func NodeExtractorQueryGauge(w http.ResponseWriter, r *http.Request) {

	lgr.Info.Println("Received request for /api/extractor/gauge")

	// Adding header so that receiver knows they are receiving JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, extractorConfig := checkNodeNameExtractor(nodeName)
	if confirmation == false {

		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Setting the gauge query
	gaugeName := r.URL.Query().Get("gauge")
	if gaugeName == "" {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve gauge name, please specify!"})
		lgr.Error.Println(
			"Failed to retrieve gauge name, not specified!")
		return
	}

	resp, err := http.Get(extractorConfig)
	if err != nil {
		lgr.Error.Println(
			"Failed to retrieve Prometheus Data from Node Extractor Response")
	}

	defer resp.Body.Close()

	// Read the body response from the Node Extractor
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		lgr.Error.Println(
			"Failed to read the Node Extractor Response")
	}

	parsed, err2 := parser.TextToMetricFamilies(bytes.NewReader(body))
	if err2 != nil {
		lgr.Error.Println("Failed to Parse the Node Extractor Response")
	}

	output := parsed[gaugeName].GetMetric()[0].GetGauge().GetValue()
	s := fmt.Sprintf("%f", output)

	json.NewEncoder(w).Encode(responses.SuccessResponse{Result: s})
	lgr.Info.Println(
		"Received request for /api/extractor/gauge responding with : ", s)
}

// NodeExtractorQueryCounter to retreive extractor data.
func NodeExtractorQueryCounter(w http.ResponseWriter, r *http.Request) {

	lgr.Info.Println("Received request for /api/extractor/counter")

	// Adding header so that receiver knows they are receiving JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Retrieving the name of the node from the query request
	nodeName := r.URL.Query().Get("name")
	confirmation, extractorConfig := checkNodeNameExtractor(nodeName)
	if confirmation == false {

		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Setting the counter query
	counterName := r.URL.Query().Get("counter")
	if counterName == "" {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve counter name, please specify!"})
		lgr.Error.Println(
			"Failed to retrieve counter name, not specified!")
		return
	}

	resp, err := http.Get(extractorConfig)
	if err != nil {
		lgr.Error.Println("Failed to retrieve Node Extractor Data")
	}

	defer resp.Body.Close()

	// Read the body response of the Node Extractor
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		lgr.Error.Println("Failed to read the Node Extractor Response")
	}

	parsed, err2 := parser.TextToMetricFamilies(bytes.NewReader(body))
	if err2 != nil {
		lgr.Error.Println("Failed to Parse the Node Extractor Response")
	}

	output := parsed[counterName].GetMetric()[0].GetCounter().GetValue()
	s := fmt.Sprintf("%f", output)

	json.NewEncoder(w).Encode(responses.SuccessResponse{Result: s})
	lgr.Info.Println(
		"Received request for /api/extractor/counter responding with : ", s)
}
