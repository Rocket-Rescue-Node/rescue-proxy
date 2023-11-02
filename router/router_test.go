package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/credentials/pb"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/test"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap/zaptest"
)

type routerTest struct {
	ctx context.Context
	pr  *ProxyRouter
}

type mockBeaconHandler struct {
	t *testing.T
}

const responseString = "curiouser and curiouser"

func (m *mockBeaconHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.t.Log("handling ", r.URL)
	w.WriteHeader(200)
	fmt.Fprintln(w, responseString)
}

func setup(t *testing.T) routerTest {
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

	mockServer := httptest.NewUnstartedServer(new(mockBeaconHandler))
	listenAddr := mockServer.Listener.Addr().String()
	// Close the listener so its port can be reused by the router
	mockServer.Close()

	cl := test.NewMockConsensusLayer(100, t.Name())
	el := test.NewMockExecutionLayer(50, 5, 100, t.Name())

	cl.AddExecutionValidators(el, t.Name())

	return routerTest{
		ctx: ctx,
		pr: &ProxyRouter{
			Addr:                 listenAddr,
			BeaconURL:            beaconURL,
			CL:                   cl,
			EL:                   el,
			Logger:               zaptest.NewLogger(t),
			CredentialSecret:     "test",
			AuthValidityWindow:   time.Hour,
			EnableSoloValidators: true,
			//GRPCAddr             string
			//GRPCBeaconURL        string
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

	cred, err := rt.pr.auth.cm.Create(time.Now(), addr, ot)
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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	time.Sleep(50 * time.Millisecond)
	rt.pr.Stop(rt.ctx)

	err := <-errs
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouterMissingAuth(t *testing.T) {
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)
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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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

func TestRouterGoodAuthSolo(t *testing.T) {
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
	rt := setup(t)
	rt.pr.EnableSoloValidators = false

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
	rt := setup(t)

	errs := make(chan error)
	go func() {
		err := rt.pr.Start()
		errs <- err
	}()

	// Give the server a second to wake up
	time.Sleep(50 * time.Millisecond)

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
