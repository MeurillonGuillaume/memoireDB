package main

import (
	"github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/role"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	node, err := role.NewNodeWithRole(&cfg.Role)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create NodeRole")
	}
	logrus.Infof("I am a cluster %s", node.Role())

	_, err = communication.NewNodeCommunicator(&cfg.Communication)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create node communicator")
	}
	logrus.Info("Created communicator")
}
