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
	IncFib1Count()
	AddFib1Value(v int)
	IncFib2Count()
	AddFib2Value(v int)
}

// Exporter is global metrics collector
var Exporter MetricsExporter

const namespace = "fibnum_observable_server"

type exporter struct {
	constFibNumGauge prom.Gauge
	fib1Counter      prom.Counter
	fib1ValueCounter prom.Counter
	fib2Counter      prom.Counter
	fib2ValueCounter prom.Counter
}

func init() {
	Exporter = &exporter{
		constFibNumGauge: prom.NewGauge(prom.GaugeOpts{
			Namespace: namespace,
			Name:      "const_fib_num",
			Help:      "showing const value of fib num",
		}),
		fib1Counter: prom.NewCounter(prom.CounterOpts{
			Namespace: namespace,
			Name:      "fib1",
			Help:      "showing fib1 num",
		}),
		fib1ValueCounter: prom.NewCounter(prom.CounterOpts{
			Namespace: namespace,
			Name:      "fib1_value",
			Help:      "showing value of fib1",
		}),
		fib2Counter: prom.NewCounter(prom.CounterOpts{
			Namespace: namespace,
			Name:      "fib2",
			Help:      "showing fib2 num",
		}),
		fib2ValueCounter: prom.NewCounter(prom.CounterOpts{
			Namespace: namespace,
			Name:      "fib2_value",
			Help:      "showing value of fib2",
		}),
	}
	prom.MustRegister(Exporter)
}

func (e *exporter) Describe(ch chan<- *prom.Desc) {
	ch <- e.constFibNumGauge.Desc()
	ch <- e.fib1Counter.Desc()
	ch <- e.fib1ValueCounter.Desc()
	ch <- e.fib2Counter.Desc()
	ch <- e.fib2ValueCounter.Desc()
}

func (e *exporter) Collect(ch chan<- prom.Metric) {
	e.constFibNumGauge.Collect(ch)
	e.fib1Counter.Collect(ch)
	e.fib1ValueCounter.Collect(ch)
	e.fib2Counter.Collect(ch)
	e.fib2ValueCounter.Collect(ch)
}

func (e *exporter) ConstFibNumValue(v int) {
	e.constFibNumGauge.Set(float64(v))
}

func (e *exporter) IncFib1Count() {
	e.fib1Counter.Inc()
}

func (e *exporter) AddFib1Value(v int) {
	e.fib1ValueCounter.Add(float64(v))
}

func (e *exporter) IncFib2Count() {
	e.fib2Counter.Inc()
}

func (e *exporter) AddFib2Value(v int) {
	e.fib2ValueCounter.Add(float64(v))
}

// Mux provides endpoint for getting custom metrics of applicaiton
func Mux() *http.ServeMux {
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	return m
}
