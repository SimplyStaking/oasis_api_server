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

// GetMemory returns the memory statistics of the current system
func GetMemory(w http.ResponseWriter, r *http.Request) {

	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Returning the memory currently being used by the system
	mem, err := memory.Get()
	if err != nil {
		lgr.Error.Println("Error while attempting to get memory of the system ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Error while attempting to get memory of the system."})
		return
	}

	// Responding with the retrieved memory statistics
	lgr.Info.Println("Request at /api/system/memory/ responding with Memory Statistics")
	json.NewEncoder(w).Encode(responses.MemoryResponse{Memory: mem})
}

// GetDisk returns the memory statistics of the current system
func GetDisk(w http.ResponseWriter, r *http.Request) {

	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Returning the disk information currently being used by the system
	dsk, err := disk.Get()
	if err != nil {
		lgr.Error.Println("Error while attempting to get disk information of the system ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Error while attempting to get disk information of the system."})
		return
	}

	// Responding with the retrieved memory statistics
	lgr.Info.Println("Request at /api/system/disk/ responding with Disk Statistics")
	json.NewEncoder(w).Encode(responses.DiskResponse{Disk: dsk})
}

// GetCPU returns the CPU statistics of the current system
func GetCPU(w http.ResponseWriter, r *http.Request) {

	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Returning the CPU currently being used by the system
	cpuinfo, err := cpu.Get()
	if err != nil {
		lgr.Error.Println("Error while attempting to get CPU information of the system ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Error while attempting to get CPU information of the system."})
		return
	}

	// Responding with the retrieved CPU statistics
	lgr.Info.Println("Request at /api/system/cpu/ responding with CPU Statistics")
	json.NewEncoder(w).Encode(responses.CPUResponse{CPU: cpuinfo})
}

// GetNetwork returns the network statistics of the current system
func GetNetwork(w http.ResponseWriter, r *http.Request) {

	// Adding a header so that the receiver knows they are receiving a JSON structure
	w.Header().Add("Content-Type", "application/json")

	// Returning the network statistics currently being used by the system
	netwrk, err := network.Get()
	if err != nil {
		lgr.Error.Println("Error while attempting to get Network information of the system ", err)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Error: "Error while attempting to get Network information of the system."})
		return
	}

	// Responding with the network memory statistics
	lgr.Info.Println("Request at /api/system/network/ responding with Network Statistics")
	json.NewEncoder(w).Encode(responses.NetworkResponse{Network: netwrk})
}
