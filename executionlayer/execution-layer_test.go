package executionlayer

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/dataprovider"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
	"go.uber.org/zap/zaptest"
)

const rocketSmoothingPool = "0x000000000000000000000000d4e96ef8eee8678dbff4d535e033ed1a4f7605b7"
const rocketTokenRETH = "0x000000000000000000000000ae78736cd615f374d3085123a210448e74fc6393"

const eip1271SmartContractValidSignerAddress = "0x1234567890123456789012345678901234567890"
const eip1271SmartContractInvalidSignerAddress = "0x1234567890123456789012345678901234567891"
const eip1271ValidSignature = "0x0000000000000000000000000000000000000000000000000000000000000456"
const eip1271InvalidSignature = "0x0000000000000000000000000000000000000000000000000000000000000789"

type elTest struct {
	t  *testing.T
	m  *mockEC
	ec *CachingExecutionLayer
}

func setup(t *testing.T, m *mockEC) *elTest {
	_, err := metrics.Init("ec_test_" + t.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(metrics.Deinit)

	out := &elTest{
		t: t,
		m: m,
	}

	if err != nil {
		t.Fatal(err)
	}
	out.ec = &CachingExecutionLayer{
		Logger:  zaptest.NewLogger(t),
		Context: t.Context(),
	}
	out.ec.DataProvider = out.m
	return out
}

type mockNode struct {
	addr        common.Address
	inSP        bool
	minipools   int
	minipoolMap map[rptypes.ValidatorPubkey]interface{}
}

type mockEC struct {
	t        *testing.T
	nodes    []*mockNode
	daoNodes []*mockNode
}

var _ dataprovider.DataProvider = &mockEC{}

func (m *mockEC) GetAllNodes(opts *bind.CallOpts) (map[common.Address]*dataprovider.NodeInfo, error) {
	out := make(map[common.Address]*dataprovider.NodeInfo)
	for _, node := range m.nodes {
		out[node.addr] = &dataprovider.NodeInfo{
			InSmoothingPool: node.inSP,
			FeeDistributor:  node.addr,
		}
	}
	return out, nil
}

func pubkeyFromMinipool(addr common.Address) string {
	// Simply left-pad with a char out to the desired length
	return fmt.Sprintf("f0f0%092s", addr.String()[2:])
}

func (m *mockEC) GetAllMinipools(nodes map[common.Address]*dataprovider.NodeInfo, opts *bind.CallOpts) (map[common.Address][]rptypes.ValidatorPubkey, error) {
	out := make(map[common.Address][]rptypes.ValidatorPubkey)
	for _, node := range m.nodes {
		node.minipoolMap = make(map[rptypes.ValidatorPubkey]interface{})
		for range node.minipools {
			pubkey := pubkeyFromMinipool(node.addr)
			h, err := hex.DecodeString(pubkey)
			if err != nil {
				return nil, errors.New("unexpected invalid pubkey")
			}
			out[node.addr] = append(out[node.addr], rptypes.BytesToValidatorPubkey(h))
			node.minipoolMap[rptypes.BytesToValidatorPubkey(h)] = struct{}{}
		}
	}
	return out, nil
}

func (m *mockEC) GetAllOdaoNodes(opts *bind.CallOpts) ([]common.Address, error) {
	out := make([]common.Address, len(m.daoNodes))
	for i, node := range m.daoNodes {
		out[i] = node.addr
	}
	return out, nil
}

func (m *mockEC) GetREthAddress() common.Address {
	return common.HexToAddress(rocketTokenRETH)
}

func (m *mockEC) GetSmoothingPoolAddress() common.Address {
	return common.HexToAddress(rocketSmoothingPool)
}

func (m *mockEC) StakewiseFeeRecipient(opts *bind.CallOpts, address common.Address) (*common.Address, error) {
	return nil, nil
}

func (m *mockEC) ValidateEIP1271(opts *bind.CallOpts, dataHash common.Hash, signature []byte, address common.Address) (bool, error) {
	if bytes.Equal(address[:], common.FromHex(eip1271SmartContractInvalidSignerAddress)) {
		return false, dataprovider.ErrNoData
	}
	if bytes.Equal(signature, common.FromHex(eip1271ValidSignature)) {
		return true, nil
	}
	return false, nil
}

func (m *mockEC) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	if number == nil {
		return &types.Header{Number: big.NewInt(1111)}, nil
	}
	return &types.Header{
		Number: number,
	}, nil
}

func (m *mockEC) RefreshAddresses(opts *bind.CallOpts) error {
	return nil
}

func TestELStartStop(t *testing.T) {
	et := setup(t, &mockEC{t,
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}
}

func TestELGetRPInfoMissing(t *testing.T) {
	et := setup(t, &mockEC{t,
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	rpinfo, err := et.ec.GetRPInfo(rptypes.BytesToValidatorPubkey([]byte{0x01}))
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if rpinfo != nil {
		t.Fatal("unexpected rp info", rpinfo)
	}
}

func TestELGetRPInfo(t *testing.T) {
	mockEC := &mockEC{t,
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	}
	et := setup(t, mockEC)
	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	for _, node := range mockEC.nodes {
		for pubkey := range node.minipoolMap {
			rpinfo, err := et.ec.GetRPInfo(pubkey)
			if err != nil {
				t.Fatal("unexpected error", err)
			}
			if rpinfo == nil {
				t.Fatal("expected rp info", rpinfo)
			}
		}
	}
}

func TestELGetRETHAddress(t *testing.T) {
	et := setup(t, &mockEC{t,
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	reth := et.ec.REthAddress()
	if reth == nil {
		t.Fatal("expected address")
	}

	if !bytes.Equal(reth.Bytes(), common.HexToAddress(rocketTokenRETH).Bytes()) {
		t.Fatal("Expected reth token address to match")
	}

}

func TestELForEaches(t *testing.T) {
	hec := &mockEC{t,
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	}
	et := setup(t, hec)

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	// Check that foreach node now iterates 3x
	nodeCount := 0
	err := et.ec.ForEachNode(func(a common.Address) bool {
		nodeCount++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	if nodeCount != 2 {
		t.Fatalf("Expected 2 nodes in foreach iterator, got: %d", nodeCount)
	}

	// Check that foreach odaonode iterates 2x
	nodeCount = 0
	err = et.ec.ForEachOdaoNode(func(a common.Address) bool {
		nodeCount++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	if nodeCount != 1 {
		t.Fatalf("Expected 1 nodes in odao foreach iterator, got: %d", nodeCount)
	}

	// Remove an odao node
	err = et.ec.cache.Load().removeOdaoNode(common.HexToAddress("0x01"))
	if err != nil {
		t.Fatal(err)
	}
	nodeCount = 0
	err = et.ec.ForEachOdaoNode(func(a common.Address) bool {
		nodeCount++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	if nodeCount != 1 {
		t.Fatalf("Expected 2 nodes in odao foreach iterator, got: %d", nodeCount)
	}

	// Get rpinfo as with the sql cache
	found := 0
	for _, n := range hec.nodes {
		for pk := range n.minipoolMap {
			ri, err := et.ec.GetRPInfo(pk)
			if err != nil {
				t.Fatal(err)
			}
			found++

			if ri.NodeAddress.String() != n.addr.String() {
				t.Fatal("Mismatched node addresses")
			}

			if n.inSP && ri.ExpectedFeeRecipient.String() != common.HexToAddress(rocketSmoothingPool).String() {
				t.Fatal("expected node to have fee recipient set to sp")
			}
		}
	}

	// Direct getnodeinfo for missing key returns an error
	_, err = et.ec.cache.Load().getNodeInfo(common.HexToAddress("0xff"))
	if err == nil {
		t.Fatal("expected error")
	} else if !errors.Is(err, &NotFoundError{}) {
		t.Fatal("unexpected error", err)
	}
}

func TestValidateEIP1271(t *testing.T) {
	et := setup(t, &mockEC{t,
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
		},
		[]*mockNode{
			{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	// Test cases
	testCases := []struct {
		name           string
		dataHash       [32]byte
		signature      []byte
		address        common.Address
		expectedResult bool
		expectedError  bool
	}{
		{
			name:           "Valid signature",
			dataHash:       [32]byte{0x00},
			signature:      common.FromHex(eip1271ValidSignature),
			address:        common.HexToAddress(eip1271SmartContractValidSignerAddress),
			expectedResult: true,
			expectedError:  false,
		},
		{
			name:           "Invalid signature",
			dataHash:       [32]byte{0x00},
			signature:      common.FromHex(eip1271InvalidSignature),
			address:        common.HexToAddress(eip1271SmartContractValidSignerAddress),
			expectedResult: false,
			expectedError:  false,
		},
		{
			name:           "Invalid contract",
			dataHash:       [32]byte{0x00},
			signature:      common.FromHex(eip1271InvalidSignature),
			address:        common.HexToAddress(eip1271SmartContractInvalidSignerAddress),
			expectedResult: false,
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			result, err := et.ec.ValidateEIP1271(ctx, tc.dataHash, tc.signature, tc.address)

			if tc.expectedError {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if result != tc.expectedResult {
					t.Errorf("Expected result %v, but got %v", tc.expectedResult, result)
				}
			}
		})
	}
}
