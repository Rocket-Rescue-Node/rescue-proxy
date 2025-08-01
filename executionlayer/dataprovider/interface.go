package dataprovider

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	GetREthAddress() common.Address
	GetSmoothingPoolAddress() common.Address
	StakewiseFeeRecipient(opts *bind.CallOpts, address common.Address) (*common.Address, error)
	ValidateEIP1271(opts *bind.CallOpts, dataHash common.Hash, signature []byte, address common.Address) (bool, error)

	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	RefreshAddresses(opts *bind.CallOpts) error
}
