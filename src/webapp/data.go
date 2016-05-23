// This file simulates a database that holds names.  It's just enough for a
// demonstration.

// To test, use
//    ~/curl -X POST http://localhost:8080/data
//    ~/curl -X GET  http://localhost:8080/data
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// The database itself will be indexed by an integer, and will match that
// integer to a name.  However, when we want to send the entire database
// contents in response to a query, we need to send the index and name.  This
// struct makes it easy to produce JSON with the response.
type Element struct {
	Index int64
	Name  string
}

// The database we are simulating holds integers and strings.  However, we
// need a bit more nuance to be correct.  First, there is a counter, which
// ensures we don't have duplicate keys.  Second, there is a Readers/Writer
// lock protecting the database, so that we don't have concurrency bugs.
var database = struct {
	// This is the readers/writer mutex that protects the data structure
	sync.RWMutex
	// the counter, so that each entry has a unique ID
	counter int64
	// a map from integers to strings, so that we can get a name by its index
	elements map[int64]string
}{elements: make(map[int64]string)}

// When a request comes in for /data/, we currently aren't very nuanced.  We
// take one of two actions:
//
// - If the request is a GET for /data/, we clump all the contents of the
// - database into an array and send it back to the client
//
// - If the request is a POST, we add a dummy data entry to the database
//
// [TODO] To make this even remotely useful, you should first update POST so
//        that it gets data from r.Body.  Then we should add DELETE and PUT
//        verbs.  And we should allow /data/1 (etc) to request individual
//        database elements
func handleData(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	if r.Method == "GET" {
		// Handle GET /data/ by getting all.  Ignore other GET requests for
		// now
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

		// NB: here's a helpful code snippet for reading JSON from the request:
		//
		// contents, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		//     ...
		// }
		// var d DataRow
		// err = json.Unmarshal(contents, &d)
		database.Lock()
		defer database.Unlock()
		database.counter += 1
		x := database.counter
		database.elements[x] = strconv.FormatInt(x, 10)
		w.Write([]byte("dummy data added"))
	}
}