package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type myData struct {
	Owner string
	Name  string
}

func main() {
	// http.HandleFunc("/api/v1/shorten", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	http.HandleFunc("/api/v1/shorten", handleAddNewRepo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleAddNewRepo(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	owner := data.Owner
	// name := data.Name
	fmt.Fprintf(w, "Hello, %q", owner)
}
