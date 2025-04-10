package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	s := http.Server{
		Handler: serveMux,
		Addr:    ":8000",
	}
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler))
	fmt.Println("Hello World")
}
