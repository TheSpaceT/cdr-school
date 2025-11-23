package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Simple counter metric
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
)

func init() {
	// Register the metric with Prometheus
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		// Increment the counter
		httpRequestsTotal.WithLabelValues("/foo").Inc()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// Expose Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
