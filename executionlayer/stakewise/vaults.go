package stakewise

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type cacheValue struct {
	mevEscrow   *common.Address
	lastChecked time.Time
}

type VaultsChecker struct {
	registryInstance *bind.BoundContract
	c                sync.Map
	ec               *ethclient.Client
}

var vaultsRegistryAbi = NewVaultsRegistry()
var vaultAbi = NewEthPrivVault()

func NewVaultsChecker(ec *ethclient.Client, vaultsRegistryAddress common.Address) *VaultsChecker {
	registryInstance := vaultsRegistryAbi.Instance(ec, vaultsRegistryAddress)

	return &VaultsChecker{
		registryInstance: registryInstance,
		ec:               ec,
		c:                sync.Map{},
	}
}

// Returns the vault's mevEscrow address if it is a vault, otherwise returns nil
func (v *VaultsChecker) IsVault(ctx context.Context, vaultAddress common.Address) (*common.Address, error) {
	callOpts := bind.CallOpts{
		Context: ctx,
	}
	var cv cacheValue
	value, ok := v.c.Load(vaultAddress)
	if ok {
		cv = value.(cacheValue)
	}
	if !ok || cv.lastChecked.Add(24*time.Hour).Before(time.Now()) {
		resp, err := v.registryInstance.CallRaw(&callOpts, vaultsRegistryAbi.PackVaults(vaultAddress))
		if err != nil {
			return nil, err
		}
		isVault, err := vaultsRegistryAbi.UnpackVaults(resp)
		if err != nil {
			return nil, err
		}
		if !isVault {
			v.c.Store(vaultAddress, cacheValue{mevEscrow: nil, lastChecked: time.Now()})
			return nil, nil
		}

		// If it is a vault, we need to get the mevEscrow address
		vaultContract := vaultAbi.Instance(v.ec, vaultAddress)
		call := vaultAbi.PackMevEscrow()
		resp, err = vaultContract.CallRaw(&callOpts, call)
		if err != nil {
			return nil, err
		}
		mevEscrow, err := vaultAbi.UnpackMevEscrow(resp)
		if err != nil {
			return nil, err
		}
		cv.mevEscrow = &mevEscrow
		cv.lastChecked = time.Now()
		v.c.Store(vaultAddress, cv)
	}

	return cv.mevEscrow, nil
}
