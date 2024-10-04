package main

import (
	"fmt"
	"ministry/config"
	"ministry/internal/pkg/app"
	"ministry/internal/storage"
	"ministry/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg := config.Init()

	log := logger.Init(&cfg.Logger)

	a := app.New(cfg, log)

	go a.Run(cfg)

	fmt.Println(".....SERVER IS RUNNING.....")

	waitSignal()
}

func waitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	defer storage.CloseStorage()
	fmt.Println(".....SERVER IS SHUTTING DOWN AND DB IS CLOSED....")
}
