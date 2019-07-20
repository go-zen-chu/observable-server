package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-zen-chu/observable-server/pkg/application"
	"github.com/go-zen-chu/observable-server/pkg/pprof"
	"github.com/go-zen-chu/observable-server/pkg/prometheus"
)

const (
	pprofPort    = "6060"
	exporterPort = "9090"
	appPort      = "8080"
)

func LogIf(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	log.Println("start fibonacci server")

	// handle signal for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}
	// different port for security
	wg.Add(1)
	pprofServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", pprofPort),
		Handler: pprof.Mux(),
	}
	go func() {
		LogIf(pprofServer.ListenAndServe())
		wg.Done()
	}()
	// different port for exporter
	wg.Add(1)
	exporterServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", exporterPort),
		Handler: prometheus.Mux(),
	}
	go func() {
		LogIf(exporterServer.ListenAndServe())
		wg.Done()
	}()
	// handle main application
	wg.Add(1)
	appServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", appPort),
		Handler: application.Mux(),
	}
	go func() {
		LogIf(appServer.ListenAndServe())
		wg.Done()
	}()

	<-sigChan

	log.Println("end fibonacci server")
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	LogIf(appServer.Shutdown(ctx))
	LogIf(exporterServer.Shutdown(ctx))
	LogIf(pprofServer.Shutdown(ctx))
	wg.Wait()
	log.Println("Bye")
}
