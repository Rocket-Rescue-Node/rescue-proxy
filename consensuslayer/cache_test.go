package consensuslayer

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
)

func TestCacheRoundTrip(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := bigcache.DefaultConfig(10 * time.Hour)

	cache, err := newValidatorCache(ctx, config)
	if err != nil {
		t.Error(err)
	}

	expectedKey, _ := hex.DecodeString("b20fb4a9340f8b23197b8449db7a5d3d8d068570a2b61a8d78817537aac4fd5645434d3e89a918a3ba9d0b7707cbeae0")
	expectedAddr, _ := hex.DecodeString("6a6d731664115Ff3C823807442a4dC94999b0923")

	cache.Set("test", &ValidatorInfo{
		Pubkey:            rptypes.BytesToValidatorPubkey(expectedKey),
		WithdrawalAddress: common.BytesToAddress(expectedAddr),
	})

	vInfo := cache.Get("test")
	if vInfo == nil {
		t.Fail()
		return
	}

	if !bytes.Equal(vInfo.Pubkey[:], expectedKey) {
		t.Fail()
	}

	if !bytes.Equal(vInfo.WithdrawalAddress[:], expectedAddr) {
		t.Fail()
	}
}
