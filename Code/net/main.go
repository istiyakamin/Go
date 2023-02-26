package main

import (
	"fmt"
	"net/http"
)

func main(){


	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/service", service)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `Welcome to homepage`);
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `Welcome to about`);
}

func service(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `Welcome to service`);
}

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `Welcome to contact`);
}