package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dpcat237/afs/config"
	"github.com/dpcat237/afs/internal/controller"
	"github.com/dpcat237/afs/internal/handler"
	"github.com/dpcat237/afs/internal/interfaces"
	"github.com/dpcat237/afs/internal/logger"
	"github.com/dpcat237/afs/internal/service/router"
)

func main() {
	cfg, err := config.LoadData()
	if err != nil {
		panic("Error initializing configuration: " + err.Error())
	}
	lgr, err := logger.New()
	if err != nil {
		panic("Error initializing logger: " + err.Error())
	}

	// Init handlers
	hndsCll := handler.NewCollector()
	// Init router controllers.
	cnts := controller.NewCollector(lgr, hndsCll)

	// Create router manager.
	rtrMng := router.NewManager(cnts, lgr, cfg.PortHTTP)

	rtrMng.Start()
	lgr.Infof("HTTP router started on port %s", cfg.PortHTTP)

	gracefulShutdown(lgr, rtrMng)
}

// gracefulShutdown stops router connections after receiving system notification.
func gracefulShutdown(lgr interfaces.Logger, rtrMng router.Manager) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	close(c)

	lgr.Info("Service stopping")
	rtrMng.Shutdown(context.Background())
	lgr.Info("Service stopped")
}
