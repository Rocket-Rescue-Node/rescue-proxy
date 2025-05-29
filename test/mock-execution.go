package test

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"

	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer"
)

type MockExecutionLayer struct {
	nodes     []*executionlayer.RPInfo
	odaoNodes []common.Address
	VMap      map[rptypes.ValidatorPubkey]*executionlayer.RPInfo
	REth      common.Address
}

func NewMockExecutionLayer(numNodes int, numOdaoNodes int, numValidators int, seed string) *MockExecutionLayer {
	hash := md5.Sum([]byte(seed))
	// Use the low 8 bytes as the seed for rand
	seedInt := binary.LittleEndian.Uint64(hash[len(hash)-8:])
	chaos := rand.NewSource(int64(seedInt))
	gen := rand.New(chaos)

	out := new(MockExecutionLayer)
	out.nodes = make([]*executionlayer.RPInfo, 0, numNodes)
	out.odaoNodes = make([]common.Address, 0, numOdaoNodes)
	out.VMap = make(map[rptypes.ValidatorPubkey]*executionlayer.RPInfo, numValidators)

	// Create a fake rETH address
	out.REth = randAddress(gen)

	// Generate numNodes random info
	for i := 0; i < numNodes; i++ {
		fr := randAddress(gen)
		info := &executionlayer.RPInfo{
			NodeAddress:          randAddress(gen),
			ExpectedFeeRecipient: &fr,
		}

		out.nodes = append(out.nodes, info)
	}

	// Generate numOdaoNodes random addresses
	for i := 0; i < numOdaoNodes; i++ {
		out.odaoNodes = append(out.odaoNodes, randAddress(gen))
	}

	// Generate numValidators random validators
	for i := 0; i < numValidators; i++ {
		// Pick a random node
		info := out.nodes[gen.Int31n(int32(numNodes))]
		out.VMap[randPubkey(gen)] = info
	}

	return out

}

func (m *MockExecutionLayer) ForEachNode(c executionlayer.ForEachNodeClosure) error {
	for _, node := range m.nodes {
		if !c(node.NodeAddress) {
			break
		}
	}

	return nil
}

func (m *MockExecutionLayer) ForEachOdaoNode(c executionlayer.ForEachNodeClosure) error {
	for _, node := range m.odaoNodes {
		if !c(node) {
			break
		}
	}

	return nil
}

func (m *MockExecutionLayer) GetRPInfo(k rptypes.ValidatorPubkey) (*executionlayer.RPInfo, error) {
	out, ok := m.VMap[k]
	if !ok {
		return nil, nil
	}
	return out, nil
}

func (m *MockExecutionLayer) REthAddress() *common.Address {
	return &m.REth
}

func (m *MockExecutionLayer) ValidateEIP1271(ctx context.Context, dataHash common.Hash, signature []byte, address common.Address) (bool, error) {
	return true, nil
}

func (m *MockExecutionLayer) StakewiseFeeRecipient(ctx context.Context, address common.Address) (*common.Address, error) {
	return nil, nil
}
