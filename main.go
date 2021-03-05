package main

import "github.com/sirupsen/logrus"

func main() {
	cfg, err := loadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	logrus.Info(cfg)
}
