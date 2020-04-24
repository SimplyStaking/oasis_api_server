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

// NodeExporterQueryGauge to retreive exporter data.
func NodeExporterQueryGauge(w http.ResponseWriter, r *http.Request) {

	lgr.Info.Println("Received request for /api/exporter/gauge")

	// Adding header so that receiver knows they are receiving JSON structure
	w.Header().Add("Content-Type", "application/json")

	//Get Node Exporter Metrics URl
	confirmation, exporterConfig := getNodeExporter()
	if confirmation == false {

		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node Exporter is not configured!"})
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

	resp, err := http.Get(exporterConfig)
	if err != nil {
		lgr.Error.Println(
			"Failed to retrieve Prometheus Data from Node Exporter Response")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to"+
			" retrieve Prometheus Data check if Node Exporter is Enabled!"})
		return
	}

	defer resp.Body.Close()

	// Read the body response from the Node Exporter
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		lgr.Error.Println(
			"Failed to read the Node Exporter Response")
	}

	parsed, err2 := parser.TextToMetricFamilies(bytes.NewReader(body))
	if err2 != nil {
		lgr.Error.Println("Failed to Parse the Node Exporter Response")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to"+
			" read Node Exporter Response."})
		return
	}

	if len(parsed[gaugeName].GetMetric()) <= 0 {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Metric name"+
		" doesn't exist!"})
		lgr.Info.Println(
			"Received request for /api/exporter/gauge but Metric name "+
			"doesn't exit!")
		return
	}

	output := parsed[gaugeName].GetMetric()[0].GetGauge().GetValue()
	s := fmt.Sprintf("%f", output)

	json.NewEncoder(w).Encode(responses.SuccessResponse{Result: s})
	lgr.Info.Println(
		"Received request for /api/exporter/gauge responding with : ", s)
}

// NodeExporterQueryCounter to retreive exporter data.
func NodeExporterQueryCounter(w http.ResponseWriter, r *http.Request) {

	lgr.Info.Println("Received request for /api/exporter/counter")

	// Adding header so that receiver knows they are receiving JSON structure
	w.Header().Add("Content-Type", "application/json")

	//Get Node Exporter Metrics URl
	confirmation, exporterConfig := getNodeExporter()
	if confirmation == false {

		// Stop the code here no need to establish connection and reply
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Node Exporter is not configured!"})
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

	resp, err := http.Get(exporterConfig)
	if err != nil {
		lgr.Error.Println("Failed to retrieve Node Exporter Data")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to"+
			" retrieve Prometheus Data check if Node Exporter is Enabled!"})
		return
	}

	defer resp.Body.Close()

	// Read the body response of the Node Exporter
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		lgr.Error.Println("Failed to read the Node Exporter Response")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to"+
			" read Node Exporter Response."})
		return
	}

	parsed, err2 := parser.TextToMetricFamilies(bytes.NewReader(body))
	if err2 != nil {
		lgr.Error.Println("Failed to Parse the Node Exporter Response")
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Failed to"+
			" Parse Node Exporter Response."})
		return
	}

	if len(parsed[counterName].GetMetric()) <= 0 {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Metric name"+
		" doesn't exist!"})
		lgr.Info.Println(
			"Received request for /api/exporter/counter but Metric name "+
			"doesn't exit!")
		return
	}

	output := parsed[counterName].GetMetric()[0].GetCounter().GetValue()
	s := fmt.Sprintf("%f", output)

	json.NewEncoder(w).Encode(responses.SuccessResponse{Result: s})
	lgr.Info.Println(
		"Received request for /api/exporter/counter responding with : ", s)
}
