package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseSubscriptionCancelledWebhook(form url.Values) (SubscriptionCancelled, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return SubscriptionCancelled{}, errors.WithStack(err)
	}
	var scw subscriptionCancelledWebhook
	if err := decoder.Decode(&scw, form); err != nil {
		return SubscriptionCancelled{}, errors.WithStack(err)
	}
	sc := SubscriptionCancelled{
		AlertName:                 Alert(scw.AlertName),
		AlertID:                   scw.AlertID,
		CancellationEffectiveDate: scw.CancellationEffectiveDate.Time(),
		CheckoutID:                scw.CheckoutID,
		Currency:                  scw.Currency,
		Email:                     scw.Email,
		EventTime:                 scw.EventTime.Time(),
		MarketingConsent:          bool(scw.MarketingConsent),
		Passthrough:               scw.Passthrough,
		Quantity:                  int(scw.Quantity),
		Status:                    Status(scw.Status),
		SubscriptionID:            scw.SubscriptionID,
		SubscriptionPlanID:        scw.SubscriptionPlanID,
		UnitPrice:                 scw.UnitPrice,
		UserID:                    scw.UserID,
	}
	return sc, nil
}

type SubscriptionCancelled struct {
	AlertName                 Alert
	AlertID                   string
	CancellationEffectiveDate time.Time
	CheckoutID                string
	Currency                  string
	Email                     string
	EventTime                 time.Time
	MarketingConsent          bool
	Passthrough               string
	Quantity                  int
	Status                    Status
	SubscriptionID            string
	SubscriptionPlanID        string
	UnitPrice                 string
	UserID                    string
}

type subscriptionCancelledWebhook struct {
	AlertName                 string     `schema:"alert_name"`
	AlertID                   string     `schema:"alert_id"`
	CancellationEffectiveDate customDate `schema:"cancellation_effective_date"`
	CheckoutID                string     `schema:"checkout_id"`
	Currency                  string     `schema:"currency"`
	Email                     string     `schema:"email"`
	EventTime                 customTime `schema:"event_time"`
	MarketingConsent          customBool `schema:"marketing_consent"`
	Passthrough               string     `schema:"passthrough"`
	Quantity                  customInt  `schema:"quantity"`
	Status                    string     `schema:"status"`
	SubscriptionID            string     `schema:"subscription_id"`
	SubscriptionPlanID        string     `schema:"subscription_plan_id"`
	UnitPrice                 string     `schema:"unit_price"`
	UserID                    string     `schema:"user_id"`
	Signature                 string     `schema:"p_signature"`
}
