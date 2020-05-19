package paddle

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"

	"github.com/pkg/errors"
)

var (
	ErrInvalidSignature = errors.New("paddle: invalid signature")
)

const signatureKey = "p_signature"

// verifier verifies signature
type verifier interface {
	Verify(publicKey *rsa.PublicKey, signature string, form url.Values) error
}

type noonVerifier struct{}

func (*noonVerifier) Verify(publicKey *rsa.PublicKey, signature string, form url.Values) error {
	return nil
}

type realVerifier struct{}

func (rv *realVerifier) Verify(publicKey *rsa.PublicKey, signature string, form url.Values) error {

	// get sha1 hash of php encoded form values
	b := rv.phpEncodeValues(form)
	h := sha1.New()
	if _, err := h.Write(b); err != nil {
		return errors.WithStack(err)
	}
	hash := h.Sum(nil)

	// decode signature
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	// verify
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, hash, decodedSignature)
	if err != nil {
		if errors.Cause(err) == rsa.ErrVerification {
			return errors.WithStack(ErrInvalidSignature)
		}
		return err
	}

	return nil
}

func (rv *realVerifier) phpEncodeValues(formValues url.Values) []byte {
	var keys []string
	for k := range formValues {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var phpEncoded = fmt.Sprintf("a:%d:{", len(keys)-1)
	for _, k := range keys {
		if k != signatureKey {
			val := formValues[k][0]
			phpEncoded += fmt.Sprintf("s:%d:\"%s\";s:%d:\"%s\";", len(k), k, len(val), val)
		}
	}
	phpEncoded += "}"

	return []byte(phpEncoded)
}
