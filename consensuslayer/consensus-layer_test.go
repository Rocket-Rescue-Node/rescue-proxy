package consensuslayer

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"go.uber.org/zap/zaptest"
)

const mainnetGenesis = `{"data":{"genesis_time":"1606824023","genesis_validators_root":"0x4b363db94e286120d76eb905340fdd4e54bfe9f06bf33ff6cf5ad27f511bfe95","genesis_fork_version":"0x00000000"}}`
const mainnetConfigSpec = `{"data":{"MAX_COMMITTEES_PER_SLOT":"64","TARGET_COMMITTEE_SIZE":"128","MAX_VALIDATORS_PER_COMMITTEE":"2048","SHUFFLE_ROUND_COUNT":"90","HYSTERESIS_QUOTIENT":"4","HYSTERESIS_DOWNWARD_MULTIPLIER":"1","HYSTERESIS_UPWARD_MULTIPLIER":"5","MIN_DEPOSIT_AMOUNT":"1000000000","MAX_EFFECTIVE_BALANCE":"32000000000","EFFECTIVE_BALANCE_INCREMENT":"1000000000","MIN_ATTESTATION_INCLUSION_DELAY":"1","SLOTS_PER_EPOCH":"32","MIN_SEED_LOOKAHEAD":"1","MAX_SEED_LOOKAHEAD":"4","EPOCHS_PER_ETH1_VOTING_PERIOD":"64","SLOTS_PER_HISTORICAL_ROOT":"8192","MIN_EPOCHS_TO_INACTIVITY_PENALTY":"4","EPOCHS_PER_HISTORICAL_VECTOR":"65536","EPOCHS_PER_SLASHINGS_VECTOR":"8192","HISTORICAL_ROOTS_LIMIT":"16777216","VALIDATOR_REGISTRY_LIMIT":"1099511627776","BASE_REWARD_FACTOR":"64","WHISTLEBLOWER_REWARD_QUOTIENT":"512","PROPOSER_REWARD_QUOTIENT":"8","INACTIVITY_PENALTY_QUOTIENT":"67108864","MIN_SLASHING_PENALTY_QUOTIENT":"128","PROPORTIONAL_SLASHING_MULTIPLIER":"1","MAX_PROPOSER_SLASHINGS":"16","MAX_ATTESTER_SLASHINGS":"2","MAX_ATTESTATIONS":"128","MAX_DEPOSITS":"16","MAX_VOLUNTARY_EXITS":"16","INACTIVITY_PENALTY_QUOTIENT_ALTAIR":"50331648","MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR":"64","PROPORTIONAL_SLASHING_MULTIPLIER_ALTAIR":"2","SYNC_COMMITTEE_SIZE":"512","EPOCHS_PER_SYNC_COMMITTEE_PERIOD":"256","MIN_SYNC_COMMITTEE_PARTICIPANTS":"1","UPDATE_TIMEOUT":"8192","INACTIVITY_PENALTY_QUOTIENT_BELLATRIX":"16777216","MIN_SLASHING_PENALTY_QUOTIENT_BELLATRIX":"32","PROPORTIONAL_SLASHING_MULTIPLIER_BELLATRIX":"3","MAX_BYTES_PER_TRANSACTION":"1073741824","MAX_TRANSACTIONS_PER_PAYLOAD":"1048576","BYTES_PER_LOGS_BLOOM":"256","MAX_EXTRA_DATA_BYTES":"32","MAX_BLS_TO_EXECUTION_CHANGES":"16","MAX_WITHDRAWALS_PER_PAYLOAD":"16","MAX_VALIDATORS_PER_WITHDRAWALS_SWEEP":"16384","PRESET_BASE":"mainnet","CONFIG_NAME":"mainnet","TERMINAL_TOTAL_DIFFICULTY":"58750000000000000000000","TERMINAL_BLOCK_HASH":"0x0000000000000000000000000000000000000000000000000000000000000000","TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH":"18446744073709551615","MIN_GENESIS_ACTIVE_VALIDATOR_COUNT":"16384","MIN_GENESIS_TIME":"1606824000","GENESIS_FORK_VERSION":"0x00000000","GENESIS_DELAY":"604800","ALTAIR_FORK_VERSION":"0x01000000","ALTAIR_FORK_EPOCH":"74240","BELLATRIX_FORK_VERSION":"0x02000000","BELLATRIX_FORK_EPOCH":"144896","CAPELLA_FORK_VERSION":"0x03000000","CAPELLA_FORK_EPOCH":"194048","DENEB_FORK_VERSION":"0x04000000","DENEB_FORK_EPOCH":"18446744073709551615","SECONDS_PER_SLOT":"12","SECONDS_PER_ETH1_BLOCK":"14","MIN_VALIDATOR_WITHDRAWABILITY_DELAY":"256","SHARD_COMMITTEE_PERIOD":"256","ETH1_FOLLOW_DISTANCE":"2048","INACTIVITY_SCORE_BIAS":"4","INACTIVITY_SCORE_RECOVERY_RATE":"16","EJECTION_BALANCE":"16000000000","MIN_PER_EPOCH_CHURN_LIMIT":"4","CHURN_LIMIT_QUOTIENT":"65536","PROPOSER_SCORE_BOOST":"40","DEPOSIT_CHAIN_ID":"1","DEPOSIT_NETWORK_ID":"1","DEPOSIT_CONTRACT_ADDRESS":"0x00000000219ab540356cbb839cbe05303d7705fa","BLS_WITHDRAWAL_PREFIX":"0x00","ETH1_ADDRESS_WITHDRAWAL_PREFIX":"0x01","DOMAIN_BEACON_PROPOSER":"0x00000000","DOMAIN_BEACON_ATTESTER":"0x01000000","DOMAIN_RANDAO":"0x02000000","DOMAIN_DEPOSIT":"0x03000000","DOMAIN_VOLUNTARY_EXIT":"0x04000000","DOMAIN_SELECTION_PROOF":"0x05000000","DOMAIN_AGGREGATE_AND_PROOF":"0x06000000","TIMELY_SOURCE_FLAG_INDEX":"0x00","TIMELY_TARGET_FLAG_INDEX":"0x01","TIMELY_HEAD_FLAG_INDEX":"0x02","TIMELY_SOURCE_WEIGHT":"14","TIMELY_TARGET_WEIGHT":"26","TIMELY_HEAD_WEIGHT":"14","SYNC_REWARD_WEIGHT":"2","PROPOSER_WEIGHT":"8","WEIGHT_DENOMINATOR":"64","DOMAIN_SYNC_COMMITTEE":"0x07000000","DOMAIN_SYNC_COMMITTEE_SELECTION_PROOF":"0x08000000","DOMAIN_CONTRIBUTION_AND_PROOF":"0x09000000","DOMAIN_BLS_TO_EXECUTION_CHANGE":"0x0a000000","TARGET_AGGREGATORS_PER_COMMITTEE":"16","RANDOM_SUBNETS_PER_VALIDATOR":"1","EPOCHS_PER_RANDOM_SUBNET_SUBSCRIPTION":"256","ATTESTATION_SUBNET_COUNT":"64","TARGET_AGGREGATORS_PER_SYNC_SUBCOMMITTEE":"16","SYNC_COMMITTEE_SUBNET_COUNT":"4"}}`
const mainnetConfigDepositContract = `{"data":{"chain_id":"1","address":"0x00000000219ab540356cbb839cbe05303d7705fa"}}`
const mainnetConfigForkSchedule = `{"data":[{"previous_version":"0x00000000","current_version":"0x00000000","epoch":"0"},{"previous_version":"0x00000000","current_version":"0x01000000","epoch":"74240"},{"previous_version":"0x01000000","current_version":"0x02000000","epoch":"144896"},{"previous_version":"0x02000000","current_version":"0x03000000","epoch":"194048"},{"previous_version":"0x03000000","current_version":"0x04000000","epoch":"18446744073709551615"}]}`

type ccTest struct {
	ccl *CachingConsensusLayer
	ctx context.Context
}

type mockHandler struct {
	t *testing.T
	h func(http.ResponseWriter, *http.Request)
}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()

	m.t.Log("handling " + path)
	// Let's return json
	w.Header()["Content-Type"] = []string{"application/json"}
	w.WriteHeader(200)
	// Mainnet responses
	switch path {
	case "/eth/v1/beacon/genesis":
		fmt.Fprintln(w, mainnetGenesis)
		return
	case "/eth/v1/config/spec":
		fmt.Fprintln(w, mainnetConfigSpec)
		return
	case "/eth/v1/config/deposit_contract":
		fmt.Fprintln(w, mainnetConfigDepositContract)
		return
	case "/eth/v1/config/fork_schedule":
		fmt.Fprintln(w, mainnetConfigForkSchedule)
		return
	case "/eth/v1/node/version":
		fmt.Fprintln(w, `{"data":{"version":"None-of-your-beeswax"}}`)
		return
	}

	// If we haven't handled it, fall through to the test-specific handler
	m.h(w, r)
}

func setup(t *testing.T, url *url.URL) ccTest {
	_, err := metrics.Init("cc_test_" + t.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(metrics.Deinit)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	t.Cleanup(cancel)

	return ccTest{
		ccl: NewCachingConsensusLayer(url, zaptest.NewLogger(t), false),
		ctx: ctx,
	}
}

func TestCLStartStop(t *testing.T) {
	s := httptest.NewServer(&mockHandler{
		t: t,
		h: func(w http.ResponseWriter, r *http.Request) {
		},
	})
	t.Cleanup(s.Close)

	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	cct := setup(t, u)
	err = cct.ccl.Init(cct.ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(cct.ccl.Deinit)
}

func TestGetValidatorInfo(t *testing.T) {
	s := httptest.NewServer(&mockHandler{
		t: t,
		h: func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.String() {
			case "/eth/v1/beacon/states/head/validators?id=100,101":
				fmt.Fprintf(w, `{"execution_optimistic":false,"data":[{"index":"100","balance":"32005252956","status":"active_ongoing","validator":{"pubkey":"0xb5bc96b70df0dfcc252c9ff0d1b42cb6dc0d55f8defa474dc0a5c7e0402c241e2850fea9c582e276b638b3c2c3a5ec55","withdrawal_credentials":"0x010000000000000000000000801e880e2e9aa87b20c9cc9ebf7375adb11eac21","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}},{"index":"101","balance":"32005224938","status":"active_ongoing","validator":{"pubkey":"0xaa160542c2b1b9dbf5e11ca044067526c6dfff65efba88ea483d49bdbe478ab7489f8b1a903ea22b6d30cfa57626ca9e","withdrawal_credentials":"0x010000000000000000000000d944cf00517e2dd8d00bd5c4ea1ec45cf3ec52db","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}}]}`)
				return
			}
		},
	})
	t.Cleanup(s.Close)

	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	cct := setup(t, u)
	err = cct.ccl.Init(cct.ctx)
	if err != nil {
		t.Fatal(err)
	}

	validatorInfo, err := cct.ccl.GetValidatorInfo([]string{"100", "101"})
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range validatorInfo {
		if !v.Is0x01 {
			t.Fatal("Validator was not identified as 0x01")
		}

		if !strings.EqualFold(v.WithdrawalAddress.String(), "0x801e880e2e9aa87b20c9cc9ebf7375adb11eac21") &&
			!strings.EqualFold(v.WithdrawalAddress.String(), "0xd944cf00517e2dd8d00bd5c4ea1ec45cf3ec52db") {

			t.Fatal("Unexpected withdrawal address", v.WithdrawalAddress.String())
		}

		if !strings.EqualFold(v.Pubkey.String(), "b5bc96b70df0dfcc252c9ff0d1b42cb6dc0d55f8defa474dc0a5c7e0402c241e2850fea9c582e276b638b3c2c3a5ec55") &&
			!strings.EqualFold(v.Pubkey.String(), "aa160542c2b1b9dbf5e11ca044067526c6dfff65efba88ea483d49bdbe478ab7489f8b1a903ea22b6d30cfa57626ca9e") {
			t.Fatal("Unexpected pubkey", v.Pubkey.String())
		}
	}

	if len(validatorInfo) != 2 {
		t.Fatal("expected 2 validators in the response")
	}

	t.Cleanup(cct.ccl.Deinit)
}

func TestGetValidatorCached(t *testing.T) {
	once := false
	s := httptest.NewServer(&mockHandler{
		t: t,
		h: func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.String() {
			case "/eth/v1/beacon/states/head/validators?id=100,101":
				if once {
					// Landmine to trigger an error if the cache is missed
					fmt.Fprintf(w, "Not json!")
					return
				}
				once = true
				fmt.Fprintf(w, `{"execution_optimistic":false,"data":[{"index":"100","balance":"32005252956","status":"active_ongoing","validator":{"pubkey":"0xb5bc96b70df0dfcc252c9ff0d1b42cb6dc0d55f8defa474dc0a5c7e0402c241e2850fea9c582e276b638b3c2c3a5ec55","withdrawal_credentials":"0x010000000000000000000000801e880e2e9aa87b20c9cc9ebf7375adb11eac21","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}},{"index":"101","balance":"32005224938","status":"active_ongoing","validator":{"pubkey":"0xaa160542c2b1b9dbf5e11ca044067526c6dfff65efba88ea483d49bdbe478ab7489f8b1a903ea22b6d30cfa57626ca9e","withdrawal_credentials":"0x010000000000000000000000d944cf00517e2dd8d00bd5c4ea1ec45cf3ec52db","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}}]}`)
				return
			}
		},
	})
	t.Cleanup(s.Close)

	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	cct := setup(t, u)
	err = cct.ccl.Init(cct.ctx)
	if err != nil {
		t.Fatal(err)
	}

	for range []int{1, 2} {
		validatorInfo, err := cct.ccl.GetValidatorInfo([]string{"100", "101"})
		if err != nil {
			t.Fatal(err)
		}

		if len(validatorInfo) != 2 {
			t.Fatal("expected 2 validators in the response")
		}

		for _, v := range validatorInfo {
			if !v.Is0x01 {
				t.Fatal("Validator was not identified as 0x01")
			}

			if !strings.EqualFold(v.WithdrawalAddress.String(), "0x801e880e2e9aa87b20c9cc9ebf7375adb11eac21") &&
				!strings.EqualFold(v.WithdrawalAddress.String(), "0xd944cf00517e2dd8d00bd5c4ea1ec45cf3ec52db") {

				t.Fatal("Unexpected withdrawal address", v.WithdrawalAddress.String())
			}

			if !strings.EqualFold(v.Pubkey.String(), "b5bc96b70df0dfcc252c9ff0d1b42cb6dc0d55f8defa474dc0a5c7e0402c241e2850fea9c582e276b638b3c2c3a5ec55") &&
				!strings.EqualFold(v.Pubkey.String(), "aa160542c2b1b9dbf5e11ca044067526c6dfff65efba88ea483d49bdbe478ab7489f8b1a903ea22b6d30cfa57626ca9e") {
				t.Fatal("Unexpected pubkey", v.Pubkey.String())
			}
		}
	}

	t.Cleanup(cct.ccl.Deinit)
}

const mockBeaconState = `{"version":"phase0","execution_optimistic":false,"finalized":false,"data":{"genesis_time":"1","genesis_validators_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","slot":"1","fork":{"previous_version":"0x00000000","current_version":"0x00000000","epoch":"1"},"latest_block_header":{"slot":"1","proposer_index":"1","parent_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","state_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","body_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"},"block_roots":["0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"],"state_roots":["0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"],"historical_roots":["0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"],"eth1_data":{"deposit_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","deposit_count":"1","block_hash":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"},"eth1_data_votes":[{"deposit_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","deposit_count":"1","block_hash":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"}],"eth1_deposit_index":"1","validators":[{"pubkey":"0x93247f2209abcacf57b75a51dafae777f9dd38bc7053d1af526f220a7489a6d3a2753e5f3e8b1cfe39b56f43611df74a","withdrawal_credentials":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","effective_balance":"1","slashed":false,"activation_eligibility_epoch":"1","activation_epoch":"1","exit_epoch":"1","withdrawable_epoch":"1"}],"balances":["1"],"randao_mixes":["0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"],"slashings":["1"],"previous_epoch_attestations":[{"aggregation_bits":"0x01","data":{"slot":"1","index":"1","beacon_block_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","source":{"epoch":"1","root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"},"target":{"epoch":"1","root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"}},"inclusion_delay":"1","proposer_index":"1"}],"current_epoch_attestations":[{"aggregation_bits":"0x01","data":{"slot":"1","index":"1","beacon_block_root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2","source":{"epoch":"1","root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"},"target":{"epoch":"1","root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"}},"inclusion_delay":"1","proposer_index":"1"}],"justification_bits":"0x01","previous_justified_checkpoint":{"epoch":"1","root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"},"current_justified_checkpoint":{"epoch":"1","root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"},"finalized_checkpoint":{"epoch":"1","root":"0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"}}}`

func TestGetValidators(t *testing.T) {
	s := httptest.NewServer(&mockHandler{
		t: t,
		h: func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.String() {
			case "/eth/v2/debug/beacon/states/finalized":
				fmt.Fprintln(w, mockBeaconState)
				return
			}
		},
	})
	t.Cleanup(s.Close)

	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	cct := setup(t, u)
	err = cct.ccl.Init(cct.ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Grab a random validator
	_, err = cct.ccl.GetValidators()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(cct.ccl.Deinit)
}

func TestOnHeadUpdate(t *testing.T) {
	s := httptest.NewServer(&mockHandler{
		t: t,
		h: func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.String() {
			case "/eth/v1/beacon/states/head/validators?id=100,101":
				fmt.Fprintf(w, `{"execution_optimistic":false,"data":[{"index":"100","balance":"32005252956","status":"active_ongoing","validator":{"pubkey":"0xb5bc96b70df0dfcc252c9ff0d1b42cb6dc0d55f8defa474dc0a5c7e0402c241e2850fea9c582e276b638b3c2c3a5ec55","withdrawal_credentials":"0x010000000000000000000000801e880e2e9aa87b20c9cc9ebf7375adb11eac21","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}},{"index":"101","balance":"32005224938","status":"active_ongoing","validator":{"pubkey":"0xaa160542c2b1b9dbf5e11ca044067526c6dfff65efba88ea483d49bdbe478ab7489f8b1a903ea22b6d30cfa57626ca9e","withdrawal_credentials":"0x010000000000000000000000d944cf00517e2dd8d00bd5c4ea1ec45cf3ec52db","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}}]}`)
				return
			}
		},
	})
	t.Cleanup(s.Close)

	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	cct := setup(t, u)
	err = cct.ccl.Init(cct.ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range []int{1, 33, 66, 99} {
		r, err := cct.ccl.GetValidatorInfo([]string{"100", "101"})
		if err != nil {
			t.Fatal(err)
		}

		for _, info := range r {
			metrics.ObserveValidator(info.WithdrawalAddress, info.Pubkey)
		}

		// Advance the internal head counter to update metrics
		cct.ccl.onHeadUpdate(&apiv1.Event{
			Data: &apiv1.HeadEvent{
				Slot: phase0.Slot(i),
			},
		})
	}

	pev := metrics.PreviousEpochValidators()
	if pev != 2 {
		t.Fatal("Expected 2 validators, got ", pev)
	}

	t.Cleanup(cct.ccl.Deinit)
}

func TestErrorPaths(t *testing.T) {
	s := httptest.NewServer(&mockHandler{
		t: t,
		h: func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.String() {
			case "/eth/v1/beacon/states/head/validators?id=100,101":
				fmt.Fprintf(w, `{"execution_optimistic":false,"data":[{"index":"100","balance":"32005252956","status":"active_ongoing","validator":{"pubkey":"0xb5bc96b70df0dfcc252c9ff0d1b42cb6dc0d55f8defa474dc0a5c7e0402c241e2850fea9c582e276b638b3c2c3a5ec55","withdrawal_credentials":"0x000000000000000000000000801e880e2e9aa87b20c9cc9ebf7375adb11eac21","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}},{"index":"101","balance":"32005224938","status":"active_ongoing","validator":{"pubkey":"0xaa160542c2b1b9dbf5e11ca044067526c6dfff65efba88ea483d49bdbe478ab7489f8b1a903ea22b6d30cfa57626ca9e","withdrawal_credentials":"0x010000000000000000000000d944cf00517e2dd8d00bd5c4ea1ec45cf3ec52db","effective_balance":"32000000000","slashed":false,"activation_eligibility_epoch":"0","activation_epoch":"0","exit_epoch":"18446744073709551615","withdrawable_epoch":"18446744073709551615"}}]}`)
				return
			}
		},
	})

	t.Cleanup(s.Close)

	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	cct := setup(t, u)
	err = cct.ccl.Init(cct.ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Make sure non-conforming events don't crash
	cct.ccl.onHeadUpdate(&apiv1.Event{
		Data: struct{}{},
	})

	// 0x00 validators trigger logging
	_, err = cct.ccl.GetValidatorInfo([]string{"100", "101"})
	if err != nil {
		t.Fatal(err)
	}

	// Invalid bn response is appropriately bubbled up
	_, err = cct.ccl.GetValidatorInfo([]string{"100", "bob"})
	if err == nil {
		t.Fatal("Expected error")
	}
}
