package metrics

import (
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
)

var currentIdx uint64
var epochs [4]sync.Map

var registry *MetricsRegistry

func NewEpoch(epoch uint64) {
	newCurrentIdx := epoch % 4
	// Clear the new index
	epochs[newCurrentIdx] = sync.Map{}
	// atomic provides us with a memory fence in go 1.19+, which means the above assignments
	// are always observable on other goroutines when the newCurrentIdx is loaded
	atomic.StoreUint64(&currentIdx, newCurrentIdx)
}

func ObserveValidator(node common.Address, pubkey rptypes.ValidatorPubkey) {
	var nodeMap *sync.Map = &sync.Map{}

	epoch := &epochs[atomic.LoadUint64(&currentIdx)]

	iface, _ := epoch.LoadOrStore(node, nodeMap)
	nodeMap = iface.(*sync.Map)
	_, _ = nodeMap.LoadOrStore(pubkey, struct{}{})
}

func previousEpochIdx() uint64 {

	currentIdx := atomic.LoadUint64(&currentIdx)
	if currentIdx == 0 {
		return 3
	}
	return currentIdx - 1
}

func PreviousEpochNodes() (out float64) {

	epoch := &epochs[previousEpochIdx()]
	epoch.Range(func(key, value any) bool {
		out += 1
		return true
	})

	return
}

func PreviousEpochValidators() (out float64) {

	epoch := &epochs[previousEpochIdx()]
	epoch.Range(func(key, value any) bool {
		m := value.(*sync.Map)
		m.Range(func(key2, value2 any) bool {
			out += 1
			return true
		})
		return true
	})

	return
}

func InitEpochMetrics() {
	registry = NewMetricsRegistry("epoch")
	registry.GaugeFunc("validators_seen", PreviousEpochValidators)
	registry.GaugeFunc("nodes_seen", PreviousEpochNodes)
}
