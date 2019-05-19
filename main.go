package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Home Page!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send email to <a href=\"mailto:support@abc.com\">support@abc.com</a>.")
		//fmt.Fprint(w, "To get in touch, please contact <a href=\"mailto:support@lenslocked.com\">support@abc.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> We could not find the page you are looking for. Sorry :( </h1>")
	}
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// reason to not use it always because of 3rd arg defined in the function
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	//router.GET("/hello/:name", Hello)
	router.GET("/hello/:name/spanish", Hello)
	//http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", router)

}
