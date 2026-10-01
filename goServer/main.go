package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}


	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello !!!")
}



func formHandler(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Parsing Form error", err)
		return
	}

	fmt.Fprint(w, "Post request success\n")
	  
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Email: %s\n", email)
	fmt.Fprintf(w, "Password: %s\n", password)

	if r.URL.Path != "/form" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}


	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	
}

func main() {
	fileSever := http.FileServer(http.Dir("./static"))


	http.Handle("/", fileSever)
	
	http.HandleFunc("/form", formHandler)
	
	http.HandleFunc("/hello", helloHandler)


	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}