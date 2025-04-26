package consensuslayer

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
)

type ValidatorInfo struct {
	Pubkey            rptypes.ValidatorPubkey
	WithdrawalAddress common.Address
	IsELWithdrawal    bool
}

func (v *ValidatorInfo) Serialize() ([]byte, error) {
	outSize := binary.Size(v)
	if outSize == -1 {
		panic("ValidatorInfo not serializable by encoding/binary")
	}
	out := make([]byte, outSize)
	written, err := binary.Encode(out, binary.NativeEndian, v)
	if err != nil {
		return nil, fmt.Errorf("couldn't serialize ValidatorInfo: %w", err)
	}
	return out[:written], nil
}

func (v *ValidatorInfo) Deserialize(buf []byte) error {
	_, err := binary.Decode(buf, binary.NativeEndian, v)
	if err != nil {
		return fmt.Errorf("couldn't deserialize ValidatorInfo: %w", err)
	}
	return nil
}
