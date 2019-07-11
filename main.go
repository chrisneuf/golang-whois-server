package main

import (
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

    //Finally, start the HTTP server on port 8080.
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}