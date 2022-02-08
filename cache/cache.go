package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	kvSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "kv_size",
		Help: "How many kv pairs are stored",
	})
	views = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "views",
		Help: "Number of views",
	})
)

func main() {
	http.HandleFunc("/", root)

	prometheus.MustRegister(views)
	prometheus.MustRegister(kvSize)
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

type KVStore struct {
	Data map[string]string `json:"data"`
}

type KVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var kv KVStore

func init() {
	kv.Data = make(map[string]string)
}

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

	var req KVPair
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	kv.Data[req.Key] = req.Value
	kvSize.Set(float64(len(kv.Data)))

	e := json.NewEncoder(w)
	err = e.Encode(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}
}
