package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir(".")))
	serveMux.Handle("/assets/logo.png", http.FileServer(http.Dir(".")))
	s := http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler))
}
