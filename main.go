package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
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
} */
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send email to <a href=\"mailto:support@abc.com\">support@abc.com</a>.")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	//r.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", r)

}
