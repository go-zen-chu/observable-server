package prometheus

import (
	"net/http"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// interface provides recording metrics
type MetricsExporter interface {
	prom.Collector
	ConstFibNumValue(v int)
	Fib1Num(v int)
	Fib1Value(v int)
	Fib2Num(v int)
	Fib2Value(v int)
}

// Exporter is global metrics collector
var Exporter MetricsExporter

const namespace = "fibnum_observable_server"

type exporter struct {
	constFibNumGauge prom.Gauge
	fib1Num          prom.Gauge
	fib1Value        prom.Gauge
	fib2Num          prom.Gauge
	fib2Value        prom.Gauge
}

func init() {
	Exporter = &exporter{
		constFibNumGauge: prom.NewGauge(prom.GaugeOpts{
			Namespace: namespace,
			Name:      "const_fib_num",
			Help:      "showing const value of fib num",
		}),
		fib1Num: prom.NewGauge(prom.GaugeOpts{
			Namespace: namespace,
			Name:      "fib1",
			Help:      "showing fib1 num",
		}),
		fib1Value: prom.NewGauge(prom.GaugeOpts{
			Namespace: namespace,
			Name:      "fib1_value",
			Help:      "showing value of fib1",
		}),
		fib2Num: prom.NewGauge(prom.GaugeOpts{
			Namespace: namespace,
			Name:      "fib2",
			Help:      "showing fib2 num",
		}),
		fib2Value: prom.NewGauge(prom.GaugeOpts{
			Namespace: namespace,
			Name:      "fib2_value",
			Help:      "showing value of fib2",
		}),
	}
	prom.MustRegister(Exporter)
}

func (e *exporter) Describe(ch chan<- *prom.Desc) {
	ch <- e.constFibNumGauge.Desc()
	ch <- e.fib1Num.Desc()
	ch <- e.fib1Value.Desc()
	ch <- e.fib2Num.Desc()
	ch <- e.fib2Value.Desc()
}

func (e *exporter) Collect(ch chan<- prom.Metric) {
	e.constFibNumGauge.Collect(ch)
	e.fib1Num.Collect(ch)
	e.fib1Value.Collect(ch)
	e.fib2Num.Collect(ch)
	e.fib2Value.Collect(ch)
}

func (e *exporter) ConstFibNumValue(v int) {
	e.constFibNumGauge.Set(float64(v))
}

func (e *exporter) Fib1Num(v int) {
	e.fib1Num.Set(float64(v))
}

func (e *exporter) Fib1Value(v int) {
	e.fib1Value.Set(float64(v))
}

func (e *exporter) Fib2Num(v int) {
	e.fib2Num.Set(float64(v))
}

func (e *exporter) Fib2Value(v int) {
	e.fib2Value.Set(float64(v))
}

// Mux provides endpoint for getting custom metrics of applicaiton
func Mux() *http.ServeMux {
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	return m
}
