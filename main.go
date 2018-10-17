package main

import (
  "log"
  "net/http"
)

// Define a Home function which writes a plain-text "Hello from Snippetbox"
// message as the HTTP response body
func Home(w http.ResponseWriter, r *http.Request) {
  // Use r.URL.Path to check whether the request URL path exactly matches "/".
  // If it doesn't, use the w.WriteHeader() method to send a 404 status code and
  // the w.Write() method to write a "Not Found" response body. Importantly, we
  // then return from the function so that the subsequent code is not executed.
  if r.URL.Path != "/" {
    // Use the http.NotFound() function to send a 404 Not Found response.
    http.NotFound(w, r)
    return
  }

  w.Write([]byte("Hello from Snippetbox!"))
}

func main() {
  // Use the http.NewServeMux() function to initialize a new serve mux. Then use
  // the mux.HandleFunc() method to register our Home function as the handler for
  // the "/" URL pattern.
  mux := http.NewServeMux()
  mux.HandleFunc("/", Home)

  // Use the http.ListenAndServe() function to start a new web server. We pass in
  // two parameters: the TCP network address to listen on (in this case ":4000")
  // and the serve mux we just created. If ListenAndServe() returns an error we
  // use the log.Fatal() function to record the error message and exit the program.
  log.Println("Starting server on :4000")

  // network address should be specified as "host:port", host can be ommitted for default to localhost
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
