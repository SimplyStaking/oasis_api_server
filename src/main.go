package main

import (
	"os"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/router"
)

// Main Function handles all possible API routes.
func main() {

	// Set Logger that will be used by API through all packages
	lgr.SetLogger(os.Stdout, os.Stdout, os.Stderr)

	// Retrieve arguments for a base directory
	args := os.Args
	baseDir := "../config"
	if len(args) > 1 {
		baseDir = args[1]
	}
	lgr.Info.Println("Using base directory: ", baseDir)

	// Start server
	err := router.StartServer(baseDir)
	if err != nil {
		lgr.Info.Println("Server Stopped")
	}
}
