package shepherd

import (
	"context"

	exCommunication "github.com/MeurillonGuillaume/memoireDB/external/communication"
	inCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/shared"
	"github.com/sirupsen/logrus"
)

type simpleShepherd struct {
	internalCommunicator  inCommunication.NodeCommunicator
	externalCommunicators []exCommunication.ClientCommunicator
}

func newSimpleShepherd(ic inCommunication.NodeCommunicator, ecs []exCommunication.ClientCommunicator) Shepherd {
	return &simpleShepherd{
		internalCommunicator:  ic,
		externalCommunicators: ecs,
	}
}

var _ Shepherd = (*simpleShepherd)(nil)

func (ss *simpleShepherd) Run(ctx context.Context) error {
	for _, ec := range ss.externalCommunicators {
		go ec.Run(ctx)
	}

	combinedChan := shared.CombineChans(ctx, exCommunication.GetCommunicatorChans(ss.externalCommunicators)...)
	for {
		select {
		case <-ctx.Done():
			return nil
		case item := <-combinedChan:
			logrus.WithField("op", item).Info("Received an operation")
		}
	}
}
