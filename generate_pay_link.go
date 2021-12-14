package paddle

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GeneratePayLinkRequest struct {
	// ProductID is The Paddle Product ID/Plan ID that you want to base this custom checkout on.
	// Required if not using custom products.
	// If no ProductID is set, custom non-subscription product checkouts can be generated instead by specifying the required fields: Title, WebhookURL and prices.
	// Note that CouponCode cannot be used with custom products.
	ProductID int
	// Title is the name of the product/title of the checkout. Required if ProductID is not set.
	Title string
	// WebhookURL is an endpoint that we will call with transaction information upon successful checkout, to allow you to fulfill the purchase.
	// Only valid (and required) if ProductID is not set. Not valid for subscription plans.
	// Note: testing on localhost is not supported. Please use an internet-accessible URL.
	WebhookURL string
	// Prices is price (s) of the checkout for a one-time purchase or the initial payment of a subscription.
	// If ProductID is set, you must also provide the price for the product’s default currency.If a given currency is enabled in the dashboard, it will default to a conversion of the product’s default currency price set in this field unless specified here as well.
	// If the currency specified is not enabled in the dashboard and the ProductID is a subscription, the currency’s RecurringPrices must be set as well.
	Prices map[string]string
	// RecurringPrices is recurring price(s) of the checkout (excluding the initial payment) only if the ProductID specified is a subscription.
	// To override the initial payment and all recurring payment amounts, both Prices and RecurringPrices must be set.
	// You must also provide the price for the subscription’s default currency.
	// If a given currency is enabled in the dashboard, it will default to a conversion of the subscription’s default currency price set in this field unless specified here as well.
	// If the currency specified is not enabled in the dashboard, the currency’s prices must be set as well.
	RecurringPrices map[string]string
	// TrialDays is for subscription plans only.
	// The number of days for the initial billing cycle.
	// If you leave this field empty, the default trial days of the plan will be used.
	// Note: The prices might additionally need to be set in order to achieve the desired behaviour (free trial/paid trial),
	// depending on whether the plan had a trial period set originally.
	TrialDays int
	// CustomMessage is a short message displayed below the product name on the checkout.
	CustomMessage string
	// CouponCode is a coupon to be applied to the checkout.
	// Note that this cannot be used with custom products, and is only valid if a ProductID is set.
	CouponCode string
	// Discountable specifies if a coupon can be applied to the checkout.
	// “Add Coupon” button on the checkout will be hidden as well if set to false.
	Discountable *bool
	// ImageURL is a URL for the product image/icon displayed on the checkout.
	ImageURL string
	// ReturnURL is a URL to redirect to once the checkout is completed.
	// If the variable {checkout_hash} is included within the URL
	// (e.g.https://example.com/thanks?checkout={checkout_hash}), the API will automatically populate the Paddle checkout ID in the redirected URL.
	ReturnURL string
	// QuantityVariable specifies if the user is allowed to alter the quantity of the checkout.
	QuantityVariable *bool
	// Quantity pre-fills the quantity selector on the checkout.Please note that free products/subscription plans are fixed to a quantity of 1.
	// Any quantity over the maximum value will default to a quantity of 1.
	Quantity int
	// Expires specifies if the checkout link should expire.
	// The generated checkout URL will be accessible until 23:59:59 (UTC) on the date specified.
	// If this is not specified, the checkout link will automatically expire 60 days after it has been first opened.
	Expires time.Time
	// Affiliates contains other Paddle vendor IDs whom you would like to split the funds from this checkout with.
	Affiliates map[string]string
	// RecurringAffiliateLimit limits the number of times other Paddle vendors will receive funds from the recurring payments (for subscription products).
	// The initial checkout payment is included in the limit.
	// If you leave this field empty, the limit will not be applied.
	// Note: if your plan has a trial period, set this to 2 or greater in order for your affiliates to correctly receive their commission on paid payments after the trial.
	RecurringAffiliateLimit int
	// MarketingConsent specifies whether you have gathered consent to market to the customer.
	// CustomerEmail is required if this property is set, and you want to opt the customer into marketing.
	MarketingConsent bool
	// CustomerEmail pre-fills the customer email field on the checkout.
	CustomerEmail string
	// CustomerCountry pre-fills the customer country field on the checkout.
	// List of supported ISO country codes - https://developer.paddle.com/reference/platform-parameters/supported-countries
	// Only pre-filled if CustomerEmail is set.
	CustomerCountry string
	// CustomerPostcode pre-fills the customer postcode field on the checkout.
	// This field is required if the CustomerCountry requires postcode.
	// List of the countries requiring this field - https://developer.paddle.com/reference/platform-parameters/supported-countries#countries-requiring-postcode
	CustomerPostcode string
	// IsRecoverable Specifies if checkout recovery emails can be sent to users who abandon the checkout process after entering their email address.
	// An additional 10% transaction fee applies to checkouts we recover.
	// This will override the checkout recovery setting specified in your page.
	IsRecoverable *bool
	// Passthrough is a string of metadata you wish to store with the checkout.
	// Will be sent alongside all webhooks associated with the order.
	// See the documentation for more information.
	Passthrough string
	// VatNumber pre-fills the sales tax identifier (VAT number) field on the checkout.
	VatNumber string
	// VatCompanyName pre-fills the Company Name field on the checkout.
	// Required if VatNumber is set.
	VatCompanyName string
	// VatStreet pre-fills the Street field on the checkout.
	// Required if VatNumber is set.
	VatStreet string
	// VatCity pre-fills the Town/City field on the checkout.
	// Required if VatNumber is set.
	VatCity string
	// VatState pre-fills the State field on the checkout.
	VatState string
	// VatCountry pre-fills the Country field on the checkout.
	// Required if VatNumber is set.
	// List of the list of supported ISO country codes - https://developer.paddle.com/reference/platform-parameters/supported-countries
	VatCountry string
	// VatPostcode pre-fills the Postcode field on the checkout.
	VatPostcode string
}

func toGeneratePayLinkRequest(request GeneratePayLinkRequest) generatePayLinkRequest {
	prices := make([]string, 0, len(request.Prices))
	for currency, price := range request.Prices {
		prices = append(prices, fmt.Sprintf("%s:%s", currency, price))
	}

	recurringPrices := make([]string, 0, len(request.RecurringPrices))
	for currency, price := range request.RecurringPrices {
		recurringPrices = append(recurringPrices, fmt.Sprintf("%s:%s", currency, price))
	}

	affiliates := make([]string, 0, len(request.Affiliates))
	for id, percent := range request.Affiliates {
		affiliates = append(affiliates, fmt.Sprintf("%s:%s", id, percent))
	}

	return generatePayLinkRequest{
		ProductID:               request.ProductID,
		Title:                   request.Title,
		WebhookURL:              request.WebhookURL,
		Prices:                  prices,
		RecurringPrices:         recurringPrices,
		TrialDays:               request.TrialDays,
		CustomMessage:           request.CustomMessage,
		CouponCode:              request.CouponCode,
		Discountable:            toCustomBoolPointer(request.Discountable),
		ImageURL:                request.ImageURL,
		ReturnURL:               request.ReturnURL,
		QuantityVariable:        toCustomBoolPointer(request.QuantityVariable),
		Quantity:                request.Quantity,
		Expires:                 toCustomDate(request.Expires),
		Affiliates:              affiliates,
		RecurringAffiliateLimit: request.RecurringAffiliateLimit,
		MarketingConsent:        request.MarketingConsent,
		CustomerEmail:           request.CustomerEmail,
		CustomerCountry:         request.CustomerCountry,
		CustomerPostcode:        request.CustomerPostcode,
		IsRecoverable:           toCustomBoolPointer(request.IsRecoverable),
		Passthrough:             request.Passthrough,
		VatNumber:               request.VatNumber,
		VatCompanyName:          request.VatCompanyName,
		VatStreet:               request.VatStreet,
		VatCity:                 request.VatCity,
		VatState:                request.VatState,
		VatCountry:              request.VatCountry,
		VatPostcode:             request.VatPostcode,
	}
}

type generatePayLinkRequest struct {
	ProductID               int         `schema:"product_id,omitempty"`
	Title                   string      `schema:"title,omitempty"`
	WebhookURL              string      `schema:"webhook_url,omitempty"`
	Prices                  []string    `schema:"prices,omitempty"`
	RecurringPrices         []string    `schema:"recurring_prices,omitempty"`
	TrialDays               int         `schema:"trial_days,omitempty"`
	CustomMessage           string      `schema:"custom_message,omitempty"`
	CouponCode              string      `schema:"coupon_code,omitempty"`
	Discountable            *customBool `schema:"discountable,omitempty"`
	ImageURL                string      `schema:"image_url,omitempty"`
	ReturnURL               string      `schema:"return_url,omitempty"`
	QuantityVariable        *customBool `schema:"quantity_variable,omitempty"`
	Quantity                int         `schema:"quantity,omitempty"`
	Expires                 customDate  `schema:"expires,omitempty"`
	Affiliates              []string    `schema:"affiliates,omitempty"`
	RecurringAffiliateLimit int         `schema:"recurring_affiliate_limit,omitempty"`
	MarketingConsent        bool        `schema:"marketing_consent,omitempty"`
	CustomerEmail           string      `schema:"customer_email,omitempty"`
	CustomerCountry         string      `schema:"customer_country,omitempty"`
	CustomerPostcode        string      `schema:"customer_postcode,omitempty"`
	IsRecoverable           *customBool `schema:"is_recoverable,omitempty"`
	Passthrough             string      `schema:"passthrough,omitempty"`
	VatNumber               string      `schema:"vat_number,omitempty"`
	VatCompanyName          string      `schema:"vat_company_name,omitempty"`
	VatStreet               string      `schema:"vat_street,omitempty"`
	VatCity                 string      `schema:"vat_city,omitempty"`
	VatState                string      `schema:"vat_state,omitempty"`
	VatCountry              string      `schema:"vat_country,omitempty"`
	VatPostcode             string      `schema:"vat_postcode,omitempty"`
}

type baseAPIResponse struct {
	Success *bool            `json:"success"`
	Err     errorAPIResponse `json:"error"`
}

func (b *baseAPIResponse) IsSuccess() bool {
	if b.Success == nil || *b.Success {
		return true
	}
	return false
}

func (b *baseAPIResponse) Error() string {
	return fmt.Sprintf("paddle error - %d:%s", b.Err.Code, b.Err.Message)
}

type errorAPIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type generatePayLinkResponse struct {
	baseAPIResponse
	Response struct {
		URL string `json:"url"`
	} `json:"response"`
}

func (c *Client) GeneratePayLink(ctx context.Context, request GeneratePayLinkRequest) (url string, err error) {
	body := toGeneratePayLinkRequest(request)
	req, err := c.NewRequest(http.MethodPost, "2.0/product/generate_pay_link", body)
	if err != nil {
		return "", errors.WithStack(err)
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data generatePayLinkResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return "", errors.WithStack(err)
		}
		if data.IsSuccess() {
			return data.Response.URL, nil
		}
		return "", &data
	}
	return "", errors.Errorf("got status code: %d", resp.StatusCode)
}
