package main

import (
	"github.com/domainr/whois"

	"encoding/json"
    "log"
    "net/http"
)

func main() {
    // Create a fileServer handler that serves our static files.
    fileServer := http.FileServer(http.Dir("static/"))

    //Pass Requests to the file server
    http.HandleFunc(
        "/",
        func(w http.ResponseWriter, r *http.Request) {
            fileServer.ServeHTTP(w, r)
        },
	)
	
	// Handle the POST request on /whois
	// client dat format: data=8.8.8.8
	http.HandleFunc(
		"/whois", 
		func(w http.ResponseWriter, r *http.Request) {
		// Verify this is POST request
			if r.Method != "POST" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Extract the encoded data to perform the whois on.
			data := r.PostFormValue("data")

			// Perform the whois query.
			result, err := whoisQuery(data)

			// Return a JSON-encoded response.
			if err != nil {
				jsonResponse(w, Response{Error: err.Error()})
				return
			}
			jsonResponse(w, Response{Result: result})
		},
	)

    //Finally, start the HTTP server on port 8080.
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type Response struct {
    Error string `json:"error"`
    Result string `json:"result"`
}

//Runs a who is query using github.com/domainr/whois
func whoisQuery(data string) (string, error) {
	response, err := whois.Fetch(data)
	if err != nil {
		return "", err
	}
	return string(response.Body), nil
}
func jsonResponse(w http.ResponseWriter, x interface{}) {
    // JSON-encode x.
    bytes, err := json.Marshal(x)
    if err != nil {
        panic(err)
    }
    // Write the encoded data to the ResponseWriter.
    // This will send the response to the client.
    w.Header().Set("Content-Type", "application/json")
    w.Write(bytes)
}
