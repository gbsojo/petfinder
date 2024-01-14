package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Animal struct {
	Name    string `json:"name"`
	Species string `json:"species"`
}

func handleDogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		fmt.Printf("GET\n")
		data := []Animal{{"Pepito", "dog"}, {"Rocco", "dog"}, {"Bella", "dog"}, {"Luna", "dog"}}

		json.NewEncoder(w).Encode(data)
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}
}

func handleCats(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Printf("GET\n")
		data := []Animal{{"Fito", "cat"}, {"Kali", "cat"}, {"Birba", "cat"}, {"Chai", "cat"}}

		json.NewEncoder(w).Encode(data)
	default:
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}
}

func main() {
	http.HandleFunc("/api/pets/dogs", handleDogs)
	http.HandleFunc("/api/pets/cats", handleCats)

	const portNumber string = ":8080"

	fmt.Printf("Starting Petfinder server in port %s...\n", portNumber)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal(err)
	}
}
