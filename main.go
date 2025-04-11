package main

import (
	"log"
	"net/http"
	"strconv"

	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	var apiCfg apiConfig
	serveMux := http.NewServeMux()
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))

	serveMux.Handle("/app", apiCfg.middlewareMetricsInc(handler))
	serveMux.Handle("/app/assets/", apiCfg.middlewareMetricsInc(handler))
	serveMux.Handle("/app/assets/logo.png", apiCfg.middlewareMetricsInc(handler))

	serveMux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	serveMux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits: " + strconv.Itoa(int(apiCfg.fileserverHits.Load()))))
	})

	serveMux.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits: " + strconv.Itoa(int(apiCfg.fileserverHits.Load()))))
	})

	s := http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	cfg.fileserverHits.Add(1)
	return next
}

func (cfg *apiConfig) middlewareMetricsReset(next http.Handler) http.Handler {
	cfg.fileserverHits.CompareAndSwap(cfg.fileserverHits.Load(), 0)
	return next
}

// func (cfg *apiConfig) getMetrics(w http.ResponseWriter, r *http.Request, h http.Handler) http.Handler {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Hits: " + strconv.Itoa(int(cfg.fileserverHits.Load()))))
// 	return h
// }
