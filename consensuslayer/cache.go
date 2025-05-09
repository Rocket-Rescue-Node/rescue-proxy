package consensuslayer

import (
	"context"

	"github.com/allegro/bigcache/v3"
)

type validatorCache struct {
	*bigcache.BigCache
}

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

	out := &ValidatorInfo{}
	if out.Deserialize(blob) != nil {
		// Deserialization returned an error, allow the user to overwrite this cache entry later
		return nil
	}
	return out
}

func (c *validatorCache) Set(index string, v *ValidatorInfo) error {
	blob, err := v.Serialize()
	if err != nil {
		return err
	}

	return c.BigCache.Set(index, blob)
}
