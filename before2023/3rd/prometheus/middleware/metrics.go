package middleware

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricsMiddleware struct {
	opsProcessed *prometheus.CounterVec
}

func NewMetricsMiddleware() *MetricsMiddleware {
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"method", "path", "statuscode"})
	return &MetricsMiddleware{
		opsProcessed: opsProcessed,
	}
}

// Metrics middleware to collect metrics from http requests
func (lm *MetricsMiddleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		wi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}
		next.ServeHTTP(wi, r)

		lm.opsProcessed.With(prometheus.Labels{"method": r.Method, "path": r.RequestURI, "statuscode": strconv.Itoa(wi.statusCode)}).Inc()
	})
}

// responseWriterInterceptor is a simple wrapper to intercept set data on a
// ResponseWriter.
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(p []byte) (int, error) {
	return w.ResponseWriter.Write(p)
}

func (w *responseWriterInterceptor) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("type assertion failed http.ResponseWriter not a http.Hijacker")
	}
	return h.Hijack()
}

func (w *responseWriterInterceptor) Flush() {
	f, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}

	f.Flush()
}
