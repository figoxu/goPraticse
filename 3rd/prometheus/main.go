package main

import (
	"github.com/alexgtn/go-middleware-metrics/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	r := mux.NewRouter()


	metricsMiddleware := middleware.NewMetricsMiddleware()

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/lemon", lemonHandler).Methods(http.MethodGet)
	r.HandleFunc("/potato", potatoHandler).Methods(http.MethodPost)

	r.Use(metricsMiddleware.Metrics)

	http.ListenAndServe(":8080", r)
}

func lemonHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Lemon"))
}

func potatoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Potato"))
}

