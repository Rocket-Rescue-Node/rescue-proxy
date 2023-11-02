package test

import (
	"math/rand"

	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
)

func randPubkey(r *rand.Rand) rptypes.ValidatorPubkey {
	out := make([]byte, rptypes.ValidatorPubkeyLength)
	r.Read(out)
	return rptypes.BytesToValidatorPubkey(out)
}

func randAddress(r *rand.Rand) common.Address {
	out := make([]byte, common.AddressLength)
	r.Read(out)
	return common.BytesToAddress(out)
}

func randValidatorState(r *rand.Rand) apiv1.ValidatorState {
	spread := int64(apiv1.ValidatorStateWithdrawalDone) - int64(apiv1.ValidatorStateUnknown)
	n := r.Int63n(spread + 1)

	return apiv1.ValidatorState(n)
}

func randWithdrawalCredentials(r *rand.Rand) []byte {
	out := make([]byte, 32)

	if r.Int63n(2) == 0 {
		r.Read(out)
		out[0] = 0x00
		return out
	}

	return rand0x01Credentials(r)
}

func rand0x01Credentials(r *rand.Rand) []byte {
	out := make([]byte, 32)

	out[0] = 0x01
	r.Read(out[12:])
	return out
}
