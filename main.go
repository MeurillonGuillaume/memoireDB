package main

import (
	"context"
	"os/signal"
	"syscall"

	externalCommunication "github.com/MeurillonGuillaume/memoireDB/external/communication"
	internalCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/role"
	"github.com/MeurillonGuillaume/memoireDB/shepherd"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cfg, err := loadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	node, err := role.NewNodeWithRole(&cfg.Role)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create NodeRole")
	}
	logrus.Infof("I am a cluster %s with name %s and I belong to the cluster %s", node.GetRole(), node.GetName(), node.GetCluster())

	ic, err := internalCommunication.NewNodeCommunicator(&cfg.InternalCommunication)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create node communicator")
	}
	logrus.Info("Created internal node communicator")

	ecs, err := externalCommunication.NewClientCommunicators(&cfg.ExternalCommunication)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create client communicator")
	}
	defer func() {
		for _, ec := range ecs {
			if err := ec.Close(); err != nil {
				logrus.WithError(err).Error("Could not close client communicator")
			}
		}
	}()

	shepherd, err := shepherd.NewShepherd(ic, ecs)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create shepherd")
	}

	if err := shepherd.Run(ctx); err != nil {
		logrus.WithError(err).Error("Could not execute shepherd")
	}

	logrus.Warn("Received exit signal")
}
