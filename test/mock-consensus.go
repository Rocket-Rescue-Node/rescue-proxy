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

	"github.com/Rocket-Rescue-Node/rescue-proxy/consensuslayer"
)

type MockConsensusLayer struct {
	validators map[string]*apiv1.Validator
	Indices    map[rptypes.ValidatorPubkey]string
}

func NewMockConsensusLayer(numValidators int, seed string) *MockConsensusLayer {
	hash := md5.Sum([]byte(seed))
	// Use the low 8 bytes as the seed for rand
	seedInt := binary.LittleEndian.Uint64(hash[len(hash)-8:])
	chaos := rand.NewSource(int64(seedInt))

	gen := rand.New(chaos)
	out := new(MockConsensusLayer)
	out.validators = make(map[string]*apiv1.Validator, numValidators)
	out.Indices = make(map[rptypes.ValidatorPubkey]string)

	for i := 0; i < numValidators; i++ {
		pubkey := randPubkey(gen)
		idx := 100 + i
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

		out.Indices[pubkey] = fmt.Sprint(idx)
	}

	return out
}

func (m *MockConsensusLayer) AddExecutionValidators(e *MockExecutionLayer, seed string) {
	hash := md5.Sum([]byte(seed))
	// Use the low 8 bytes as the seed for rand
	seedInt := binary.LittleEndian.Uint64(hash[len(hash)-8:])
	chaos := rand.NewSource(int64(seedInt))

	gen := rand.New(chaos)
	i := len(m.validators)
	for pubkey := range e.VMap {
		idx := 100 + i
		balance := phase0.Gwei(gen.Int63())
		m.validators[fmt.Sprint(idx)] = &apiv1.Validator{
			Index:   phase0.ValidatorIndex(idx),
			Balance: balance,
			Status:  randValidatorState(gen),
			Validator: &phase0.Validator{
				PublicKey:             phase0.BLSPubKey(pubkey),
				WithdrawalCredentials: randELCredentials(gen),
				EffectiveBalance:      balance,
				Slashed:               gen.Int63n(10) == 0,
			},
		}
		m.Indices[pubkey] = fmt.Sprint(idx)
		i++
	}

	for vault := range e.SWVaults {
		// add additional validators to the end of the list
		idx := 100 + i
		balance := phase0.Gwei(gen.Int63())
		withdrawalCreds := [32]byte{}
		withdrawalCreds[0] = 0x01
		copy(withdrawalCreds[12:], vault[:])
		pubkey := randPubkey(gen)
		m.validators[fmt.Sprint(idx)] = &apiv1.Validator{
			Index:   phase0.ValidatorIndex(idx),
			Balance: balance,
			Status:  randValidatorState(gen),
			Validator: &phase0.Validator{
				PublicKey:             phase0.BLSPubKey(pubkey),
				WithdrawalCredentials: withdrawalCreds[:],
				EffectiveBalance:      balance,
			},
		}
		m.Indices[pubkey] = fmt.Sprint(idx)
		i++
	}
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

		if bytes.HasPrefix(v.Validator.WithdrawalCredentials, []byte{0x01}) ||
			bytes.HasPrefix(v.Validator.WithdrawalCredentials, []byte{0x02}) {

			out[k].IsELWithdrawal = true
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
