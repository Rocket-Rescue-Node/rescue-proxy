package consensuslayer

import (
	"context"

	"github.com/allegro/bigcache/v3"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
)

const pubkeyLength = 48
const withdrawalLength = 20

type validatorCache struct {
	*bigcache.BigCache
}

// The cache value will be a byte slice of 68 length
// First 48 bytes for the publick key
// Last 20 bytes for the address of the 0x01 credential, or a guardian address if a BLS key.

func newValidatorCache(ctx context.Context, config bigcache.Config) (*validatorCache, error) {
	bc, err := bigcache.New(ctx, config)
	if err != nil {
		return nil, err
	}

	return &validatorCache{bc}, nil
}

func (c *validatorCache) Get(index string) *ValidatorInfo {
	blob, err := c.BigCache.Get(index)
	if err != nil {
		return nil
	}

	if len(blob) != pubkeyLength+withdrawalLength {
		return nil
	}

	out := ValidatorInfo{}
	out.Pubkey = rptypes.BytesToValidatorPubkey(blob[:pubkeyLength])
	out.WithdrawalAddress = common.BytesToAddress(blob[pubkeyLength:])
	return &out
}

func (c *validatorCache) Set(index string, v *ValidatorInfo) {
	var blob [pubkeyLength + withdrawalLength]byte

	copy(blob[:], v.Pubkey[:])
	copy(blob[pubkeyLength:], v.WithdrawalAddress[:])

	_ = c.BigCache.Set(index, blob[:])
}
