package paddle

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

var decoder *schema.Decoder

func init() {
	decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
}

type customInt int

func (ci *customInt) UnmarshalText(text []byte) error {
	s := string(text)
	i, err := strconv.Atoi(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*ci = customInt(i)
	return nil
}

func GetAlertName(form url.Values) (Alert, error) {
	var an alertName
	if err := decoder.Decode(&an, form); err != nil {
		return "", errors.WithStack(err)
	}
	return Alert(an.AlertName), nil
}

type WebhookClient struct {
	publicKey *rsa.PublicKey
	verifier  verifier
}

func NewWebhookClient(publicKey string) (*WebhookClient, error) {
	pubKey, err := parsePublicKey(publicKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &WebhookClient{
		publicKey: pubKey,
		verifier:  new(realVerifier),
	}, nil
}

func (c *WebhookClient) SetVerification(b bool) {
	if b {
		c.verifier = new(realVerifier)
	} else {
		c.verifier = new(noonVerifier)
	}
}

func parsePublicKey(pubKey string) (*rsa.PublicKey, error) {
	decoded, _ := pem.Decode([]byte(pubKey))
	if decoded == nil {
		return nil, errors.New("invalid public key")
	}
	re, err := x509.ParsePKIXPublicKey(decoded.Bytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pub := re.(*rsa.PublicKey)
	return pub, nil
}
