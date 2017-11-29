package main

import (
	"fmt"
	"net/http"
)

func main() {
	//use handlerFunc for the request
	http.HandleFunc("/", handlerFunc)
	//start the server
	//localhost is omitted, nil means use the built-in serve mux
	//that comes with http pkg
	http.ListenAndServe(":3000", nil)
}

//handlerFunc that handles all the incoming request
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}
