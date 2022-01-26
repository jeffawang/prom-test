package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "some_metric",
		Help: "Just a test metric",
	})
	views = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "views",
		Help: "Number of views",
	})
)

func main() {
	http.HandleFunc("/", root)

	prometheus.MustRegister(views)
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	views.Inc()
	w.Write([]byte("why hello there"))
}
