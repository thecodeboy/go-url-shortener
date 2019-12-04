package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/catinello/base62"
)

type payload struct {
	URL  string
	Name string
}

// DB is our mock database ;)
var DB map[int]string = make(map[int]string)
var index int = 0

func main() {
	http.HandleFunc("/api/v1/shorten", handleShorten)
	http.HandleFunc("/api/v1/redirect", handleRedirect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleShorten(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data payload
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	longURL := data.URL
	index = index + 1
	DB[index] = longURL
	fmt.Println(DB)
	// fmt.Fprintf(w, "longURL, %q", longURL)
	// Return JSON response
	data.URL = ""
	data.Name = strconv.Itoa(index)
	jsonResponse, _ := json.Marshal(data)
	w.Write(jsonResponse)
}

func handleRedirect(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data payload
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	hash := data.Name
	id, _ := base62.Decode(hash)
	// Return JSON response
	data.Name = ""
	data.URL = DB[id]
	jsonResponse, _ := json.Marshal(data)
	w.Write(jsonResponse)
}
