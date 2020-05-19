package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *Client) ParseSubscriptionPaymentRefundedWebhook(form url.Values) (SubscriptionPaymentRefunded, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return SubscriptionPaymentRefunded{}, errors.WithStack(err)
	}
	var webhook subscriptionPaymentRefundedWebhook
	if err := decoder.Decode(&webhook, form); err != nil {
		return SubscriptionPaymentRefunded{}, errors.WithStack(err)
	}
	sr := SubscriptionPaymentRefunded{
		AlertName:               Alert(webhook.AlertName),
		AlertID:                 webhook.AlertID,
		Amount:                  webhook.Amount,
		BalanceCurrency:         webhook.BalanceCurrency,
		BalanceEarningsDecrease: webhook.BalanceEarningsDecrease,
		BalanceFeeRefund:        webhook.BalanceFeeRefund,
		BalanceGrossRefund:      webhook.BalanceGrossRefund,
		BalanceTaxRefund:        webhook.BalanceTaxRefund,
		CheckoutID:              webhook.CheckoutID,
		Currency:                webhook.Currency,
		EarningsDecrease:        webhook.EarningsDecrease,
		Email:                   webhook.Email,
		EventTime:               time.Time(webhook.EventTime),
		FeeRefund:               webhook.FeeRefund,
		GrossRefund:             webhook.GrossRefund,
		InitialPayment:          bool(webhook.InitialPayment),
		Instalments:             webhook.Instalments,
		MarketingConsent:        bool(webhook.MarketingConsent),
		OrderID:                 webhook.OrderID,
		Passthrough:             webhook.Passthrough,
		Quantity:                int(webhook.Quantity),
		RefundReason:            webhook.RefundReason,
		RefundType:              RefundType(webhook.RefundType),
		Status:                  Status(webhook.Status),
		SubscriptionID:          webhook.SubscriptionID,
		SubscriptionPaymentID:   webhook.SubscriptionPaymentID,
		SubscriptionPlanID:      webhook.SubscriptionPlanID,
		TaxRefund:               webhook.TaxRefund,
		UnitPrice:               webhook.UnitPrice,
		UserID:                  webhook.UserID,
	}
	return sr, nil
}

type SubscriptionPaymentRefunded struct {
	AlertName               Alert
	AlertID                 string
	Amount                  string
	BalanceCurrency         string
	BalanceEarningsDecrease string
	BalanceFeeRefund        string
	BalanceGrossRefund      string
	BalanceTaxRefund        string
	CheckoutID              string
	Currency                string
	EarningsDecrease        string
	Email                   string
	EventTime               time.Time
	FeeRefund               string
	GrossRefund             string
	InitialPayment          bool
	Instalments             string
	MarketingConsent        bool
	OrderID                 string
	Passthrough             string
	Quantity                int
	RefundReason            string
	RefundType              RefundType
	Status                  Status
	SubscriptionID          string
	SubscriptionPaymentID   string
	SubscriptionPlanID      string
	TaxRefund               string
	UnitPrice               string
	UserID                  string
}

type subscriptionPaymentRefundedWebhook struct {
	AlertName               string     `schema:"alert_name"`
	AlertID                 string     `schema:"alert_id"`
	Amount                  string     `schema:"amount"`
	BalanceCurrency         string     `schema:"balance_currency"`
	BalanceEarningsDecrease string     `schema:"balance_earnings_decrease"`
	BalanceFeeRefund        string     `schema:"balance_fee_refund"`
	BalanceGrossRefund      string     `schema:"balance_gross_refund"`
	BalanceTaxRefund        string     `schema:"balance_tax_refund"`
	CheckoutID              string     `schema:"checkout_id"`
	Currency                string     `schema:"currency"`
	EarningsDecrease        string     `schema:"earnings_decrease"`
	Email                   string     `schema:"email"`
	EventTime               customTime `schema:"event_time"`
	FeeRefund               string     `schema:"fee_refund"`
	GrossRefund             string     `schema:"gross_refund"`
	InitialPayment          customBool `schema:"initial_payment"`
	Instalments             string     `schema:"instalments"`
	MarketingConsent        customBool `schema:"marketing_consent"`
	OrderID                 string     `schema:"order_id"`
	Passthrough             string     `schema:"passthrough"`
	Quantity                customInt  `schema:"quantity"`
	RefundReason            string     `schema:"refund_reason"`
	RefundType              string     `schema:"refund_type"`
	Status                  string     `schema:"status"`
	SubscriptionID          string     `schema:"subscription_id"`
	SubscriptionPaymentID   string     `schema:"subscription_payment_id"`
	SubscriptionPlanID      string     `schema:"subscription_plan_id"`
	TaxRefund               string     `schema:"tax_refund"`
	UnitPrice               string     `schema:"unit_price"`
	UserID                  string     `schema:"user_id"`
	Signature               string     `schema:"p_signature"`
}
