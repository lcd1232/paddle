package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// PaymentDisputeStatus defines the status of a payment dispute (chargeback).
type PaymentDisputeStatus string

const (
	// PaymentDisputeStatusOpen means the dispute is ongoing and awaiting a
	// decision from the bank/payment processor.
	PaymentDisputeStatusOpen PaymentDisputeStatus = "open"
	// PaymentDisputeStatusClosed means the dispute has been resolved.
	PaymentDisputeStatusClosed PaymentDisputeStatus = "closed"
)

// ParsePaymentDisputeCreatedWebhook verifies the signature of a
// payment_dispute_created webhook and parses its form values into a
// PaymentDisputeCreated. It returns an error if the signature is invalid.
func (c *WebhookClient) ParsePaymentDisputeCreatedWebhook(form url.Values) (PaymentDisputeCreated, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return PaymentDisputeCreated{}, errors.WithStack(err)
	}
	var webhook paymentDisputeCreatedWebhook
	if err := decoder.Decode(&webhook, form); err != nil {
		return PaymentDisputeCreated{}, errors.WithStack(err)
	}
	d := PaymentDisputeCreated{
		AlertName:        Alert(webhook.AlertName),
		AlertID:          webhook.AlertID,
		Amount:           webhook.Amount,
		BalanceAmount:    webhook.BalanceAmount,
		BalanceCurrency:  webhook.BalanceCurrency,
		BalanceFee:       webhook.BalanceFee,
		CheckoutID:       webhook.CheckoutID,
		Currency:         webhook.Currency,
		Email:            webhook.Email,
		EventTime:        webhook.EventTime.Time(),
		FeeUSD:           webhook.FeeUSD,
		MarketingConsent: bool(webhook.MarketingConsent),
		OrderID:          webhook.OrderID,
		Passthrough:      webhook.Passthrough,
		Status:           PaymentDisputeStatus(webhook.Status),
	}
	return d, nil
}

// PaymentDisputeCreated is the payload of a payment_dispute_created webhook,
// sent when a customer disputes (charges back) a payment.
type PaymentDisputeCreated struct {
	AlertName        Alert
	AlertID          string
	Amount           string
	BalanceAmount    string
	BalanceCurrency  string
	BalanceFee       string
	CheckoutID       string
	Currency         string
	Email            string
	EventTime        time.Time
	FeeUSD           string
	MarketingConsent bool
	OrderID          string
	// Passthrough carries the seller-defined data sent with the checkout.
	Passthrough string
	Status      PaymentDisputeStatus
}

type paymentDisputeCreatedWebhook struct {
	AlertName        string     `schema:"alert_name"`
	AlertID          string     `schema:"alert_id"`
	Amount           string     `schema:"amount"`
	BalanceAmount    string     `schema:"balance_amount"`
	BalanceCurrency  string     `schema:"balance_currency"`
	BalanceFee       string     `schema:"balance_fee"`
	CheckoutID       string     `schema:"checkout_id"`
	Currency         string     `schema:"currency"`
	Email            string     `schema:"email"`
	EventTime        customTime `schema:"event_time"`
	FeeUSD           string     `schema:"fee_usd"`
	MarketingConsent customBool `schema:"marketing_consent"`
	OrderID          string     `schema:"order_id"`
	Passthrough      string     `schema:"passthrough"`
	Status           string     `schema:"status"`
	Signature        string     `schema:"p_signature"`
}
