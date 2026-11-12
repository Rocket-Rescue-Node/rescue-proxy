package dataprovider

import (
	"testing"

	rptypes "github.com/rocket-pool/smartnode/bindings/types"
)

func TestBytesToValidatorPubkey(t *testing.T) {
	valid := make([]byte, rptypes.ValidatorPubkeyLength)
	valid[0] = 0xab

	pk, err := bytesToValidatorPubkey(valid)
	if err != nil {
		t.Fatalf("expected valid 48-byte pubkey to succeed: %v", err)
	}
	if pk[0] != 0xab {
		t.Fatalf("unexpected pubkey contents: %x", pk)
	}

	for _, bad := range [][]byte{
		nil,
		{},
		{0x01},
		make([]byte, rptypes.ValidatorPubkeyLength-1),
		make([]byte, rptypes.ValidatorPubkeyLength+1),
	} {
		if _, err := bytesToValidatorPubkey(bad); err == nil {
			t.Fatalf("expected error for pubkey length %d", len(bad))
		}
	}
}
