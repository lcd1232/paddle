package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseSubscriptionUpdatedWebhook(form url.Values) (SubscriptionUpdated, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return SubscriptionUpdated{}, errors.WithStack(err)
	}
	var suw subscriptionUpdatedWebhook
	if err := decoder.Decode(&suw, form); err != nil {
		return SubscriptionUpdated{}, errors.WithStack(err)
	}
	sc := SubscriptionUpdated{
		AlertName:             Alert(suw.AlertName),
		AlertID:               suw.AlertID,
		UpdateURL:             suw.UpdateURL,
		CancelURL:             suw.CancelURL,
		CheckoutID:            suw.CheckoutID,
		Currency:              suw.Currency,
		Email:                 suw.Email,
		EventTime:             time.Time(suw.EventTime),
		MarketingConsent:      bool(suw.MarketingConsent),
		NewPrice:              suw.NewPrice,
		NewQuantity:           int(suw.NewQuantity),
		NewUnitPrice:          suw.NewUnitPrice,
		NextBillDate:          time.Time(suw.NextBillDate),
		Passthrough:           suw.Passthrough,
		Status:                Status(suw.Status),
		SubscriptionID:        suw.SubscriptionID,
		SubscriptionPlanID:    suw.SubscriptionPlanID,
		UserID:                suw.UserID,
		OldNextBillDate:       time.Time(suw.OldNextBillDate),
		OldPrice:              suw.OldPrice,
		OldQuantity:           int(suw.OldQuantity),
		OldStatus:             Status(suw.OldStatus),
		OldSubscriptionPlanID: suw.OldSubscriptionPlanID,
		OldUnitPrice:          suw.OldUnitPrice,
		PausedAt:              time.Time(suw.PausedAt),
		PausedFrom:            time.Time(suw.PausedFrom),
		PausedReason:          PausedReason(suw.PausedReason),
	}
	return sc, nil
}

type SubscriptionUpdated struct {
	AlertName Alert
	AlertID   string

	UpdateURL string
	CancelURL string

	CheckoutID       string
	Currency         string
	Email            string
	EventTime        time.Time
	MarketingConsent bool

	NewPrice     string
	NewQuantity  int
	NewUnitPrice string
	NextBillDate time.Time

	Passthrough        string
	Status             Status
	SubscriptionID     string
	SubscriptionPlanID string
	UserID             string

	OldNextBillDate       time.Time
	OldPrice              string
	OldQuantity           int
	OldStatus             Status
	OldSubscriptionPlanID string
	OldUnitPrice          string

	PausedAt     time.Time
	PausedFrom   time.Time
	PausedReason PausedReason
}

type subscriptionUpdatedWebhook struct {
	AlertName             string     `schema:"alert_name"`
	AlertID               string     `schema:"alert_id"`
	CancelURL             string     `schema:"cancel_url"`
	CheckoutID            string     `schema:"checkout_id"`
	Email                 string     `schema:"email"`
	EventTime             customTime `schema:"event_time"`
	MarketingConsent      customBool `schema:"marketing_consent"`
	NewPrice              string     `schema:"new_price"`
	NewQuantity           customInt  `schema:"new_quantity"`
	NewUnitPrice          string     `schema:"new_unit_price"`
	NextBillDate          customDate `schema:"next_bill_date"`
	OldPrice              string     `schema:"old_price"`
	OldQuantity           customInt  `schema:"old_quantity"`
	OldUnitPrice          string     `schema:"old_unit_price"`
	Currency              string     `schema:"currency"`
	Passthrough           string     `schema:"passthrough"`
	Status                string     `schema:"status"`
	SubscriptionID        string     `schema:"subscription_id"`
	SubscriptionPlanID    string     `schema:"subscription_plan_id"`
	UserID                string     `schema:"user_id"`
	UpdateURL             string     `schema:"update_url"`
	OldNextBillDate       customDate `schema:"old_next_bill_date"`
	OldStatus             string     `schema:"old_status"`
	OldSubscriptionPlanID string     `schema:"old_subscription_plan_id"`
	PausedAt              customTime `schema:"paused_at"`
	PausedFrom            customDate `schema:"paused_from"`
	PausedReason          string     `schema:"paused_reason"`
	Signature             string     `schema:"p_signature"`
}
