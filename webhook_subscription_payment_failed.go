package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseSubscriptionPaymentFailedWebhook(form url.Values) (SubscriptionPaymentFailed, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return SubscriptionPaymentFailed{}, errors.WithStack(err)
	}
	var spfw subscriptionPaymentFailedWebhook
	if err := decoder.Decode(&spfw, form); err != nil {
		return SubscriptionPaymentFailed{}, errors.WithStack(err)
	}
	spf := SubscriptionPaymentFailed{
		AlertName:             Alert(spfw.AlertName),
		AlertID:               spfw.AlertID,
		Amount:                spfw.Amount,
		CancelURL:             spfw.CancelURL,
		CheckoutID:            spfw.CheckoutID,
		Currency:              spfw.Currency,
		Email:                 spfw.Email,
		EventTime:             spfw.EventTime.Time(),
		MarketingConsent:      bool(spfw.MarketingConsent),
		NextRetryDate:         spfw.NextRetryDate.Time(),
		Passthrough:           spfw.Passthrough,
		Quantity:              int(spfw.Quantity),
		Status:                Status(spfw.Status),
		SubscriptionID:        spfw.SubscriptionID,
		SubscriptionPlanID:    spfw.SubscriptionPlanID,
		UnitPrice:             spfw.UnitPrice,
		UpdateURL:             spfw.UpdateURL,
		SubscriptionPaymentID: spfw.SubscriptionPaymentID,
		Installments:          int(spfw.Installments),
		OrderID:               spfw.OrderID,
		UserID:                spfw.UserID,
		AttemptNumber:         int(spfw.AttemptNumber),
	}
	return spf, nil
}

type SubscriptionPaymentFailed struct {
	AlertName             Alert
	AlertID               string
	Amount                string
	CancelURL             string
	CheckoutID            string
	Currency              string
	Email                 string
	EventTime             time.Time
	MarketingConsent      bool
	NextRetryDate         time.Time
	Passthrough           string
	Quantity              int
	Status                Status
	SubscriptionID        string
	SubscriptionPlanID    string
	UnitPrice             string
	UpdateURL             string
	SubscriptionPaymentID string
	Installments          int
	OrderID               string
	UserID                string
	AttemptNumber         int
}

type subscriptionPaymentFailedWebhook struct {
	AlertName             string     `schema:"alert_name"`
	AlertID               string     `schema:"alert_id"`
	Amount                string     `schema:"amount"`
	CancelURL             string     `schema:"cancel_url"`
	CheckoutID            string     `schema:"checkout_id"`
	Currency              string     `schema:"currency"`
	Email                 string     `schema:"email"`
	EventTime             customTime `schema:"event_time"`
	MarketingConsent      customBool `schema:"marketing_consent"`
	NextRetryDate         customDate `schema:"next_retry_date"`
	Passthrough           string     `schema:"passthrough"`
	Quantity              customInt  `schema:"quantity"`
	Status                string     `schema:"status"`
	SubscriptionID        string     `schema:"subscription_id"`
	SubscriptionPlanID    string     `schema:"subscription_plan_id"`
	UnitPrice             string     `schema:"unit_price"`
	UpdateURL             string     `schema:"update_url"`
	SubscriptionPaymentID string     `schema:"subscription_payment_id"`
	Installments          customInt  `schema:"installments"`
	OrderID               string     `schema:"order_id"`
	UserID                string     `schema:"user_id"`
	AttemptNumber         customInt  `schema:"attempt_number"`
	Signature             string     `schema:"p_signature"`
}
