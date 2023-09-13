package router

import (
	"crypto/sha256"
	"testing"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/credentials"
	"github.com/Rocket-Pool-Rescue-Node/credentials/pb"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
)

var nodeId = []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99,
	0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x00}

func setup(t *testing.T) func() {
	_, err := metrics.Init("authentication_test_" + t.Name())
	if err != nil {
		t.Error(err)
	}

	cm := credentials.NewCredentialManager(sha256.New, []byte("test"))
	InitAuth(cm, time.Minute*5)
	return func() {
		metrics.Deinit()
		DeinitAuth()
	}
}

func TestValidCredential(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Create a valid credential
	cred, err := cm.Create(time.Now(), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Error(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Error(err)
	}

	// Ensure the credential passes authentication
	_, authErr := authenticate(username, password)
	if authErr != nil {
		t.Error(authErr)
	}
}

func TestExpiredCredential(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Create a valid credential
	cred, err := cm.Create(time.Now().Add(-time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Error(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Error(err)
	}

	_, authErr := authenticate(username, password)
	if authErr == nil {
		t.Fail()
	}
}

func TestEmptyUsername(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Create a valid credential
	cred, err := cm.Create(time.Now().Add(-time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Error(err)
	}

	// Convert to username/password
	username := ""
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Error(err)
	}

	_, authErr := authenticate(username, password)
	if authErr == nil {
		t.Fail()
	}
}

func TestEmptyPassword(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Create a valid credential
	cred, err := cm.Create(time.Now().Add(-time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Error(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password := ""

	_, authErr := authenticate(username, password)
	if authErr == nil {
		t.Fail()
	}
}

func TestFutureCredential(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// Create a valid credential
	cred, err := cm.Create(time.Now().Add(time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Error(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Error(err)
	}

	// Ensure the credential passes authentication
	_, authErr := authenticate(username, password)
	if authErr != nil {
		t.Error(authErr)
	}
}
