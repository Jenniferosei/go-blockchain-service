package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    RequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "api_requests_total",
            Help: "Total number of API requests",
        },
        []string{"path"},
    )
)

func InitMetrics() {
    prometheus.MustRegister(RequestsTotal)
}
