// When we get a request for a file ("/public/..."), this code will translate
// the request to the folder where the file is stored, and will then serve
// the file
package main

import (
	"log"
	"net/http"
	"os"
)

// ensure that the given path is a valid file, and not a directory
//
// NB: we have rather verbose logging here, so that we can see when
//     strange requests come in
func isValidFile(path string) bool {
	file, err := os.Open(cfg.FilePath + "/" + path)
	if err != nil {
		log.Println("open() error:", path, err) // file doesn't exist
		return false
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Println("stat() error:", path, err) // very rare error
		return false
	}

	if stat.IsDir() {
		log.Println("dir() error:", path) // request for directory
		return false
	}
	log.Println("Valid request for file", path)
	return true
}

// Handle the request for a file
func handleFile(w http.ResponseWriter, r *http.Request) {
	// Strip the leading "/public/" from the requested path
	requestedFile := r.URL.Path[len("/public/"):]
	// Serve the file if it's valid, otherwise send an error message
	if isValidFile(requestedFile) {
		http.ServeFile(w, r, cfg.FilePath+"/"+requestedFile)
	} else {
		http.NotFound(w, r)
	}
}