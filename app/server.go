package main

import (
  "math/rand"
  "net/http"
  "time"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

var reqLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
  Name: "http_request_duration_seconds",
  Help: "Request latency",
  Buckets: prometheus.DefBuckets,
})

var reqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
  Name: "http_requests_total",
  Help: "Requests by status",
}, []string{"code"})

func handler(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  if rand.Intn(10) < 2 { // 20% error
    w.WriteHeader(500); reqTotal.WithLabelValues("500").Inc()
  } else {
    time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
    w.WriteHeader(200); reqTotal.WithLabelValues("200").Inc()
    w.Write([]byte("ok"))
  }
  reqLatency.Observe(time.Since(start).Seconds())
}

func main() {
  prometheus.MustRegister(reqLatency, reqTotal)
  http.Handle("/metrics", promhttp.Handler())
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8081", nil)
}
