package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseSubscriptionPaymentSucceededWebhook(form url.Values) (SubscriptionPaymentSucceeded, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return SubscriptionPaymentSucceeded{}, errors.WithStack(err)
	}
	var suw subscriptionPaymentSucceededWebhook
	if err := decoder.Decode(&suw, form); err != nil {
		return SubscriptionPaymentSucceeded{}, errors.WithStack(err)
	}
	sc := SubscriptionPaymentSucceeded{
		AlertName:             Alert(suw.AlertName),
		AlertID:               suw.AlertID,
		BalanceCurrency:       suw.BalanceCurrency,
		BalanceEarnings:       suw.BalanceEarnings,
		BalanceFee:            suw.BalanceFee,
		BalanceGross:          suw.BalanceGross,
		BalanceTax:            suw.BalanceTax,
		CheckoutID:            suw.CheckoutID,
		Country:               suw.Country,
		Coupon:                suw.Coupon,
		Currency:              suw.Currency,
		CustomerName:          suw.CustomerName,
		Earnings:              suw.Earnings,
		Email:                 suw.Email,
		EventTime:             suw.EventTime.Time(),
		Fee:                   suw.Fee,
		InitialPayment:        bool(suw.InitialPayment),
		Instalments:           suw.Instalments,
		MarketingConsent:      bool(suw.MarketingConsent),
		NextBillDate:          suw.NextBillDate.Time(),
		NextPaymentAmount:     suw.NextPaymentAmount,
		OrderID:               suw.OrderID,
		Passthrough:           suw.Passthrough,
		PaymentMethod:         PaymentMethod(suw.PaymentMethod),
		PaymentTax:            suw.PaymentTax,
		PlanName:              suw.PlanName,
		Quantity:              int(suw.Quantity),
		ReceiptURL:            suw.ReceiptURL,
		SaleGross:             suw.SaleGross,
		Status:                Status(suw.Status),
		SubscriptionID:        suw.SubscriptionID,
		SubscriptionPaymentID: suw.SubscriptionPaymentID,
		SubscriptionPlanID:    suw.SubscriptionPlanID,
		UnitPrice:             suw.UnitPrice,
		UserID:                suw.UserID,
	}
	return sc, nil
}

type SubscriptionPaymentSucceeded struct {
	AlertName             Alert
	AlertID               string
	BalanceCurrency       string
	BalanceEarnings       string
	BalanceFee            string
	BalanceGross          string
	BalanceTax            string
	CheckoutID            string
	Country               string
	Coupon                string
	Currency              string
	CustomerName          string
	Earnings              string
	Email                 string
	EventTime             time.Time
	Fee                   string
	InitialPayment        bool
	Instalments           string
	MarketingConsent      bool
	NextBillDate          time.Time
	NextPaymentAmount     string
	OrderID               string
	Passthrough           string
	PaymentMethod         PaymentMethod
	PaymentTax            string
	PlanName              string
	Quantity              int
	ReceiptURL            string
	SaleGross             string
	Status                Status
	SubscriptionID        string
	SubscriptionPaymentID string
	SubscriptionPlanID    string
	UnitPrice             string
	UserID                string
}

type subscriptionPaymentSucceededWebhook struct {
	AlertName             string     `schema:"alert_name"`
	AlertID               string     `schema:"alert_id"`
	BalanceCurrency       string     `schema:"balance_currency"`
	BalanceEarnings       string     `schema:"balance_earnings"`
	BalanceFee            string     `schema:"balance_fee"`
	BalanceGross          string     `schema:"balance_gross"`
	BalanceTax            string     `schema:"balance_tax"`
	CheckoutID            string     `schema:"checkout_id"`
	Country               string     `schema:"country"`
	Coupon                string     `schema:"coupon"`
	Currency              string     `schema:"currency"`
	CustomerName          string     `schema:"customer_name"`
	Earnings              string     `schema:"earnings"`
	Email                 string     `schema:"email"`
	EventTime             customTime `schema:"event_time"`
	Fee                   string     `schema:"fee"`
	InitialPayment        customBool `schema:"initial_payment"`
	Instalments           string     `schema:"instalments"`
	MarketingConsent      customBool `schema:"marketing_consent"`
	NextBillDate          customDate `schema:"next_bill_date"`
	NextPaymentAmount     string     `schema:"next_payment_amount"`
	OrderID               string     `schema:"order_id"`
	Passthrough           string     `schema:"passthrough"`
	PaymentMethod         string     `schema:"payment_method"`
	PaymentTax            string     `schema:"payment_tax"`
	PlanName              string     `schema:"plan_name"`
	Quantity              customInt  `schema:"quantity"`
	ReceiptURL            string     `schema:"receipt_url"`
	SaleGross             string     `schema:"sale_gross"`
	Status                string     `schema:"status"`
	SubscriptionID        string     `schema:"subscription_id"`
	SubscriptionPaymentID string     `schema:"subscription_payment_id"`
	SubscriptionPlanID    string     `schema:"subscription_plan_id"`
	UnitPrice             string     `schema:"unit_price"`
	UserID                string     `schema:"user_id"`
	Signature             string     `schema:"p_signature"`
}
