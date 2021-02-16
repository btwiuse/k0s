package basicauth

import (
	"bytes"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/lunny/tango"
)

type Auther interface {
	AskAuth() bool
}

type NoAuth struct {
}

func (NoAuth) AskAuth() bool {
	return false
}

type BasicAuth struct {
	Realm    string
	User     string
	Password string
}

func New(name, pass string) *BasicAuth {
	return &BasicAuth{"Restricted", name, pass}
}

func (b *BasicAuth) authenticate(r *http.Request) bool {
	const basicScheme string = "Basic "

	// Confirm the request is sending Basic Authentication credentials.
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, basicScheme) {
		return false
	}

	// Get the plain-text username and password from the request
	// The first six characters are skipped - e.g. "Basic ".
	str, err := base64.StdEncoding.DecodeString(auth[len(basicScheme):])
	if err != nil {
		return false
	}

	// Split on the first ":" character only, with any subsequent colons assumed to be part
	// of the password. Note that the RFC2617 standard does not place any limitations on
	// allowable characters in the password.
	creds := bytes.SplitN(str, []byte(":"), 2)

	// Equalize lengths of supplied and required credentials
	// by hashing them
	givenUser := sha256.Sum256(creds[0])
	givenPass := sha256.Sum256(creds[1])
	requiredUser := sha256.Sum256([]byte(b.User))
	requiredPass := sha256.Sum256([]byte(b.Password))

	if len(creds) != 2 {
		return false
	}

	// Compare the supplied credentials to those set in our options
	if subtle.ConstantTimeCompare(givenUser[:], requiredUser[:]) == 1 &&
		subtle.ConstantTimeCompare(givenPass[:], requiredPass[:]) == 1 {
		return true
	}

	return false
}

func (b *BasicAuth) Handle(ctx *tango.Context) {
	if action := ctx.Action(); action != nil {
		if a, ok := action.(Auther); ok {
			if !a.AskAuth() {
				ctx.Next()
				return
			}
		}
	}

	if b.authenticate(ctx.Req()) {
		ctx.Next()
		return
	}

	ctx.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm=%q`, b.Realm))
	ctx.Unauthorized()
	return
}

// defaultUnauthorizedHandler provides a default HTTP 401 Unauthorized response.
func defaultUnauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
