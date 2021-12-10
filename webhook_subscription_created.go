package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseSubscriptionCreatedWebhook(form url.Values) (SubscriptionCreated, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return SubscriptionCreated{}, errors.WithStack(err)
	}
	var scw subscriptionCreatedWebhook
	if err := decoder.Decode(&scw, form); err != nil {
		return SubscriptionCreated{}, errors.WithStack(err)
	}
	sc := SubscriptionCreated{
		AlertName:          Alert(scw.AlertName),
		AlertID:            scw.AlertID,
		CancelURL:          scw.CancelURL,
		CheckoutID:         scw.CheckoutID,
		Currency:           scw.Currency,
		Email:              scw.Email,
		EventTime:          scw.EventTime.Time(),
		MarketingConsent:   bool(scw.MarketingConsent),
		NextBillDate:       time.Time(scw.NextBillDate),
		Passthrough:        scw.Passthrough,
		Quantity:           int(scw.Quantity),
		Source:             scw.Source,
		Status:             Status(scw.Status),
		SubscriptionID:     scw.SubscriptionID,
		SubscriptionPlanID: scw.SubscriptionPlanID,
		UnitPrice:          scw.UnitPrice,
		UserID:             scw.UserID,
		UpdateURL:          scw.UpdateURL,
	}
	return sc, nil
}

type SubscriptionCreated struct {
	AlertName          Alert
	AlertID            string
	CancelURL          string
	CheckoutID         string
	Currency           string
	Email              string
	EventTime          time.Time
	MarketingConsent   bool
	NextBillDate       time.Time
	Passthrough        string
	Quantity           int
	Source             string
	Status             Status
	SubscriptionID     string
	SubscriptionPlanID string
	UnitPrice          string
	UserID             string
	UpdateURL          string
}

type subscriptionCreatedWebhook struct {
	AlertName          string     `schema:"alert_name"`
	AlertID            string     `schema:"alert_id"`
	CancelURL          string     `schema:"cancel_url"`
	CheckoutID         string     `schema:"checkout_id"`
	Currency           string     `schema:"currency"`
	Email              string     `schema:"email"`
	EventTime          customTime `schema:"event_time"`
	MarketingConsent   customBool `schema:"marketing_consent"`
	NextBillDate       customDate `schema:"next_bill_date"`
	Passthrough        string     `schema:"passthrough"`
	Quantity           customInt  `schema:"quantity"`
	Source             string     `schema:"source"`
	Status             string     `schema:"status"`
	SubscriptionID     string     `schema:"subscription_id"`
	SubscriptionPlanID string     `schema:"subscription_plan_id"`
	UnitPrice          string     `schema:"unit_price"`
	UserID             string     `schema:"user_id"`
	UpdateURL          string     `schema:"update_url"`
	Signature          string     `schema:"p_signature"`
}
