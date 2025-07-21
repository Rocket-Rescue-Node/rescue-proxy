package consensuslayer

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
)

const pubkeyLength = 48
const withdrawalLength = 20

func TestValidatorInfoBinaryRoundTrip(t *testing.T) {
	hash := md5.Sum([]byte(t.Name()))
	seedInt := binary.LittleEndian.Uint64(hash[len(hash)-8:])
	chaos := rand.NewSource(int64(seedInt))
	r := rand.New(chaos)

	pubkeyBuffer := make([]byte, pubkeyLength)
	addr := common.Address{}
	for range 20 {
		n, err := r.Read(pubkeyBuffer)
		if err != nil {
			t.Fatal(err)
		}
		if n != pubkeyLength {
			t.Fatalf("Expected to read 48 random bytes, read %d", n)
		}

		n, err = r.Read(addr[:])
		if err != nil {
			t.Fatal(err)
		}
		if n != withdrawalLength {
			t.Fatalf("Expected to read 20 random bytes, read %d", n)
		}

		vi := ValidatorInfo{
			Pubkey:            rptypes.BytesToValidatorPubkey(pubkeyBuffer),
			WithdrawalAddress: addr,
			IsELWithdrawal:    r.Int()%2 == 0,
		}
		t.Logf("vi: %+v\n", vi)

		// convert vi to binary and back
		blob, err := vi.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		dvi := new(ValidatorInfo)
		err = dvi.Deserialize(blob)
		if err != nil {
			t.Fatal(err)
		}

		// Compare to ensure dvi == vi
		if !bytes.Equal(dvi.Pubkey[:], vi.Pubkey[:]) ||
			!bytes.Equal(dvi.WithdrawalAddress[:], vi.WithdrawalAddress[:]) ||
			dvi.IsELWithdrawal != vi.IsELWithdrawal {
			t.Fatalf("dvi does not match vi: %+v\n\n vs \n\n %+v", dvi, vi)
		}
	}
}
