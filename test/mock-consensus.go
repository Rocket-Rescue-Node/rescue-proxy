package test

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/rand"

	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
)

type MockConsensusLayer struct {
	validators map[string]*apiv1.Validator
}

func NewMockConsensusLayer(numValidators int, seed string) *MockConsensusLayer {
	hash := md5.Sum([]byte(seed))
	// Use the low 8 bytes as the seed for rand
	seedInt := binary.LittleEndian.Uint64(hash[len(hash)-8:])
	chaos := rand.NewSource(int64(seedInt))

	gen := rand.New(chaos)
	out := new(MockConsensusLayer)
	out.validators = make(map[string]*apiv1.Validator, numValidators)

	for i := 0; i < numValidators; i++ {
		pubkey := randPubkey(gen)
		idx := gen.Int63n(int64(numValidators) * 2)
		balance := phase0.Gwei(gen.Int63())
		out.validators[fmt.Sprint(idx)] = &apiv1.Validator{
			Index:   phase0.ValidatorIndex(idx),
			Balance: balance,
			Status:  randValidatorState(gen),
			Validator: &phase0.Validator{
				PublicKey:             phase0.BLSPubKey(pubkey),
				WithdrawalCredentials: randWithdrawalCredentials(gen),
				EffectiveBalance:      balance,
				Slashed:               gen.Int63n(10) == 0,
			},
		}
	}

	return out
}

func (m *MockConsensusLayer) GetValidatorInfo(idx []string) (map[string]*consensuslayer.ValidatorInfo, error) {
	out := make(map[string]*consensuslayer.ValidatorInfo)
	for _, k := range idx {
		v, ok := m.validators[k]
		if !ok {
			continue
		}

		out[k] = &consensuslayer.ValidatorInfo{
			Pubkey: rptypes.BytesToValidatorPubkey(v.Validator.PublicKey[:]),
		}

		if bytes.HasPrefix(v.Validator.WithdrawalCredentials, []byte{0x01}) {
			out[k].Is0x01 = true
			out[k].WithdrawalAddress = common.BytesToAddress(v.Validator.WithdrawalCredentials)
		}
	}

	return out, nil
}

func (m *MockConsensusLayer) GetValidators() ([]*apiv1.Validator, error) {
	out := make([]*apiv1.Validator, 0, len(m.validators))

	for _, v := range m.validators {
		out = append(out, v)
	}

	return out, nil
}
