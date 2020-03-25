package handlers

import (
	"encoding/json"
	"net/http"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
)

// GetMemory returns memory statistics of current system
func GetMemory(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Returning memory currently being used by system
	mem, err := memory.Get()
	if err != nil {
		lgr.Error.Println(
			"Error while attempting to get memory of system ", err)

		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Error while attempting to get memory of system."})
		return
	}

	// Responding with retrieved memory statistics
	lgr.Info.Println("Request at /api/system/memory/ responding with Memory Statistics")
	json.NewEncoder(w).Encode(responses.MemoryResponse{Memory: mem})
}

// GetDisk returns memory statistics of current system
func GetDisk(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Returning disk information currently being used by system
	dsk, err := disk.Get()
	if err != nil {
		lgr.Error.Println(
			"Error while attempting to get disk information of system ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Error while attempting to get disk information of system."})
		return
	}

	// Responding with retrieved memory statistics
	lgr.Info.Println("Request at /api/system/disk/ responding with Disk Statistics")
	json.NewEncoder(w).Encode(responses.DiskResponse{Disk: dsk})
}

// GetCPU returns CPU statistics of current system
func GetCPU(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Returning CPU currently being used by system
	cpuinfo, err := cpu.Get()
	if err != nil {
		lgr.Error.Println(
			"Error while attempting to get CPU information of system ", err)

		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Error while attempting to get CPU information of system."})
		return
	}

	// Responding with retrieved CPU statistics
	lgr.Info.Println("Request at /api/system/cpu/ responding with CPU Statistics")
	json.NewEncoder(w).Encode(responses.CPUResponse{CPU: cpuinfo})
}

// GetNetwork returns network statistics of current system
func GetNetwork(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")

	// Returning network statistics currently being used by system
	netwrk, err := network.Get()
	if err != nil {
		lgr.Error.Println(
			"Error while attempting to get Network information of system ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			Error: "Error while attempting to get Network information of system."})
		return
	}

	// Responding with network memory statistics
	lgr.Info.Println("Request at /api/system/network/ responding with Network Statistics")
	json.NewEncoder(w).Encode(responses.NetworkResponse{Network: netwrk})
}
