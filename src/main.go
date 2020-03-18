package main

import (
	"io/ioutil"
	"os"

	api_router "github.com/SimplyVC/oasis_api_server/src/api_router"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

//Main Function handles all the possible API routes.
func main() {
	//Set the Logger that will be used by the API through all the packages
	lgr.SetLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	//Start the server
	err := api_router.StartServer()
	if err != nil {
		lgr.Info.Println("Server Stopped")
	}
}
