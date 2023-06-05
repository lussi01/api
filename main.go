package main

import "net/http"
import "fmt"
import "encoding/json"

type Person struct {
	name string
	age string
	gender string
	job string
}

func getAllPersons(w http.ResponseWriter, r *http.Request) {
	persons := []Person{
		{ "mario", "21", "male", "developer"},
		{ "luigi", "23", "male", "developer"},
	}
	json.NewEncoder(w).Encode(persons)
}
		
func main() {
	http.HandleFunc("/persons", getAllPersons)
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe("localhost:8000", nil)
}