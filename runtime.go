package main

import (
	"context"
	"fmt"

	externalCommunication "github.com/MeurillonGuillaume/memoireDB/external/communication"
	internalCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/datastore"
	"github.com/MeurillonGuillaume/memoireDB/shepherd"
	"github.com/sirupsen/logrus"
)

/*
Run is used to create a memoireDB instance which can be called from any source.

External clients can import the run function and create a local instance.

The client will live as long as the context is alive.
*/
func Run(ctx context.Context) (err error) {
	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("could not load configuration: %w", err)
	}

	ic, err := internalCommunication.NewNodeCommunicator(&cfg.InternalCommunication)
	if err != nil {
		return fmt.Errorf("could not create cluster internal node communicator: %w", err)
	}
	logrus.Info("Created cluster internal node communicator")

	ecs, err := externalCommunication.NewClientCommunicators(&cfg.ExternalCommunication)
	if err != nil {
		return fmt.Errorf("could not create external communicator: %w", err)
	}
	defer func() {
		for _, ec := range ecs {
			if err := ec.Close(); err != nil {
				logrus.WithError(err).Error("Could not close client communicator")
			}
		}
	}()

	ds, err := datastore.NewDatastore(cfg.Datastore)
	if err != nil {
		return fmt.Errorf("could not create datastore: %w", err)
	}
	defer func() {
		if err := ds.Close(); err != nil {
			logrus.WithError(err).Error("Could not close datastore properly")
		}
	}()

	shepherd, err := shepherd.NewShepherd(ic, ecs, ds)
	if err != nil {
		return fmt.Errorf("could not create shepherd: %w", err)
	}

	if err := shepherd.Run(ctx); err != nil {
		return fmt.Errorf("could no longer keep shepherd alive: %w", err)
	}

	logrus.WithField("signal", ctx.Err()).Warn("Received exit signal")
	return nil
}
