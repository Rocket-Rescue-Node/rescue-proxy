package dataprovider

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
)

type NodeInfo struct {
	InSmoothingPool bool
	FeeDistributor  common.Address
}

type DataProvider interface {
	GetAllNodes(opts *bind.CallOpts) (map[common.Address]*NodeInfo, error)
	GetAllMinipools(nodeMap map[common.Address]*NodeInfo, opts *bind.CallOpts) (map[common.Address][]rptypes.ValidatorPubkey, error)
	GetAllOdaoNodes(opts *bind.CallOpts) ([]common.Address, error)
}
