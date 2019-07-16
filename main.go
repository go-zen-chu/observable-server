package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-zen-chu/observable-server/pkg/application"
	"github.com/go-zen-chu/observable-server/pkg/pprof"
	"github.com/go-zen-chu/observable-server/pkg/prometheus"
)

const (
	pprofPort    = "6060"
	exporterPort = "9090"
	appPort      = "8080"
)

func main() {
	log.Println("start fibonacci server")
	wg := &sync.WaitGroup{}
	// different port for security
	wg.Add(1)
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%s", pprofPort), pprof.PProfMux())
		wg.Done()
	}()
	// different port for exporter
	wg.Add(1)
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%s", exporterPort), prometheus.PrometheusMux())
		wg.Done()
	}()
	// handle main application
	http.ListenAndServe(fmt.Sprintf(":%s", appPort), application.AppMux())

	wg.Wait()
	log.Println("end fibonacci server")
}
