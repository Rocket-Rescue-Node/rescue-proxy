package executionlayer

import (
	"bytes"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap/zaptest"
	"golang.org/x/exp/slices"
)

var upgrader = websocket.Upgrader{}

const rocketStorage = "0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46"
const rocketNodeManager = "0x00000000000000000000000089f478e6cc24f052103628f36598d4c14da3d287"
const rocketNodeDistributorFactory = "0x000000000000000000000000e228017f77b3e0785e794e4c0a8a6b935bb4037c"
const rocketMinipoolManager = "0x0000000000000000000000006d010c43d4e96d74c422f2e27370af48711b49bf"
const rocketDAONodeTrusted = "0x000000000000000000000000b8e783882b11Ff4f6Cef3C501EA0f4b960152cc9"
const rocketSmoothingPool = "0x000000000000000000000000d4e96ef8eee8678dbff4d535e033ed1a4f7605b7"
const rocketTokenRETH = "0x000000000000000000000000ae78736cd615f374d3085123a210448e74fc6393"

//go:embed test-data/block-by-number.txt
var blockByNumberFmt string

//go:embed test-data/call-result-fmt.txt
var callResultFmt string

//go:embed test-data/rocket-node-manager-abi.txt
var rocketNodeManagerAbi string

//go:embed test-data/rocket-minipool-manager-abi.txt
var rocketMinipoolManagerAbi string

//go:embed test-data/rocket-smoothing-pool-abi.txt
var rocketSmoothingPoolAbi string

//go:embed test-data/rocket-token-reth-abi.txt
var rocketTokenRETHAbi string

//go:embed test-data/rocket-dao-node-trusted-actions-abi.txt
var rocketDAONodeTrustedActionsAbi string

//go:embed test-data/rocket-node-distributor-factory-abi.txt
var rocketNodeDistributorFactoryAbi string

//go:embed test-data/rocket-dao-node-trusted-abi.txt
var rocketDAONodeTrustedAbi string

type mockEC interface {
	Serve(int, []byte) (int, []byte)
}

type elTest struct {
	t  *testing.T
	m  mockEC
	ec *CachingExecutionLayer
}

func (e *elTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		e.t.Fatal(err)
	}
	defer c.Close()
	for {
		mt, data, err := c.ReadMessage()
		if err != nil {
			e.t.Log("mockEC recv err:", err)
			return
		}

		e.t.Logf("mockEC recv: %d - %s\n", mt, string(data))
		mt, data = e.m.Serve(mt, data)
		e.t.Logf("mockEC resp: %d - %s\n", mt, string(data))

		err = c.WriteMessage(mt, data)
		if err != nil {
			e.t.Log("mockEC resp err:", err)
			return
		}
	}
}

func setup(t *testing.T, m mockEC) *elTest {
	_, err := metrics.Init("cc_test_" + t.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(metrics.Deinit)

	out := &elTest{
		t: t,
		m: m,
	}

	s := httptest.NewServer(out)
	url, err := url.Parse(s.URL)
	// replace scheme
	url.Scheme = "ws"
	if err != nil {
		t.Fatal(err)
	}
	out.ec = &CachingExecutionLayer{
		Logger:            zaptest.NewLogger(t),
		ECURL:             url,
		RocketStorageAddr: rocketStorage,
	}
	t.Cleanup(s.Close)
	return out
}

func intToHex(i int) string {
	return fmt.Sprintf("0x%064x", i)
}

func addrToMinipool(idx uint64, addr common.Address) string {
	return fmt.Sprintf("0x%044x%s", idx+1, addr.String()[2+20:])
}

type mockNode struct {
	addr        common.Address
	inSP        bool
	minipools   int
	minipoolMap map[rptypes.ValidatorPubkey]interface{}
}

type happyEC struct {
	t        *testing.T
	nodes    []*mockNode
	daoNodes []*mockNode
}

// Lifted from geth source https://github.com/ethereum/go-ethereum/blob/master/rpc/json.go#L51
type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type call struct {
	To   common.Address `json:"to"`
	From common.Address `json:"from"`
	Data string         `json:"data"`
}

func (e *happyEC) Serve(mt int, data []byte) (int, []byte) {
	m := jsonrpcMessage{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		e.t.Fatal(err)
	}

	var resp string
	switch m.Method {
	case "eth_getBlockByNumber":
		resp = fmt.Sprintf(blockByNumberFmt, m.ID, "0x11af2c8")
	case "eth_subscribe", "eth_unsubscribe":
		resp = fmt.Sprintf(callResultFmt, m.ID, "0x")
	case "eth_call":
		var paramsArray []json.RawMessage
		err := json.Unmarshal(m.Params, &paramsArray)
		if err != nil {
			e.t.Fatal(err)
		}

		if len(paramsArray) == 0 {
			e.t.Fatal("eth call with 0-len params")
		}

		var callMsg call
		err = json.Unmarshal(paramsArray[0], &callMsg)
		if err != nil {
			e.t.Fatal(err)
		}

		if len(callMsg.Data) == 0 {
			e.t.Fatal("eth call with 0-len data")
		}

		switch callMsg.To.String() {
		case rocketStorage:
			selector := callMsg.Data[:10]
			input := callMsg.Data[10:]
			switch selector {
			// Get Address
			case "0x21f8a721":
				switch input {
				case "af00be55c9fb8f543c04e0aa0d70351b880c1bfafffd15b60065a4a50c85ec94":
					e.t.Log("Returning RocketNodeManager address")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketNodeManager)
				case "e9dfec9339b94a131861a58f1bb4ac4c1ce55c7ffe8550e0b6ebcfde87bb012f":
					e.t.Log("Returning RocketMinipoolManager address")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketMinipoolManager)
				case "822231720aef9b264db1d9ca053137498f759c28b243f45c44db1d39d6bce46e":
					e.t.Log("Returning RocketSmoothingPool address")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketSmoothingPool)
				case "e3744443225bff7cc22028be036b80de58057d65a3fdca0a3df329f525e31ccc":
					e.t.Log("Returning RocketTokenRETH address")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketTokenRETH)
				case "0db039c744237e4ce5ef78a0f054ce1d90d4c567f771ca22f1b89eed7a7b901c":
					e.t.Log("Returning RocketDAONodeTrustedActions address")
					resp = fmt.Sprintf(callResultFmt, m.ID, "0x000000000000000000000000029d946f28f93399a5b0d09c879fc8c94e596aeb")
				case "ea051094896ef3b09ab1b794ad5ea695a5ff3906f74a9328e2c16db69d0f3123":
					e.t.Log("Returning RocketNodeDistributorFactory address")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketNodeDistributorFactory)
				case "9a354e1bb2e38ca826db7a8d061cfb0ed7dbd83d241a2cbe4fd5218f9bb4333f":
					e.t.Log("Returning RocketDAONodeTrusted address")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketDAONodeTrusted)

				default:
					e.t.Log("Unhandled GetAddress", input)
				}
			// Get String
			case "0x986e791a":
				switch input {
				case "b665755e7f514adae7d03140292e555a67796a7f6d6193f2b69e1988efc42a7c":
					e.t.Log("Returning RocketNodeManager abi")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketNodeManagerAbi)
				case "8dbcb47ea1b95b945ad8a07ea9995ed9f9f05c9d32b7abf92ab673ab5c0e88f4":
					e.t.Log("Returning RocketMinipoolManager abi")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketMinipoolManagerAbi)
				case "68df011a367b483345047bcd57214093f1a4920f99771f783e52523d2e8c9359":
					e.t.Log("Returning RocketSmoothingPool abi")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketSmoothingPoolAbi)
				case "66fa1687b0fe549b3c17422e7889850e38d08ecd92c902a63818ba19c20be1f8":
					e.t.Log("Returning RocketTokenRETH abi")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketTokenRETHAbi)
				case "b76727ffbb601cda11c24a90a1e005855a7b67865beb7b9fe18c606e85b5bb9f":
					e.t.Log("Returning RocketDAONodeTrustedActions abi")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketDAONodeTrustedActionsAbi)
				case "fa18b901ccc1aac57eee8923beef96533e7f0a4140718a753ba8efe9b253649b":
					e.t.Log("Returning RocketNodeDistributorFactory abi")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketNodeDistributorFactoryAbi)
				case "72496a0f6ba2c8ba96a29f6abeb2147ad89ec172d1a8ce84fd85828fd8475ed4":
					e.t.Log("Returning RocketDAONodeTrusted abi")
					resp = fmt.Sprintf(callResultFmt, m.ID, rocketDAONodeTrustedAbi)
				default:
					e.t.Log("Unhandled GetString", input)
				}
			default:
				e.t.Log("Unhandled rocketStorage selector", selector)
			}

		case common.HexToAddress(rocketNodeManager).String():
			selector := callMsg.Data[:10]
			input := callMsg.Data[10:]
			switch selector {
			// GetNodeCount
			case "0x39bf397e":
				count := len(e.nodes)
				resp = fmt.Sprintf(callResultFmt, m.ID, intToHex(count))
			// GetNodeAt
			case "0xba75d806":
				// Make sure input is in range
				idx, err := hex.DecodeString(input)
				if err != nil {
					e.t.Fatal(err)
				}
				i := big.NewInt(0).SetBytes(idx).Uint64()
				if i >= uint64(len(e.nodes)) {
					e.t.Fatal("Out-of-bounds node requested")
				}
				n := e.nodes[i]
				addrStr := n.addr.String()[2:]
				resp = fmt.Sprintf(callResultFmt, m.ID, "0x000000000000000000000000"+addrStr)
			// GetSmoothingPoolRegistrationState
			case "0xa4cef9dd":
				// The input is an address
				addr := common.HexToAddress(input)
				for _, n := range e.nodes {
					if n.addr != addr {
						continue
					}
					if n.inSP {
						resp = fmt.Sprintf(callResultFmt, m.ID, "0x0000000000000000000000000000000000000000000000000000000000000001")
					} else {
						resp = fmt.Sprintf(callResultFmt, m.ID, "0x0000000000000000000000000000000000000000000000000000000000000000")
					}
					break
				}

			default:
				e.t.Log("Unhandled rocketNodeManager selector", selector)
			}

		case common.HexToAddress(rocketNodeDistributorFactory).String():
			selector := callMsg.Data[:10]
			input := callMsg.Data[10:]
			switch selector {
			// GetProxyAddress(address)
			case "0xfa2a5b01":
				// Use the reverse of the the node address as the fr
				// The input is an address
				addr := common.HexToAddress(input).Bytes()
				slices.Reverse(addr)
				// Return fee recipient address
				resp = fmt.Sprintf(callResultFmt, m.ID, "0x000000000000000000000000"+common.BytesToAddress(addr).String()[2:])
			default:
				e.t.Log("Unhandled rocketNodeDistributorFactory selector", selector)
			}
		case common.HexToAddress(rocketMinipoolManager).String():
			selector := callMsg.Data[:10]
			input := callMsg.Data[10:]
			switch selector {
			// GetNodeMinipoolCount(address)
			case "0x1ce9ec33":
				// The input is an address
				addr := common.HexToAddress(input)
				for _, n := range e.nodes {
					if n.addr != addr {
						continue
					}
					resp = fmt.Sprintf(callResultFmt, m.ID, intToHex(n.minipools))
					break
				}
			// GetNodeMinipoolAt(address,idx)
			case "0x8b300029":
				// The input is an address and idx
				addr := common.HexToAddress(input[:64])
				idx, err := strconv.ParseUint(input[64+48:], 16, 32)
				if err != nil {
					e.t.Fatal(err)
				}
				for _, n := range e.nodes {
					if n.addr != addr {
						continue
					}
					if idx >= uint64(n.minipools) {
						e.t.Fatal("idx exceeds node count", idx, n.minipools)
					}
					// Prepend the minipool number to the node address
					resp = fmt.Sprintf(callResultFmt, m.ID, addrToMinipool(idx, addr))
					break
				}
			// GetMinipoolPubkey(address)
			case "0x3eb535e9":
				// The input is an address
				addr := common.HexToAddress(input)
				// Simply left-pad with a char out to the desired length
				pubkey := fmt.Sprintf("f0f0%092s", addr.String()[2:])
				h, err := hex.DecodeString(pubkey)
				if err != nil {
					e.t.Fatal(err)
				}
				typ, err := abi.NewType("bytes", "bytes", nil)
				if err != nil {
					e.t.Fatal(err)
				}

				args := abi.Arguments{abi.Argument{Type: typ, Name: "blah"}}

				solBytes, err := args.Pack(h)
				if err != nil {
					e.t.Fatal(err)
				}

				// Find the node and add the pubkey to its map, for the tests to reference
				for _, n := range e.nodes {
					if !bytes.HasSuffix(addr.Bytes(), n.addr.Bytes()[10:]) {
						continue
					}
					if n.minipoolMap == nil {
						n.minipoolMap = make(map[rptypes.ValidatorPubkey]interface{})
					}
					pk := rptypes.BytesToValidatorPubkey(h)
					n.minipoolMap[pk] = struct{}{}
				}

				resp = fmt.Sprintf(callResultFmt, m.ID, fmt.Sprintf("0x%s", hex.EncodeToString(solBytes)))

			default:
				e.t.Log("Unhandled rocketMinipoolManager selector", selector)
			}
		case common.HexToAddress(rocketDAONodeTrusted).String():
			selector := callMsg.Data[:10]
			input := callMsg.Data[10:]
			switch selector {
			// GetMemberCount()
			case "0x997072f7":
				count := len(e.daoNodes)
				resp = fmt.Sprintf(callResultFmt, m.ID, intToHex(count))
			// GetMemberAt(uint256)
			case "0xe992c817":
				// Input is just a number
				i, err := strconv.ParseUint(input, 16, 64)
				if err != nil {
					e.t.Fatal(err)
				}

				if i >= uint64(len(e.daoNodes)) {
					e.t.Fatal("index too large for dao nodes")
				}

				n := e.daoNodes[int(i)]
				resp = fmt.Sprintf(callResultFmt, m.ID, "0x000000000000000000000000"+n.addr.String()[2:])
			default:
				e.t.Log("Unhandled rocketDAONodeTrusted selector", selector)
			}

		default:
			e.t.Log("Unhandled contract call", callMsg.To)
		}
	default:
		e.t.Log("Unhandled eth rpc", m.Method)
	}

	return mt, []byte(resp)
}

func TestELStartStop(t *testing.T) {
	et := setup(t, &happyEC{t,
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	errs := make(chan error)
	go func() {
		if err := et.ec.Start(); err != nil {
			errs <- err
		}
		close(errs)
	}()

	// Wait for connection
	<-et.ec.connected

	et.ec.Stop()
	err := <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestSQLELCache(t *testing.T) {
	cachePath, err := os.MkdirTemp("", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	hec := &happyEC{t,
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
			&mockNode{
				addr:      common.HexToAddress("0x01"),
				inSP:      false,
				minipools: 0,
			},
		},
	}
	et := setup(t, hec)

	et.ec.CachePath = cachePath

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	errs := make(chan error)
	go func() {
		if err := et.ec.Start(); err != nil {
			errs <- err
		}
		close(errs)
	}()

	// Wait for connection
	<-et.ec.connected

	addr := common.HexToAddress("0x1234567891232222222212345678912345678900")
	addr2 := common.HexToAddress("0x1234567891255555555555545678912345678900")
	ni := nodeInfo{true, addr2}

	// Add a node to the cache
	err = et.ec.cache.addNodeInfo(addr, &ni)
	if err != nil {
		t.Fatal(err)
	}

	// Restart
	et.ec.Stop()
	err = <-errs
	if err != nil {
		t.Fatal(err)
	}

	metrics.Deinit()

	et = setup(t, &happyEC{t,
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	et.ec.CachePath = cachePath

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	errs = make(chan error)
	go func() {
		if err := et.ec.Start(); err != nil {
			errs <- err
		}
		close(errs)
	}()

	// Wait for connection
	<-et.ec.connected

	// Check that not found errors come back ok
	_, err = et.ec.cache.getNodeInfo(common.HexToAddress("0x0"))
	if err == nil {
		t.Fatal("expected error")
	} else if !strings.EqualFold(err.Error(), "key not found in cache") {
		t.Fatal("unexpected error", err)
	}

	// Check that the node was loaded from disk
	cachedNi, err := et.ec.cache.getNodeInfo(addr)
	if err != nil {
		t.Fatal(err)
	}

	// Check that foreach node now iterates 3x
	nodeCount := 0
	err = et.ec.ForEachNode(func(a common.Address) bool {
		nodeCount++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	if nodeCount != 3 {
		t.Fatalf("Expected 3 nodes in foreach iterator, got: %d", nodeCount)
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

	if nodeCount != 2 {
		t.Fatalf("Expected 2 nodes in odao foreach iterator, got: %d", nodeCount)
	}

	// Remove an odao node
	err = et.ec.cache.removeOdaoNode(common.HexToAddress("0x01"))
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

	// Check that you can get the node addr from the pubkeys
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

	if found == 0 {
		t.Fatal("Didn't find any cached data")
	}

	if cachedNi.feeDistributor.String() != addr2.String() {
		t.Fatalf("unexpected fee recipient from cache: %s expected: %s", cachedNi.feeDistributor.String(), addr2.String())
	}

	et.ec.Stop()
	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestELGetRPInfoMissing(t *testing.T) {
	et := setup(t, &happyEC{t,
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	errs := make(chan error)
	go func() {
		if err := et.ec.Start(); err != nil {
			errs <- err
		}
		close(errs)
	}()

	// Wait for connection
	<-et.ec.connected

	rpinfo, err := et.ec.GetRPInfo(rptypes.BytesToValidatorPubkey([]byte{0x01}))
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if rpinfo != nil {
		t.Fatal("unexpected rp info", rpinfo)
	}

	et.ec.Stop()
	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestELGetRETHAddress(t *testing.T) {
	et := setup(t, &happyEC{t,
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000222222222222222222222222222222"),
				inSP:      false,
				minipools: 0,
			},
		},
	})

	if err := et.ec.Init(); err != nil {
		t.Fatal(err)
	}

	errs := make(chan error)
	go func() {
		if err := et.ec.Start(); err != nil {
			errs <- err
		}
		close(errs)
	}()

	// Wait for connection
	<-et.ec.connected

	reth := et.ec.REthAddress()
	if reth == nil {
		t.Fatal("expected address")
	}

	if !bytes.Equal(reth.Bytes(), common.HexToAddress(rocketTokenRETH).Bytes()) {
		t.Fatal("Expected reth token address to match")
	}

	et.ec.Stop()
	err := <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestELForEaches(t *testing.T) {
	hec := &happyEC{t,
		[]*mockNode{
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000001234567899876543210"),
				inSP:      true,
				minipools: 1,
			},
			&mockNode{
				addr:      common.HexToAddress("0x0000000000000000000002234567899876543210"),
				inSP:      false,
				minipools: 3,
			},
		},
		[]*mockNode{
			&mockNode{
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

	errs := make(chan error)
	go func() {
		if err := et.ec.Start(); err != nil {
			errs <- err
		}
		close(errs)
	}()

	// Wait for connection
	<-et.ec.connected

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
	err = et.ec.cache.removeOdaoNode(common.HexToAddress("0x01"))
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

	// Couple bonuse checks for coverage. highest block doesn't increase when it would be backwards
	et.ec.cache.setHighestBlock(big.NewInt(0))
	if et.ec.cache.getHighestBlock().Uint64() == 0 {
		t.Fatal("block should not have decreased")
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
	_, err = et.ec.cache.getNodeInfo(common.HexToAddress("0xff"))
	if err == nil {
		t.Fatal("expected error")
	} else if !errors.Is(err, &NotFoundError{}) {
		t.Fatal("unexpected error", err)
	}

	et.ec.Stop()
	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}
