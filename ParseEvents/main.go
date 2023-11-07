package main

import (
	"ParseEvents/config"
	"ParseEvents/contracts"
	"context"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	cfg := config.CreateConfig()

	event := contracts.NewEvents(cfg)
	event.Start(ctx)
}
