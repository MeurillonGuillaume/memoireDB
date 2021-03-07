package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/role"
	"github.com/sirupsen/logrus"
)

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	cfg, err := loadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	node, err := role.NewNodeWithRole(&cfg.Role)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create NodeRole")
	}
	logrus.Infof("I am a cluster %s with name %s and I belong to the cluster %s", node.GetRole(), node.GetName(), node.GetCluster())

	_, err = communication.NewNodeCommunicator(&cfg.Communication)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create node communicator")
	}
	logrus.Info("Created communicator")

	logrus.WithField("signal", <-sigChan).Warn("Received exit signal")
}
