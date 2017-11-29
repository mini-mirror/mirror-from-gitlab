package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	http.ListenAndServe(":3000", router)
	//use handlerFunc for the request, it uses the default serve mux: &http.ServeMux
	//http.HandleFunc("/", handlerFunc)

	//start the server
	//localhost is omitted, nil means use the built-in serve mux
	//that comes with http pkg
	//http.ListenAndServe(":3000", nil)
}

//handlerFunc that handles all the incoming request
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "please contact <a href=\"mailto:eleanore.jin@gmail.com\">eleanore.jin@gmail.com</a>")
	} else {
		w.WriteHeader(404)
		fmt.Fprint(w, "<h1>Page Not Found</h1><p>please contact us</p>")
	}
}

//Index ...
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//Hello ...
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
