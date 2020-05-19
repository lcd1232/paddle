package paddle

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

var decoder *schema.Decoder

func init() {
	decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
}

type customBool bool

func (cb *customBool) UnmarshalText(text []byte) error {
	s := string(text)
	b, err := parseBool(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*cb = customBool(b)
	return nil
}

type customTime time.Time

func (ct *customTime) UnmarshalText(text []byte) error {
	s := string(text)
	t, err := parseTime(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*ct = customTime(t)
	return nil
}

type customDate time.Time

func (cd *customDate) UnmarshalText(text []byte) error {
	s := string(text)
	t, err := parseDate(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*cd = customDate(t)
	return nil
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

func parseTime(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return t, nil
}

func parseDate(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return t, nil
}

func parseBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return b, nil
}

type Client struct {
	publicKey *rsa.PublicKey
	verifier  verifier
}

func NewClient(publicKey string) (*Client, error) {
	pubKey, err := parsePublicKey(publicKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Client{
		publicKey: pubKey,
		verifier:  new(realVerifier),
	}, nil
}

func (c *Client) SetVerification(b bool) {
	if b {
		c.verifier = new(realVerifier)
	} else {
		c.verifier = new(noonVerifier)
	}
}

func parsePublicKey(pubKey string) (*rsa.PublicKey, error) {
	decoded, _ := pem.Decode([]byte(pubKey))
	re, err := x509.ParsePKIXPublicKey(decoded.Bytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pub := re.(*rsa.PublicKey)
	return pub, nil
}
