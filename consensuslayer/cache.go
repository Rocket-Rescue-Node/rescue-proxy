package consensuslayer

import (
	"context"

	"github.com/allegro/bigcache/v3"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
)

const pubkeyLength = 48
const withdrawalLength = 20

// Leave a byte for Is0x01
const blobLength = pubkeyLength + withdrawalLength + 1

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

	if len(blob) != blobLength {
		return nil
	}

	out := ValidatorInfo{}
	out.Pubkey = rptypes.BytesToValidatorPubkey(blob[:pubkeyLength])
	out.WithdrawalAddress = common.BytesToAddress(blob[pubkeyLength : pubkeyLength+withdrawalLength])
	if blob[pubkeyLength+withdrawalLength] == 0x01 {
		out.Is0x01 = true
	}
	return &out
}

func (c *validatorCache) Set(index string, v *ValidatorInfo) error {
	var blob [blobLength]byte

	copy(blob[:], v.Pubkey[:])
	copy(blob[pubkeyLength:], v.WithdrawalAddress[:])
	if v.Is0x01 {
		copy(blob[pubkeyLength+withdrawalLength:], []byte{0x01})
	}

	return c.BigCache.Set(index, blob[:])
}
