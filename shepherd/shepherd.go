package shepherd

import (
	"context"

	exCommunication "github.com/MeurillonGuillaume/memoireDB/external/communication"
	inCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
)

type Shepherd interface {
	Run(ctx context.Context) error
}

func NewShepherd(ic inCommunication.NodeCommunicator, ec []exCommunication.ClientCommunicator) (s Shepherd, err error) {
	return newSimpleShepherd(ic, ec), nil
}
