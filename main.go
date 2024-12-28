package main

import (
	"log"
	"net/http"

	"tdriven.com/dependencyinjection"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
