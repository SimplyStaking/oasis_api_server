package main

import (
	"io/ioutil"
	"os"

	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/router"
)

// Main Function handles all possible API routes.
func main() {

	// Set Logger that will be used by API through all packages
	lgr.SetLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	// Start server
	err := router.StartServer()
	if err != nil {
		lgr.Info.Println("Server Stopped")
	}
}
