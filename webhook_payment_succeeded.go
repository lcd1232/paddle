package paddle

import (
	"net/url"
	"time"
)

func (c *Client) ParsePaymentSucceededWebhook(form url.Values) (PaymentSucceeded, error) {
	//signature := form.Get(signatureKey)
	//if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
	//	return PaymentSucceeded{}, errors.WithStack(err)
	//}
	//var suw subscriptionPaymentSucceededWebhook
	//if err := decoder.Decode(&suw, form); err != nil {
	//	return SubscriptionPaymentSucceeded{}, errors.WithStack(err)
	//}
	//sc := PaymentSucceeded{
	//	AlertName:             Alert(suw.AlertName),
	//	AlertID:               suw.AlertID,
	//	BalanceCurrency:       suw.BalanceCurrency,
	//	BalanceEarnings:       suw.BalanceEarnings,
	//	BalanceFee:            suw.BalanceFee,
	//	BalanceGross:          suw.BalanceGross,
	//	BalanceTax:            suw.BalanceTax,
	//	CheckoutID:            suw.CheckoutID,
	//	Country:               suw.Country,
	//	Coupon:                suw.Coupon,
	//	Currency:              suw.Currency,
	//	CustomerName:          suw.CustomerName,
	//	Earnings:              suw.Earnings,
	//	Email:                 suw.Email,
	//	EventTime:             time.Time(suw.EventTime),
	//	Fee:                   suw.Fee,
	//	InitialPayment:        bool(suw.InitialPayment),
	//	Instalments:           suw.Instalments,
	//	MarketingConsent:      bool(suw.MarketingConsent),
	//	NextBillDate:          time.Time(suw.NextBillDate),
	//	NextPaymentAmount:     suw.NextPaymentAmount,
	//	OrderID:               suw.OrderID,
	//	Passthrough:           suw.Passthrough,
	//	PaymentMethod:         PaymentMethod(suw.PaymentMethod),
	//	PaymentTax:            suw.PaymentTax,
	//	PlanName:              suw.PlanName,
	//	Quantity:              int(suw.Quantity),
	//	ReceiptURL:            suw.ReceiptURL,
	//	SaleGross:             suw.SaleGross,
	//	Status:                Status(suw.Status),
	//	SubscriptionID:        suw.SubscriptionID,
	//	SubscriptionPaymentID: suw.SubscriptionPaymentID,
	//	SubscriptionPlanID:    suw.SubscriptionPlanID,
	//	UnitPrice:             suw.UnitPrice,
	//	UserID:                suw.UserID,
	//}
	//return sc, nil
	return PaymentSucceeded{}, nil
}

type PaymentSucceeded struct {
	AlertName         Alert
	AlertID           string
	BalanceCurrency   string
	BalanceEarnings   string
	BalanceFee        string
	BalanceGross      string
	BalanceTax        string
	CheckoutID        string
	Country           string
	Coupon            string
	Currency          string
	CustomerName      string
	Earnings          string
	Email             string
	EventTime         time.Time
	Fee               string
	Instalments       string
	MarketingConsent  bool
	OrderID           string
	IP                string
	Passthrough       string
	PaymentMethod     PaymentMethod
	PaymentTax        string
	ProductID         string
	ProductName       string
	Quantity          int
	ReceiptURL        string
	SaleGross         string
	UsedPriceOverride bool
}

type paymentSucceededWebhook struct {
	AlertName         string     `schema:"alert_name"`
	AlertID           string     `schema:"alert_id"`
	BalanceCurrency   string     `schema:"balance_currency"`
	BalanceEarnings   string     `schema:"balance_earnings"`
	BalanceFee        string     `schema:"balance_fee"`
	BalanceGross      string     `schema:"balance_gross"`
	BalanceTax        string     `schema:"balance_tax"`
	CheckoutID        string     `schema:"checkout_id"`
	Country           string     `schema:"country"`
	Coupon            string     `schema:"coupon"`
	Currency          string     `schema:"currency"`
	CustomerName      string     `schema:"customer_name"`
	Earnings          string     `schema:"earnings"`
	Email             string     `schema:"email"`
	EventTime         customTime `schema:"event_time"`
	Fee               string     `schema:"fee"`
	Instalments       string     `schema:"instalments"`
	MarketingConsent  customBool `schema:"marketing_consent"`
	OrderID           string     `schema:"order_id"`
	IP                string     `schema:"ip"`
	Passthrough       string     `schema:"passthrough"`
	PaymentMethod     string     `schema:"payment_method"`
	PaymentTax        string     `schema:"payment_tax"`
	ProductID         string     `schema:"product_id"`
	ProductName       string     `schema:"product_name"`
	Quantity          customInt  `schema:"quantity"`
	ReceiptURL        string     `schema:"receipt_url"`
	SaleGross         string     `schema:"sale_gross"`
	UsedPriceOverride customBool `schema:"used_price_override"`
	Signature         string     `schema:"p_signature"`
}
