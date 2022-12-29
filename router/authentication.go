package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/credentials"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authenticationError struct {
	msg        string
	httpStatus int
	grpcCode   codes.Code
}

func (a *authenticationError) Error() string {
	return "authentication failed, " + a.msg
}

func (a *authenticationError) GRPCError() error {
	return status.Error(a.grpcCode, a.Error())
}

var (
	metricsRegistry    *metrics.MetricsRegistry
	authValidityWindow time.Duration
	cm                 *credentials.CredentialManager
)

func malformed(err error) *authenticationError {
	return &authenticationError{
		msg:        "malformed credentials: " + err.Error(),
		httpStatus: http.StatusUnauthorized,
		grpcCode:   codes.Unauthenticated,
	}
}

func invalid(err error) *authenticationError {
	return &authenticationError{
		msg:        "invalid credentials: " + err.Error(),
		httpStatus: http.StatusUnauthorized,
		grpcCode:   codes.Unauthenticated,
	}
}

func expired() *authenticationError {
	return &authenticationError{
		msg:        "expired credentials",
		httpStatus: http.StatusUnauthorized,
		grpcCode:   codes.PermissionDenied,
	}
}

// authenticate returns nil if the username/password are valid and current
// username/password must be base64url encoded
// otherwise, it returns an authentication error
func authenticate(username, password string) (*credentials.AuthenticatedCredential, *authenticationError) {

	ac := credentials.AuthenticatedCredential{}
	if len(username) == 0 || len(password) == 0 {
		metricsRegistry.Counter("malformed").Inc()
		return nil, malformed(fmt.Errorf("username or password missing"))
	}

	err := ac.Base64URLDecode(username, password)
	if err != nil {
		metricsRegistry.Counter("malformed").Inc()
		return nil, malformed(err)
	}

	err = cm.Verify(&ac)
	if err != nil {
		metricsRegistry.Counter("invalid").Inc()
		return nil, invalid(err)
	}

	// Grab the timestamp and make sure the credential is recent enough
	ts := time.Unix(ac.Credential.Timestamp, 0)
	now := time.Now()
	if ts.Before(now) && now.Sub(ts) > authValidityWindow {
		metricsRegistry.Counter("expired").Inc()
		return nil, expired()
	}

	metricsRegistry.Counter("valid").Inc()
	return &ac, nil
}

func InitAuth(credentialManager *credentials.CredentialManager, validityWindow time.Duration) {
	authValidityWindow = validityWindow
	cm = credentialManager
	metricsRegistry = metrics.NewMetricsRegistry("authentication")
}

func DeinitAuth() {
	authValidityWindow = 0
	cm = nil
	metricsRegistry = nil
}
