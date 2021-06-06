package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := Run(ctx); err != nil {
		logrus.WithError(err).Fatal("MemoireDB runtime failed")
	}
}
