package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParsePaymentSucceededWebhook(form url.Values) (PaymentSucceeded, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return PaymentSucceeded{}, errors.WithStack(err)
	}
	var psw paymentSucceededWebhook
	if err := decoder.Decode(&psw, form); err != nil {
		return PaymentSucceeded{}, errors.WithStack(err)
	}
	ps := PaymentSucceeded{
		AlertName:         Alert(psw.AlertName),
		AlertID:           psw.AlertID,
		BalanceCurrency:   psw.BalanceCurrency,
		BalanceEarnings:   psw.BalanceEarnings,
		BalanceFee:        psw.BalanceFee,
		BalanceGross:      psw.BalanceGross,
		BalanceTax:        psw.BalanceTax,
		CheckoutID:        psw.CheckoutID,
		Country:           psw.Country,
		Coupon:            psw.Coupon,
		Currency:          psw.Currency,
		CustomerName:      psw.CustomerName,
		Earnings:          psw.Earnings,
		Email:             psw.Email,
		EventTime:         psw.EventTime.Time(),
		Fee:               psw.Fee,
		MarketingConsent:  bool(psw.MarketingConsent),
		OrderID:           psw.OrderID,
		IP:                psw.IP,
		Passthrough:       psw.Passthrough,
		PaymentMethod:     PaymentMethod(psw.PaymentMethod),
		PaymentTax:        psw.PaymentTax,
		ProductID:         psw.ProductID,
		ProductName:       psw.ProductName,
		Quantity:          int(psw.Quantity),
		ReceiptURL:        psw.ReceiptURL,
		SaleGross:         psw.SaleGross,
		UsedPriceOverride: bool(psw.UsedPriceOverride),
	}
	return ps, nil
}

type PaymentSucceeded struct {
	AlertName        Alert
	AlertID          string
	BalanceCurrency  string
	BalanceEarnings  string
	BalanceFee       string
	BalanceGross     string
	BalanceTax       string
	CheckoutID       string
	Country          string
	Coupon           string
	Currency         string
	CustomerName     string
	Earnings         string
	Email            string
	EventTime        time.Time
	Fee              string
	MarketingConsent bool
	OrderID          string
	// IP defines user IP.
	// Deprecated, see https://developer.paddle.com/webhook-reference/one-off-purchase-alerts/payment-succeeded
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
