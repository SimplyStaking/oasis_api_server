package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/prometheus/common/expfmt"
)

var parser expfmt.TextParser

// PrometheusQueryGauge to retreive prometheus data.
func PrometheusQueryGauge(w http.ResponseWriter, r *http.Request) {

	lgr.Info.Println("Received request for /api/prometheus/gauge")

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, prometheusConfig := checkNodeNamePrometheus(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Setting gauge query
	gaugeName := r.URL.Query().Get("gauge")
	if gaugeName == "" {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve gauge name, please specify!"})
		lgr.Error.Println("Failed to retrieve gauge name, not specified!")
		return
	}

	resp, err := http.Get(prometheusConfig)
	if err != nil {
		lgr.Error.Println("Failed to retrieve Prometheus Data")
	}

	defer resp.Body.Close()

	// Read body response of Prometheus Configuration
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		lgr.Error.Println("Failed to read Prometheus Response")
	}

	parsed, err2 := parser.TextToMetricFamilies(bytes.NewReader(body))
	if err2 != nil {
		lgr.Error.Println("Failed to Parse Prometheus Response")
	}

	output := parsed[gaugeName].GetMetric()[0].GetGauge().GetValue()
	s := fmt.Sprintf("%f", output)

	json.NewEncoder(w).Encode(responses.SuccessResponse{Result: s})
	lgr.Info.Println(
		"Received request for /api/prometheus/gauge responding with : ", s)
}

// PrometheusQueryCounter to retreive prometheus data.
func PrometheusQueryCounter(w http.ResponseWriter, r *http.Request) {

	lgr.Info.Println("Received request for /api/prometheus/counter")

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Retrieving name of node from query request
	nodeName := r.URL.Query().Get("name")
	confirmation, prometheusConfig := checkNodeNamePrometheus(nodeName)
	if confirmation == false {

		// Stop code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node name requested doesn't exist"})
		return
	}

	// Setting counter query
	counterName := r.URL.Query().Get("counter")
	if counterName == "" {
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Failed to retrieve counter name, please specify!"})
		lgr.Error.Println("Failed to retrieve counter name, not specified!")
		return
	}

	resp, err := http.Get(prometheusConfig)
	if err != nil {
		lgr.Error.Println("Failed to retrieve Prometheus Data")
	}

	defer resp.Body.Close()

	// Read body response of Prometheus Configuration
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		lgr.Error.Println("Failed to read Prometheus Response")
	}

	parsed, err2 := parser.TextToMetricFamilies(bytes.NewReader(body))
	if err2 != nil {
		lgr.Error.Println("Failed to Parse Prometheus Response")
	}

	output := parsed[counterName].GetMetric()[0].GetCounter().GetValue()
	s := fmt.Sprintf("%f", output)

	json.NewEncoder(w).Encode(responses.SuccessResponse{Result: s})
	lgr.Info.Println(
		"Received request for /api/prometheus/counter responding with : ", s)
}
