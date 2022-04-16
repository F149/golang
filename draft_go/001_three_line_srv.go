package main

import (
	"log"
	"net/http"
)

// ------------ V.1 ------------
//
//
//func main() {
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Three line server answer")
//	})
//
//	log.Println("Start http server on port 8081")
//	log.Fatal(http.ListenAndServe(":8081", nil))
//}

// ------------ V.2 ------------

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Simple three line server\n"))
}

func main() {
	http.HandleFunc("/", Handler)

	log.Println("Start http server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
