package app

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func CreateApp() *gin.Engine {
	router := gin.Default()

	// Redis
	InitRedis()

	router.NoRoute(ProxyController)

	// Pometheus config
	RegisterMetrics()
	router.Use(PrometheusMiddleware())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}
