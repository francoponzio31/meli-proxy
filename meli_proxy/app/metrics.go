package app

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Defines a counter for HTTP requests
var RequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests received",
	},
	[]string{"path", "method"},
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ignore requests to the /metrics endpoint
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		c.Next()

		path := c.Request.URL.Path
		RequestCount.WithLabelValues(path, c.Request.Method).Inc()
	}
}

func RegisterMetrics() {
	// Register custom metrics to the Prometheus logger
	prometheus.MustRegister(RequestCount)
}
