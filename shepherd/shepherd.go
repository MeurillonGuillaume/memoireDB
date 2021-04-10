package shepherd

import (
	"context"

	exCommunication "github.com/MeurillonGuillaume/memoireDB/external/communication"
	inCommunication "github.com/MeurillonGuillaume/memoireDB/internal/communication"
	"github.com/MeurillonGuillaume/memoireDB/internal/datastore"
)

type Shepherd interface {
	Run(ctx context.Context) error
}

func NewShepherd(
	ic inCommunication.NodeCommunicator,
	ec []exCommunication.ClientCommunicator,
	ds datastore.Store,
) (s Shepherd, err error) {
	return newSimpleShepherd(ic, ec, ds), nil
}
