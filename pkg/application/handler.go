package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-zen-chu/observable-server/pkg/prometheus"
)

const fibNum = 40

// high computation
func fib1(n int) int {
	if n < 2 {
		return n
	}
	val := fib1(n-1) + fib1(n-2)
	prometheus.Exporter.Fib1Num(n)
	prometheus.Exporter.Fib1Value(val)
	return val
}

// low computation
func fib2(n int) int {
	if n < 2 {
		return n
	}
	p2 := 1
	p1 := 1
	i := 3 // start idx from 3
	for ; i < n; i++ {
		tmp := p1
		p1 = p2 + p1
		p2 = tmp
		prometheus.Exporter.Fib2Num(i)
		prometheus.Exporter.Fib2Value(p1)
		// sleep for getting metric
		time.Sleep(50 * time.Millisecond)
	}
	prometheus.Exporter.Fib2Num(i)
	prometheus.Exporter.Fib2Value(p2 + p1)
	return p2 + p1
}

func fib1Handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := fib1(fibNum)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "fib1: n=%d, result=%d, time=%s", fibNum, res, elapsed)
}

func fib2Handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := fib2(fibNum)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "fib2: n=%d, result=%d, time=%s", fibNum, res, elapsed)
}

// Mux provides application route handling
func Mux() *http.ServeMux {
	m := http.NewServeMux()
	prometheus.Exporter.ConstFibNumValue(fibNum)
	m.HandleFunc("/fib1", fib1Handler)
	m.HandleFunc("/fib2", fib2Handler)
	return m
}
