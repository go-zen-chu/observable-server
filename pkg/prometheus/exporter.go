package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    exampleCount = prometheus.NewCounter(prometheus.CounterOpts{
        Namespace: namespace,
        Name:      "example_count",
        Help:      "example counter help",
    })
    exampleGauge = prometheus.NewGauge(prometheus.GaugeOpts{
        Namespace: namespace,
        Name:      "example_gauge",
        Help:      "example gauge help",
	})
)

type customCollector struct {}

func (c customCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- exampleCount.Desc()
    ch <- exampleGauge.Desc()
}

func (c customCollector) Collect(ch chan<- prometheus.Metric) {
    ch <- prometheus.MustNewConstMetric(
        exampleCount.Desc(),
        prometheus.CounterValue,
        float64(exampleValue),
    )
    ch <- prometheus.MustNewConstMetric(
        exampleGauge.Desc(),
        prometheus.GaugeValue,
        float64(exampleValue),
    )
}

// PrometheusMux provides endpoint for getting custom metrics of applicaiton
func PrometheusMux() *http.ServerMux {
	m := http.NewServeMux()
	c := &customCollector{}
	prometheus.MustRegister(c)

	m.HandleFunc("/metrics", promhttp.Handler())
	return m
}
