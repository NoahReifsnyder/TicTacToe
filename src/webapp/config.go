// This file exposes the Config object, which is populated with application
// configuration information.  The config information is loaded from a JSON
// file.
package main

// In Go, this is how we declare dependencies on other libraries
import (
	"encoding/json"
	"log"
	"os"
)

// Configuration information for our app.  For now, the only configuration is
// the port on which to run the server
//
// NB: field names must start with Capital letter for JSON parse to work
type Config struct {
	AppPort  string `json:"AppPort"`
	FilePath string `json:"FilePath"`
}

// A global variable storing the configuration information for the app
var cfg Config

// Load a JSON file that has all the config information for our app, and put
// the JSON contents into the cfg variable
func loadConfig(cfgFileName string) {
	// Load the file, fail on error
	f, err := os.Open(cfgFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse the file contents into the cfg object
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
}