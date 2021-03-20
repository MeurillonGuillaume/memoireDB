package shepherd

import (
	"context"

	exCommunication "github.com/MeurillonGuillaume/memoireDB/external/communication"
	inCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
)

type Shepherd interface {
	Run(ctx context.Context, cc []exCommunication.ClientCommunicator, ic inCommunication.NodeCommunicator) error
}
