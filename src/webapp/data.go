// This file simulates a database that holds names.  It's just enough for a
// demonstration.

// To test, use
//    ~/curl -X POST http://localhost:8080/data
//    ~/curl -X GET  http://localhost:8080/data

// WARNING: this is not production-ready code.  For example, it doesn't 
// even handle PUT and DELETE verbs

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// The database we are simulating has two field: an integer that is indexed,
// and a string that is not.  For simplicity, we represent the database as a
// map of integers to strings.  This means we can only find a string if we
// know its integer, and thus we don't have to store the integer and string
// in the database.
//
// The trick is that the database also needs a monotonically increasing 
// counter, so that we can assign new index values to new entries, and
// this is a WEB APP, so we need to be able to handle the possibility that
// two clients make requests simultaneously.  Thus we need a counter and 
// a lock as part of the simulated database
var database = struct {
	// This is the readers/writer mutex that protects the data structure
	sync.RWMutex
	// the counter, so that each entry has a unique ID
	counter int64
	// a map from integers to strings, so that we can get a name by its index
	elements map[int64]string
}{elements: make(map[int64]string)}

// Helper type: when a user GETs on /data/, we want to send the entire set 
// of rows from the database.  We will want to send the index of the row,
// as well as the row contents, but in the database we only store the row
// contents, with the index being implicit.  Making an explicit Element type
// for rows of the reply ensures that we can use json.Marshal() without pain. 
type Element struct {
	Index int64   // The index of the row
	Name  string  // The row contents.  For now, it's just a string
}

// When a request comes in for /data/, we currently aren't very nuanced.  We
// take one of two actions:
//
// - If the request is a GET for /data/, we clump all the contents of the
// - database into an array and send it back to the client
//
// - If the request is a POST, we add a dummy data entry to the database
//
// [TODO] To make this even remotely useful, you should first update POST so
//        that it gets data from r.Body.  Then you should add DELETE and PUT
//        verbs.  And you should allow /data/1 (etc) to request individual
//        database elements
func handleData(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	if r.Method == "GET" {
		// Handle GET /data/ by getting all.  Ignore single-item GET requests 
		// for now
		if r.URL.Path == "/data/" {
			var data []Element
			database.RLock()
			defer database.RUnlock()
			for k, v := range database.elements {
				data = append(data, Element{k, v})
			}
			payload, err := json.Marshal(data)
			if err != nil {
				log.Println("JSON error", err.Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Write([]byte(string(payload)))
		} else {
			w.Write([]byte("Individual GETs are not yet implemented"))
		}
	} else if r.Method == "POST" {
		// Handle POST /data/ by adding a new element to the database...
		database.Lock()
		defer database.Unlock()
		database.counter += 1
		x := database.counter
		database.elements[x] = strconv.FormatInt(x, 10)
		w.Write([]byte("dummy data added"))

		// NB: here's a helpful code snippet for reading JSON from the request:
		//
		// contents, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		//     ...
		// }
		// var d DataRow // NB: you'll need to define the DataRow type
		// err = json.Unmarshal(contents, &d)
	} else {
		w.Write([]byte("Method not yet supported"))
	}
}