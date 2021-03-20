package shepherd

import (
	"context"

	"github.com/MeurillonGuillaume/memoireDB/external/communication"
	inCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/shared"
	"github.com/sirupsen/logrus"
)

type simpleShepherd struct{}

var _ Shepherd = (*simpleShepherd)(nil)

func (ss *simpleShepherd) Run(ctx context.Context, clientCommunicators []communication.ClientCommunicator, _ inCommunication.NodeCommunicator) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case item := <-shared.CombineChans(ctx, communication.GetCommunicatorChans(clientCommunicators)...):
			logrus.Info(item)
		}
	}
}
