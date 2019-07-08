package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("my-app/build"))
	for {
		http.Handle("/", fs)

		log.Println("Listening...")
		http.ListenAndServe(":4000", nil)
	}
}
