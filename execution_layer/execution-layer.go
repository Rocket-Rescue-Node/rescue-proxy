package execution_layer

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
)

// Encapsulate a EL client and present a facade with only the data we care about accessing
type ExecutionLayer struct {
	rp *rocketpool.RocketPool
}

func NewExecutionLayer(ecUrl *url.URL, rocketStorageAddr string) (*ExecutionLayer, error) {
	out := &ExecutionLayer{}
	client, err := ethclient.Dial(ecUrl.String())
	if err != nil {
		return nil, err
	}
	out.rp, err = rocketpool.NewRocketPool(client, common.HexToAddress(rocketStorageAddr))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// TODO functions that match:
// func (e *ExecutionLayer) name() (foo, error) {}
// that abstract access to the EL, using rocketpool-go
