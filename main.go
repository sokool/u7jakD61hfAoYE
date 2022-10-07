package main

import (
	"log"
	"net/http"
	"rollee/dictionary"
)

func main() {
	s := dictionary.NewStorage()
	h := dictionary.NewHTTP(s)
	http.Handle("/dictionary", h)

	log.Printf("Dictionary - implementation based on prefix tree data structure and with use of DFS algorithm to " +
		"read words from given prefixes")
	log.Printf("Please use followed HTTP endpoints: \n")
	log.Printf("[POST] http://localhost:8080/dictionary - to add word as json, in request body \n")
	log.Printf("[GET]  http://localhost:8080/dictionary?prefix={word} - to read word by it's prefix\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
