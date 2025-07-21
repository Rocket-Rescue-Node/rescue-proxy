package executionlayer

import "github.com/ethereum/go-ethereum/common"

type stateContracts struct {
	MulticallAddress    common.Address
	StateManagerAddress common.Address
}

var stateContractsMap map[common.Address]stateContracts

const rocketStorageMainnetString = "0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46"
const multicallMainnetString = "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696"
const balancbatcherMainnetString = "0xb1f8e55c7f64d203c1400b9d8555d050f94adf39"

const rocketStorageTestnetString = "0x594Fb75D3dc2DFa0150Ad03F99F97817747dd4E1"
const multicallTestnetString = "0xc5fA61aA6Ec012d1A2Ea38f31ADAf4D06c8725E7"
const balancbatcherTestnetString = "0xB80b500CF68a956b6f149F1C48E8F07EEF4486Ce"

func init() {
	stateContractsMap = map[common.Address]stateContracts{
		common.HexToAddress(rocketStorageMainnetString): {
			MulticallAddress:    common.HexToAddress(multicallMainnetString),
			StateManagerAddress: common.HexToAddress(balancbatcherMainnetString),
		},
		common.HexToAddress(rocketStorageTestnetString): {
			MulticallAddress:    common.HexToAddress(multicallTestnetString),
			StateManagerAddress: common.HexToAddress(balancbatcherTestnetString),
		},
	}
}
