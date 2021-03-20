package main

import (
	"os"
	"os/signal"
	"syscall"

	externalCommunication "github.com/MeurillonGuillaume/memoireDB/external/communication"
	internalCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/role"
	"github.com/sirupsen/logrus"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	cfg, err := loadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	node, err := role.NewNodeWithRole(&cfg.Role)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create NodeRole")
	}
	logrus.Infof("I am a cluster %s with name %s and I belong to the cluster %s", node.GetRole(), node.GetName(), node.GetCluster())

	_, err = internalCommunication.NewNodeCommunicator(&cfg.InternalCommunication)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create node communicator")
	}
	logrus.Info("Created internal node communicator")

	_, err = externalCommunication.NewClientCommunicators(&cfg.ExternalCommunication)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create client communicator")
	}

	logrus.WithField("signal", <-sigChan).Warn("Received exit signal")
}
