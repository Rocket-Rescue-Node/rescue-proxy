package executionlayer

import (
	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/dataprovider"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
)

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "Key not found in cache"
}

// cachedValidator maps a pubkey to its owning node and whether it belongs to a megapool.
type cachedValidator struct {
	NodeAddress common.Address
	IsMegapool  bool
}

type Cache interface {
	init() error
	getMinipoolNode(rptypes.ValidatorPubkey) (cachedValidator, error)
	addMinipoolNode(rptypes.ValidatorPubkey, common.Address) error
	addMegapoolNode(rptypes.ValidatorPubkey, common.Address) error
	getNodeInfo(common.Address) (*dataprovider.NodeInfo, error)
	addNodeInfo(common.Address, *dataprovider.NodeInfo) error
	forEachNode(ForEachNodeClosure) error
	addOdaoNode(common.Address) error
	removeOdaoNode(common.Address) error
	forEachOdaoNode(ForEachNodeClosure) error
}
