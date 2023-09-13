package metrics

import (
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
)

type maps struct {
	solo sync.Map
	rp   sync.Map
}

var currentIdx uint64
var epochs [4]maps

var registry *MetricsRegistry

func OnHead(epoch uint64) {

	// Bunch 2 epochs into the same bucket
	// Even epochs do, which means currendIdx advances once for every 2 epochs
	// Divide by two before modulo four in order to have a ring buffer of 8 epochs,
	// two per bucket.
	epoch /= 2

	newCurrentIdx := epoch % 4
	if newCurrentIdx == currentIdx {
		return
	}

	registry.Counter("head_advanced").Inc()

	// Clear the new index
	epochs[newCurrentIdx] = maps{}
	// atomic provides us with a memory fence in go 1.19+, which means the above assignments
	// are always observable on other goroutines when the newCurrentIdx is loaded
	atomic.StoreUint64(&currentIdx, newCurrentIdx)
}

func ObserveSoloValidator(node common.Address, pubkey rptypes.ValidatorPubkey) {
	registry.Counter("observed_solo_validator").Inc()
	var nodeMap *sync.Map = &sync.Map{}

	epoch := &epochs[atomic.LoadUint64(&currentIdx)]

	iface, _ := epoch.solo.LoadOrStore(node, nodeMap)
	nodeMap = iface.(*sync.Map)
	_, _ = nodeMap.LoadOrStore(pubkey, struct{}{})
}

func ObserveValidator(node common.Address, pubkey rptypes.ValidatorPubkey) {
	registry.Counter("observed_validator").Inc()
	var nodeMap *sync.Map = &sync.Map{}

	epoch := &epochs[atomic.LoadUint64(&currentIdx)]

	iface, _ := epoch.rp.LoadOrStore(node, nodeMap)
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
	epoch.rp.Range(func(key, value any) bool {
		out += 1
		return true
	})

	return
}

func PreviousEpochValidators() (out float64) {

	epoch := &epochs[previousEpochIdx()]
	epoch.rp.Range(func(key, value any) bool {
		m := value.(*sync.Map)
		m.Range(func(key2, value2 any) bool {
			out += 1
			return true
		})
		return true
	})

	return
}

func PreviousEpochSoloNodes() (out float64) {

	epoch := &epochs[previousEpochIdx()]
	epoch.solo.Range(func(key, value any) bool {
		out += 1
		return true
	})

	return
}

func PreviousEpochSoloValidators() (out float64) {

	epoch := &epochs[previousEpochIdx()]
	epoch.solo.Range(func(key, value any) bool {
		m := value.(*sync.Map)
		m.Range(func(key2, value2 any) bool {
			out += 1
			return true
		})
		return true
	})

	return
}

func CurrentIdx() (out float64) {

	return float64(currentIdx)
}

func PreviousIdx() (out float64) {

	return float64(previousEpochIdx())
}

func InitEpochMetrics() {
	registry = NewMetricsRegistry("epoch")
	registry.GaugeFunc("validators_seen", PreviousEpochValidators)
	registry.GaugeFunc("nodes_seen", PreviousEpochNodes)
	registry.GaugeFunc("validators_seen_solo", PreviousEpochSoloValidators)
	registry.GaugeFunc("nodes_seen_solo", PreviousEpochSoloNodes)
	registry.GaugeFunc("current_idx", CurrentIdx)
	registry.GaugeFunc("previous_idx", PreviousIdx)
}
