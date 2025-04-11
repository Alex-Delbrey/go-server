package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	dir := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	serveMux.Handle("/app", dir)
	serveMux.Handle("/app/assets/", dir)
	serveMux.Handle("/app/assets/logo.png", dir)

	serveMux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	s := http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler))
}
