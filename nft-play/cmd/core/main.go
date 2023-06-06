package main

import (
	"context"
	"metamask-auth/config"
	"metamask-auth/internal/program"
	appCloser "metamask-auth/pkg/app_closer"
	"metamask-auth/pkg/logger"
	"metamask-auth/pkg/storage"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	cfg := config.CreateConfig()
	mainlog := logger.FromEnv("[main]")
	closer := appCloser.InitCloser(mainlog)

	db, err := storage.InitDBConnection(ctx, cfg, closer)
	if err != nil {
		mainlog.Panic("pgConnector creation error")
	}

	coreProgram := program.NewCoreProgram(cfg, logger.FromEnv("[coreProgram]"), closer, db)
	go coreProgram.Run()

	select {
	case <-ctx.Done():
		go func() {
			closer.CloseAll()
		}()
		stop()
		os.Exit(0)
	}
}
