package router

import (
	"crypto/sha256"
	"strings"
	"testing"
	"time"

	"github.com/Rocket-Rescue-Node/credentials"
	"github.com/Rocket-Rescue-Node/credentials/pb"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
)

var nodeId = []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99,
	0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x00}

func setupAuthTest(t *testing.T) *auth {
	_, err := metrics.Init("authentication_test_" + t.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(metrics.Deinit)

	cm := credentials.NewCredentialManager(sha256.New, []byte("test"))
	a := initAuth(cm)
	return a
}

func TestValidCredential(t *testing.T) {
	a := setupAuthTest(t)

	// Create a valid credential
	cred, err := a.cm.Create(time.Now(), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	// Ensure the credential passes authentication
	_, authErr := a.authenticate(username, password)
	if authErr != nil {
		t.Fatal(authErr)
	}
}

func TestExpiredCredential(t *testing.T) {
	a := setupAuthTest(t)

	// Create a valid credential
	cred, err := a.cm.Create(time.Now().Add(-(time.Hour * 24 * 30)), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	_, authErr := a.authenticate(username, password)
	if authErr == nil {
		t.Fatal("expired credential should produce authentication error")
	}
}

func TestEmptyUsername(t *testing.T) {
	a := setupAuthTest(t)

	// Create a valid credential
	cred, err := a.cm.Create(time.Now().Add(-time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := ""
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	_, authErr := a.authenticate(username, password)
	if authErr == nil {
		t.Fatal("missing username should produce authentication error")
	}
}

func TestEmptyPassword(t *testing.T) {
	a := setupAuthTest(t)

	// Create a valid credential
	cred, err := a.cm.Create(time.Now().Add(-time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password := ""

	_, authErr := a.authenticate(username, password)
	if authErr == nil {
		t.Fatal("missing password should produce authentication error")
	}
}

func TestBadBase64(t *testing.T) {
	a := setupAuthTest(t)

	// Create a valid credential
	cred, err := a.cm.Create(time.Now().Add(-time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password := "a space is invalid base64"

	_, authErr := a.authenticate(username, password)
	if authErr == nil {
		t.Fatal("invalid base64 encoding should produce authentication error")
	}
}

func TestBadSecret(t *testing.T) {
	a := setupAuthTest(t)

	// Create a CM with a different secret
	cm := credentials.NewCredentialManager(sha256.New, []byte("wrong"))

	// Create a valid credential, but with the wrong secret
	cred, err := cm.Create(time.Now().Add(-time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	_, authErr := a.authenticate(username, password)
	if authErr == nil {
		t.Fatal("incorrect hmac secret should produce authentication error")
	}
}

func TestFutureCredential(t *testing.T) {
	a := setupAuthTest(t)

	// Create a valid credential
	cred, err := a.cm.Create(time.Now().Add(time.Hour), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	// Ensure the credential passes authentication
	_, authErr := a.authenticate(username, password)
	if authErr != nil {
		t.Fatal(authErr)
	}
}

func TestErrorMessages(t *testing.T) {
	a := setupAuthTest(t)

	// Create an expired credential
	cred, err := a.cm.Create(time.Now().Add(-(time.Hour * 24 * 30)), nodeId, pb.OperatorType_OT_ROCKETPOOL)
	if err != nil {
		t.Fatal(err)
	}

	// Convert to username/password
	username := cred.Base64URLEncodeUsername()
	password, err := cred.Base64URLEncodePassword()
	if err != nil {
		t.Fatal(err)
	}

	_, authErr := a.authenticate(username, password)
	if authErr == nil {
		t.Fatal("expired credential should produce authentication error for string testing")
	}

	if !strings.Contains(authErr.Error(), "authentication failed") {
		t.Fatal("'authentication failed' should be in error message")
	}

	if !strings.Contains(authErr.GRPCError().Error(), "authentication failed") {
		t.Fatal("'authentication failed' should be in grpc error message")
	}
}
