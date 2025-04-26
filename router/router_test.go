package router

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Rocket-Rescue-Node/credentials"
	"github.com/Rocket-Rescue-Node/credentials/pb"
	gbp "github.com/Rocket-Rescue-Node/guarded-beacon-proxy"
	"github.com/Rocket-Rescue-Node/rescue-proxy/config"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/Rocket-Rescue-Node/rescue-proxy/test"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/metadata"
)

type routerTest struct {
	ctx   context.Context
	pr    *ProxyRouter
	start func()
}

type mockBeaconHandler struct {
	t *testing.T
}

const responseString = "curiouser and curiouser"

func (m *mockBeaconHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.t.Log("handling ", r.URL)
	w.WriteHeader(200)
	_, _ = fmt.Fprintln(w, responseString)
}

func setup(t *testing.T, errs chan error) routerTest {
	_, err := metrics.Init("router_test_" + t.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(metrics.Deinit)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	t.Cleanup(cancel)

	beacon := httptest.NewServer(&mockBeaconHandler{
		t: t,
	})
	t.Cleanup(beacon.Close)

	beaconURL, err := url.Parse(beacon.URL)
	if err != nil {
		t.Fatal(err)
	}

	httpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	cl := test.NewMockConsensusLayer(100, t.Name())
	el := test.NewMockExecutionLayer(50, 5, 100, t.Name())

	cl.AddExecutionValidators(el, t.Name())

	pr := &ProxyRouter{
		Addr:                 httpListener.Addr().String(),
		BeaconURL:            beaconURL,
		CL:                   cl,
		EL:                   el,
		Logger:               zaptest.NewLogger(t),
		CredentialSecrets:    config.CredentialSecrets{[]byte("test"), []byte("test2")},
		EnableSoloValidators: true,
	}
	pr.Init()
	return routerTest{
		ctx: ctx,
		pr:  pr,
		start: func() {
			errs <- pr.Serve(httpListener, nil)
		},
	}
}

func (rt routerTest) validAuth(t *testing.T, solo bool) (string, string) {
	// Grab a node id
	var addr []byte
	err := rt.pr.EL.(*test.MockExecutionLayer).ForEachNode(func(a common.Address) bool {
		addr = a.Bytes()
		return false
	})
	if err != nil {
		t.Fatal(err)
	}

	ot := pb.OperatorType_OT_ROCKETPOOL
	if solo {
		ot = pb.OperatorType_OT_SOLO
	}

	cred, err := rt.pr.auth.credentialManager.Create(time.Now(), addr, ot)
	if err != nil {
		t.Fatal(err)
	}

	pw, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	return cred.Base64URLEncodeUsername(), pw
}

func TestRouterStartStop(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	rt.pr.Stop(rt.ctx)

	err := <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterMissingAuth(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	resp, err := http.Get("http://" + rt.pr.Addr)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 401 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterBadAuth(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	username, pw := rt.validAuth(t, false)
	resp, err := http.Get("http://" + username + ":" + strings.ToLower(pw) + "@" + rt.pr.Addr)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 401 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterGoodAuth(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	username, pw := rt.validAuth(t, false)
	resp, err := http.Get("http://" + username + ":" + pw + "@" + rt.pr.Addr)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(body)) != responseString {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterGoodAuthPartner(t *testing.T) {
	var addr []byte
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()
	err := rt.pr.EL.(*test.MockExecutionLayer).ForEachNode(func(a common.Address) bool {
		addr = a.Bytes()
		return false
	})
	if err != nil {
		t.Fatal(err)
	}

	ot := pb.OperatorType_OT_ROCKETPOOL

	cm := credentials.NewCredentialManager([]byte("test2"))
	cred, err := cm.Create(time.Now(), addr, ot)
	if err != nil {
		t.Fatal(err)
	}

	username := cred.Base64URLEncodeUsername()
	pw, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("http://" + username + ":" + pw + "@" + rt.pr.Addr)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(body)) != responseString {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterGoodAuthSolo(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	username, pw := rt.validAuth(t, true)
	resp, err := http.Get("http://" + username + ":" + pw + "@" + rt.pr.Addr)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(body)) != responseString {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterGoodAuthSoloBackoff(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)
	rt.pr.EnableSoloValidators = false

	go rt.start()

	username, pw := rt.validAuth(t, true)
	resp, err := http.Get("http://" + username + ":" + pw + "@" + rt.pr.Addr)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 429 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterPBPSolo(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	// Grab the list of validators from the mock client
	valis, err := rt.pr.CL.GetValidators()
	if err != nil {
		t.Fatal(err)
	}

	// Find a validator that is 0x01
	var fr common.Address
	var index phase0.ValidatorIndex
	for _, v := range valis {

		// Make sure it's not a RP validator
		info, err := rt.pr.EL.GetRPInfo(rptypes.BytesToValidatorPubkey(v.Validator.PublicKey[:]))
		if err != nil {
			t.Fatal(err)
		}
		if info != nil {
			continue
		}

		withdrawalCreds := v.Validator.WithdrawalCredentials

		if bytes.HasPrefix(withdrawalCreds, []byte{0x01}) {
			fr = common.BytesToAddress(withdrawalCreds)
			index = v.Index
			break
		}
	}

	username, pw := rt.validAuth(t, true)
	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/prepare_beacon_proposer",
		"application/json",
		strings.NewReader(fmt.Sprintf(`
			[{
				"validator_index": "%s",
				"fee_recipient": "%s"
			}]`, fmt.Sprint(index), fr.String()),
		),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(body)) != responseString {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterPBPSoloUnseen(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	username, pw := rt.validAuth(t, true)
	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/prepare_beacon_proposer",
		"application/json",
		strings.NewReader(fmt.Sprintf(`
			[{
				"validator_index": "%s",
				"fee_recipient": "%s"
			}]`, "1010101", "0xabcf8e0d4e9587369b2301d0790347320302cc09"),
		),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 400 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var eMap map[string]string
	err = json.Unmarshal(body, &eMap)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(eMap["error"], "unknown validator index") {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterPBPSoloBadFeeRecipient(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	// Grab the list of validators from the mock client
	valis, err := rt.pr.CL.GetValidators()
	if err != nil {
		t.Fatal(err)
	}

	// Find a validator that is 0x01
	var index phase0.ValidatorIndex
	for _, v := range valis {

		// Make sure it's not a RP validator
		info, err := rt.pr.EL.GetRPInfo(rptypes.BytesToValidatorPubkey(v.Validator.PublicKey[:]))
		if err != nil {
			t.Fatal(err)
		}
		if info != nil {
			continue
		}
		withdrawalCreds := v.Validator.WithdrawalCredentials

		if bytes.HasPrefix(withdrawalCreds, []byte{0x01}) {
			index = v.Index
			break
		}
	}

	// Sneaky check for rp->solo sharing
	username, pw := rt.validAuth(t, false)
	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/prepare_beacon_proposer",
		"application/json",
		strings.NewReader(fmt.Sprintf(`
			[{
				"validator_index": "%s",
				"fee_recipient": "%s"
			}]`, fmt.Sprint(index), "0xabcf8e0d4e9587369b2301d0790347320302cc09"),
		),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 403 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var eMap map[string]string
	err = json.Unmarshal(body, &eMap)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(eMap["error"], "attempting to set fee recipient") {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterPBPRP(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	// Grab a couple validators
	vMap := rt.pr.EL.(*test.MockExecutionLayer).VMap
	rEth := rt.pr.EL.(*test.MockExecutionLayer).REth
	mockIndices := rt.pr.CL.(*test.MockConsensusLayer).Indices

	pubkeys := make([]rptypes.ValidatorPubkey, 0)
	frs := make([]*common.Address, 0)
	indices := make([]string, 0)
	for pubkey, info := range vMap {
		if len(pubkeys) == 2 {
			break
		}
		frs = append(frs, info.ExpectedFeeRecipient)
		pubkeys = append(pubkeys, pubkey)
		indices = append(indices, mockIndices[pubkey])

	}

	username, pw := rt.validAuth(t, false)
	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/prepare_beacon_proposer",
		"application/json",
		strings.NewReader(fmt.Sprintf(`
			[{
				"validator_index": "%s",
				"fee_recipient": "%s"
			},{
				"validator_index": "%s",
				"fee_recipient": "%s"
			}]`,
			indices[0],
			frs[0].String(),
			indices[1],
			rEth.String()),
		),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(body)) != responseString {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterPBPRPCheater(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	// Grab a couple validators
	vMap := rt.pr.EL.(*test.MockExecutionLayer).VMap
	mockIndices := rt.pr.CL.(*test.MockConsensusLayer).Indices

	pubkeys := make([]rptypes.ValidatorPubkey, 0)
	frs := make([]*common.Address, 0)
	indices := make([]string, 0)
	for pubkey, info := range vMap {
		if len(pubkeys) == 2 {
			break
		}
		frs = append(frs, info.ExpectedFeeRecipient)
		pubkeys = append(pubkeys, pubkey)
		indices = append(indices, mockIndices[pubkey])

	}

	// Sneaky check for solo->rp sharing
	username, pw := rt.validAuth(t, true)
	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/prepare_beacon_proposer",
		"application/json",
		strings.NewReader(fmt.Sprintf(`
			[{
				"validator_index": "%s",
				"fee_recipient": "%s"
			},{
				"validator_index": "%s",
				"fee_recipient": "%s"
			}]`,
			indices[0],
			frs[0].String(),
			indices[1],
			"0xabcf8e0d4e9587369b2301d0790347320302cc09"),
		),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 409 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var eMap map[string]string
	err = json.Unmarshal(body, &eMap)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(eMap["error"], "actual fee recipient 0xabcf8e0d4e9587369b2301d0790347320302cc09 didn't match expected fee recipient") {
		t.Fatal("unexpected response", string(body))
	}

	rt.pr.Stop(rt.ctx)

	err = <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterRVSolo(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	// Grab RP validators to exclude
	vMap := rt.pr.EL.(*test.MockExecutionLayer).VMap

	// Grab all validators to pick from
	validators, err := rt.pr.CL.GetValidators()
	if err != nil {
		t.Fatal(err)
	}

	var pubkey rptypes.ValidatorPubkey
	for _, v := range validators {
		key := v.Validator.PublicKey
		// Convert to rptypes
		pubkey = rptypes.BytesToValidatorPubkey(key[:])

		_, ok := vMap[pubkey]
		if !ok {
			// This isn't a rp validator, keep it
			break
		}
	}

	body := fmt.Sprintf(`
			[{
				"message": {
					"gas_limit": "1",
					"timestamp": "1",
					"pubkey": "%s",
					"fee_recipient": "%s"
				},
				"signature": "0x1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505cc411d61252fb6cb3fa0017b679f8bb2305b26a285fa2737f175668d0dff91cc1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505"
			}]`,
		pubkey.String(),
		"0xabcf8e0d4e9587369b2301d0790347320302cc09")
	t.Log("body", body)
	username, pw := rt.validAuth(t, true)
	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/register_validator",
		"application/json",
		strings.NewReader(body),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}
}

func TestRouterRVSoloMalformed(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	body := fmt.Sprintf(`
			[{
				"message": {
					"gas_limit": "1",
					"timestamp": "1",
					"pubkey": "%s",
					"fee_recipient": "%s"
				},
				"signature": "0x1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505cc411d61252fb6cb3fa0017b679f8bb2305b26a285fa2737f175668d0dff91cc1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505"
			}]`,
		"bob",
		"0xabcf8e0d4e9587369b2301d0790347320302cc09")
	t.Log("body", body)

	username, pw := rt.validAuth(t, true)
	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/register_validator",
		"application/json",
		strings.NewReader(body),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 400 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var eMap map[string]string
	err = json.Unmarshal(respBody, &eMap)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(eMap["error"], "error parsing pubkey from request body: Invalid validator public key hex string bob: invalid length 3") {
		t.Fatal("unexpected status", eMap["error"])
	}
}

func TestRouterRVRP(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	// Grab a couple validators
	vMap := rt.pr.EL.(*test.MockExecutionLayer).VMap

	pubkeys := make([]rptypes.ValidatorPubkey, 0)
	frs := make([]*common.Address, 0)
	for pubkey, info := range vMap {
		if len(pubkeys) == 2 {
			break
		}
		frs = append(frs, info.ExpectedFeeRecipient)
		pubkeys = append(pubkeys, pubkey)

	}

	username, pw := rt.validAuth(t, false)

	rEth := rt.pr.EL.(*test.MockExecutionLayer).REth

	body := fmt.Sprintf(`
			[{
				"message": {
					"gas_limit": "1",
					"timestamp": "1",
					"pubkey": "%s",
					"fee_recipient": "%s"
				},
				"signature": "0x1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505cc411d61252fb6cb3fa0017b679f8bb2305b26a285fa2737f175668d0dff91cc1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505"
			},
			{
				"message": {
					"gas_limit": "1",
					"timestamp": "1",
					"pubkey": "%s",
					"fee_recipient": "%s"
				},
				"signature": "0x1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505cc411d61252fb6cb3fa0017b679f8bb2305b26a285fa2737f175668d0dff91cc1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505"
			}
			]`,
		pubkeys[0].String(),
		frs[0].String(),
		pubkeys[1].String(),
		rEth.String(),
	)
	t.Log("body", body)

	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/register_validator",
		"application/json",
		strings.NewReader(body),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}
}

func TestRouterRVCheater(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	// Grab a couple validators
	vMap := rt.pr.EL.(*test.MockExecutionLayer).VMap

	pubkeys := make([]rptypes.ValidatorPubkey, 0)
	frs := make([]*common.Address, 0)
	for pubkey, info := range vMap {
		if len(pubkeys) == 2 {
			break
		}
		frs = append(frs, info.ExpectedFeeRecipient)
		pubkeys = append(pubkeys, pubkey)

	}

	username, pw := rt.validAuth(t, false)

	body := fmt.Sprintf(`
			[{
				"message": {
					"gas_limit": "1",
					"timestamp": "1",
					"pubkey": "%s",
					"fee_recipient": "%s"
				},
				"signature": "0x1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505cc411d61252fb6cb3fa0017b679f8bb2305b26a285fa2737f175668d0dff91cc1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505"
			},
			{
				"message": {
					"gas_limit": "1",
					"timestamp": "1",
					"pubkey": "%s",
					"fee_recipient": "%s"
				},
				"signature": "0x1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505cc411d61252fb6cb3fa0017b679f8bb2305b26a285fa2737f175668d0dff91cc1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505"
			}
			]`,
		pubkeys[0].String(),
		frs[0].String(),
		pubkeys[1].String(),
		"0xabcf8e0d4e9587369b2301d0790347320302cc09",
	)
	t.Log("body", body)

	resp, err := http.Post(
		"http://"+username+":"+pw+"@"+rt.pr.Addr+"/eth/v1/validator/register_validator",
		"application/json",
		strings.NewReader(body),
	)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if resp.StatusCode != 409 {
		t.Fatal("unexpected status code", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var eMap map[string]string
	err = json.Unmarshal(respBody, &eMap)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(eMap["error"], "actual fee recipient 0xabcf8e0d4e9587369b2301d0790347320302cc09 didn't match expected fee recipient") {
		t.Fatal("unexpected status", eMap["error"])
	}
}

func TestRouterGRPCAuth(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	username, pw := rt.validAuth(t, false)
	md := metadata.New(map[string]string{
		"rprnauth": fmt.Sprintf("%s:%s", username, pw),
	})

	authStatus, ctx, err := rt.pr.grpcAuthenticate(md)
	if err != nil {
		t.Fatal(err)
	}

	if authStatus != gbp.Allowed {
		t.Fatal("Unexpected authStatus", authStatus)
	}

	val := ctx.Value(prContextNodeAddrKey)
	node, ok := val.([]byte)
	if !ok {
		t.Fatal("Couldn't get node address from ctx")
	}

	if !strings.EqualFold(username, base64.URLEncoding.EncodeToString(node)) {
		t.Fatal("incorrect node address from ctx")
	}

	val = ctx.Value(prContextOperatorTypeKey)
	ot, ok := val.(pb.OperatorType)
	if !ok {
		t.Fatal("Couldn't get operator type from ctx")
	}
	if ot != pb.OperatorType_OT_ROCKETPOOL {
		t.Fatal("Unexpected operator type from ctx", ot)
	}
}

func TestRouterGRPCAuthSolo(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	username, pw := rt.validAuth(t, true)
	md := metadata.New(map[string]string{
		"rprnauth": fmt.Sprintf("%s:%s", username, pw),
	})

	authStatus, ctx, err := rt.pr.grpcAuthenticate(md)
	if err != nil {
		t.Fatal(err)
	}

	if authStatus != gbp.Allowed {
		t.Fatal("Unexpected authStatus", authStatus)
	}

	val := ctx.Value(prContextNodeAddrKey)
	node, ok := val.([]byte)
	if !ok {
		t.Fatal("Couldn't get node address from ctx")
	}

	if !strings.EqualFold(username, base64.URLEncoding.EncodeToString(node)) {
		t.Fatal("incorrect node address from ctx")
	}

	val = ctx.Value(prContextOperatorTypeKey)
	ot, ok := val.(pb.OperatorType)
	if !ok {
		t.Fatal("Couldn't get operator type from ctx")
	}
	if ot != pb.OperatorType_OT_SOLO {
		t.Fatal("Unexpected operator type from ctx", ot)
	}
}

func TestRouterGRPCAuthMissing(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	md := metadata.New(map[string]string{})

	authStatus, _, err := rt.pr.grpcAuthenticate(md)
	if err == nil || err.Error() != "headers missing" {
		t.Fatal("expected error about missing headers", err)
	}

	if authStatus != gbp.Unauthorized {
		t.Fatal("Unexpected authStatus", authStatus)
	}

}

func TestRouterGRPCAuthMissingColon(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	md := metadata.New(map[string]string{
		"rprnauth": "we're all mad here",
	})

	authStatus, _, err := rt.pr.grpcAuthenticate(md)
	if err == nil || err.Error() != "headers invalid" {
		t.Fatal("expected error about invalid headers", err)
	}

	if authStatus != gbp.Unauthorized {
		t.Fatal("Unexpected authStatus", authStatus)
	}

}

func TestRouterGRPCAuthMalformed(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)

	go rt.start()

	username, pw := rt.validAuth(t, false)
	md := metadata.New(map[string]string{
		"rprnauth": fmt.Sprintf("%s:%s", username, strings.ToLower(pw)),
	})

	authStatus, _, err := rt.pr.grpcAuthenticate(md)
	if err == nil || !strings.HasPrefix(err.Error(), "authentication failed, malformed credentials") {
		t.Fatal("expected error about malformed headers", err)
	}

	if authStatus != gbp.Unauthorized {
		t.Fatal("Unexpected authStatus", authStatus)
	}

}

func TestRouterGRPCAuthSoloBackoff(t *testing.T) {
	errs := make(chan error)
	rt := setup(t, errs)
	rt.pr.EnableSoloValidators = false

	go rt.start()

	username, pw := rt.validAuth(t, true)
	md := metadata.New(map[string]string{
		"rprnauth": fmt.Sprintf("%s:%s", username, pw),
	})

	authStatus, _, err := rt.pr.grpcAuthenticate(md)
	if err == nil || !strings.HasPrefix(err.Error(), "solo validator support was manually disabled") {
		t.Fatal("expected error about backoff", err)
	}

	if authStatus != gbp.TooManyRequests {
		t.Fatal("Unexpected authStatus", authStatus)
	}
}
