package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_request_duration_ms",
	Help:    "Duration of HTTP requests in ms",
	Buckets: prometheus.DefBuckets,
}, []string{"path", "method", "status"})

var http500Errors = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_500_errors",
	Help: "Number of 500 errors",
}, []string{"path", "method"})

func InitMetrics() {
	prometheus.MustRegister(httpDuration, http500Errors)

	log.Info().Msg("Prometheus metrics initialized")
}

func MetricsHttpHandler() http.HandlerFunc {
	return promhttp.Handler().ServeHTTP
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		duration := time.Since(startTime).Milliseconds()
		status := ww.Status()

		routePattern := chi.RouteContext(r.Context()).RoutePattern()

		if status >= 500 {
			http500Errors.WithLabelValues(routePattern, r.Method).Inc()
		}

		httpDuration.WithLabelValues(
			routePattern,
			r.Method,
			strconv.Itoa(status),
		).Observe(float64(duration))
	})
}
