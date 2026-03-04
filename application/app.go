package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	// Business metric
	opsProcessed prometheus.Counter

	// HTTP metrics
	httpRequestsTotal    *prometheus.CounterVec
	httpRequestDuration  *prometheus.HistogramVec
	httpRequestsInFlight prometheus.Gauge
}

func newMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		opsProcessed: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
			Help: "Total number of processed operations",
		}),

		httpRequestsTotal: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),

		httpRequestDuration: promauto.With(reg).NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request latency",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "path"},
		),

		httpRequestsInFlight: promauto.With(reg).NewGauge(
			prometheus.GaugeOpts{
				Name: "http_requests_in_flight",
				Help: "Current number of in-flight HTTP requests",
			},
		),
	}

	return m
}

func main() {
	reg := prometheus.NewRegistry()

	// Register runtime & process metrics
	reg.MustRegister(
		prometheus.NewGoCollector(),
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
	)

	m := newMetrics(reg)

	mux := http.NewServeMux()

	// Example business endpoint
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		m.httpRequestsInFlight.Inc()
		defer m.httpRequestsInFlight.Dec()

		// Simulate work
		time.Sleep(100 * time.Millisecond)

		m.opsProcessed.Inc()

		status := http.StatusOK
		w.WriteHeader(status)
		w.Write([]byte("ok"))

		duration := time.Since(start).Seconds()

		m.httpRequestsTotal.
			WithLabelValues(r.Method, "/process", http.StatusText(status)).
			Inc()

		m.httpRequestDuration.
			WithLabelValues(r.Method, "/process").
			Observe(duration)
	})

	mux.Handle("/process", handler)
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	server := &http.Server{
		Addr:    ":2112",
		Handler: mux,
	}

	log.Println("Server running on :2112")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
