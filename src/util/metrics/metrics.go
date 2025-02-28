package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
)

const (
	namespace = "link"
	subsystem = "gin-use"
)

// metricsRequestsTotal metrics for request total 计数器（Counter）
var metricsRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "request(s) total",
	},
	[]string{"method", "path"},
)

// metricsRequestsCost metrics for requests cost 累积直方图（Histogram）
var metricsRequestsCost = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_cost",
		Help:      "request(s) cost seconds",
	},
	[]string{"method", "path", "success", "http_code", "business_code", "cost_seconds", "trace_id"},
)

func init() {
	prometheus.MustRegister(metricsRequestsTotal, metricsRequestsCost)
}

// RecordMetrics 记录指标
func RecordMetrics(method, uri string, success bool, httpCode, businessCode int, costSeconds float64, traceId string) {
	metricsRequestsTotal.With(prometheus.Labels{
		"method": method,
		"path":   uri,
	}).Inc()

	metricsRequestsCost.With(prometheus.Labels{
		"method":        method,
		"path":          uri,
		"success":       cast.ToString(success),
		"http_code":     cast.ToString(httpCode),
		"business_code": cast.ToString(businessCode),
		"cost_seconds":  cast.ToString(costSeconds),
		"trace_id":      traceId,
	}).Observe(costSeconds)
}
