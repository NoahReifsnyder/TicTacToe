// A very simple web application in Go
package main

import (
	"flag"
	"log"
	"net/http"
)

// The main function just configures resources and starts listening for new
// web requests
func main() {
	// parse command line options
	configPath := flag.String("configfile", "config.json", "Path to the configuration (JSON) file")
	flag.Parse()

	// load the JSON config file
	loadConfig(*configPath)

	// set up a few routes
	http.HandleFunc("/public/", handleFile)
	http.HandleFunc("/data/", handleData)

	// my new route
	http.HandleFunc("/spear/", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hi Spear"))
	})

	// print a diagnostic message and start the server
	log.Println("Server running on port " + cfg.AppPort)
	http.ListenAndServe(":"+cfg.AppPort, nil)
}