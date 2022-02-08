package main

import (
	"encoding/json"
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
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	}
}

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var kv KV

func get(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	err := e.Encode(kv)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var req KV
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	kv = req

	e := json.NewEncoder(w)
	err = e.Encode(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}
}
