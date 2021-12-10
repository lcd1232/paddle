package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseInvoiceSentWebhook(form url.Values) (InvoiceSent, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return InvoiceSent{}, errors.WithStack(err)
	}
	var isw invoiceSentWebhook
	if err := decoder.Decode(&isw, form); err != nil {
		return InvoiceSent{}, errors.WithStack(err)
	}
	is := InvoiceSent{
		AlertName:                    Alert(isw.AlertName),
		AlertID:                      isw.AlertID,
		PaymentID:                    isw.PaymentID,
		Amount:                       isw.Amount,
		SaleGross:                    isw.SaleGross,
		TermDays:                     isw.TermDays,
		Status:                       Status(isw.Status),
		PurchaseOrderNumber:          isw.PurchaseOrderNumber,
		InvoicedAt:                   time.Time(isw.InvoicedAt),
		Currency:                     isw.Currency,
		ProductID:                    isw.ProductID,
		ProductName:                  isw.ProductName,
		ProductAdditionalInformation: isw.ProductAdditionalInformation,
		CustomerID:                   isw.CustomerID,
		CustomerName:                 isw.CustomerName,
		Email:                        isw.Email,
		CustomerVatNumber:            isw.CustomerVatNumber,
		CustomerCompanyNumber:        isw.CustomerCompanyNumber,
		CustomerAddress:              isw.CustomerAddress,
		CustomerCity:                 isw.CustomerCity,
		CustomerState:                isw.CustomerState,
		CustomerZipcode:              isw.CustomerZipcode,
		Country:                      isw.Country,
		ContractID:                   isw.ContractID,
		ContractStartDate:            time.Time(isw.ContractStartDate),
		ContractEndDate:              time.Time(isw.ContractEndDate),
		Passthrough:                  isw.Passthrough,
		DateCreated:                  time.Time(isw.DateCreated),
		BalanceCurrency:              isw.BalanceCurrency,
		PaymentTax:                   isw.PaymentTax,
		PaymentMethod:                PaymentMethod(isw.PaymentMethod),
		Fee:                          isw.Fee,
		Earnings:                     isw.Earnings,
		EventTime:                    time.Time(isw.EventTime),
	}
	return is, nil
}

type InvoiceSent struct {
	AlertName                    Alert
	AlertID                      string
	PaymentID                    string
	Amount                       string
	SaleGross                    string
	TermDays                     string
	Status                       Status
	PurchaseOrderNumber          string
	InvoicedAt                   time.Time
	Currency                     string
	ProductID                    string
	ProductName                  string
	ProductAdditionalInformation string
	CustomerID                   string
	CustomerName                 string
	Email                        string
	CustomerVatNumber            string
	CustomerCompanyNumber        string
	CustomerAddress              string
	CustomerCity                 string
	CustomerState                string
	CustomerZipcode              string
	Country                      string
	ContractID                   string
	ContractStartDate            time.Time
	ContractEndDate              time.Time
	Passthrough                  string
	DateCreated                  time.Time
	BalanceCurrency              string
	PaymentTax                   string
	PaymentMethod                PaymentMethod
	Fee                          string
	Earnings                     string
	EventTime                    time.Time
}

type invoiceSentWebhook struct {
	AlertName                    string     `schema:"alert_name"`
	AlertID                      string     `schema:"alert_id"`
	PaymentID                    string     `schema:"payment_id"`
	Amount                       string     `schema:"amount"`
	SaleGross                    string     `schema:"sale_gross"`
	TermDays                     string     `schema:"term_days"`
	Status                       string     `schema:"status"`
	PurchaseOrderNumber          string     `schema:"purchase_order_number"`
	InvoicedAt                   customTime `schema:"invoiced_at"`
	Currency                     string     `schema:"currency"`
	ProductID                    string     `schema:"product_id"`
	ProductName                  string     `schema:"product_name"`
	ProductAdditionalInformation string     `schema:"product_additional_information"`
	CustomerID                   string     `schema:"customer_id"`
	CustomerName                 string     `schema:"customer_name"`
	Email                        string     `schema:"email"`
	CustomerVatNumber            string     `schema:"customer_vat_number"`
	CustomerCompanyNumber        string     `schema:"customer_company_number"`
	CustomerAddress              string     `schema:"customer_address"`
	CustomerCity                 string     `schema:"customer_city"`
	CustomerState                string     `schema:"customer_state"`
	CustomerZipcode              string     `schema:"customer_zipcode"`
	Country                      string     `schema:"country"`
	ContractID                   string     `schema:"contract_id"`
	ContractStartDate            customDate `schema:"contract_start_date"`
	ContractEndDate              customDate `schema:"contract_end_date"`
	Passthrough                  string     `schema:"passthrough"`
	DateCreated                  customDate `schema:"date_created"`
	BalanceCurrency              string     `schema:"balance_currency"`
	PaymentTax                   string     `schema:"payment_tax"`
	PaymentMethod                string     `schema:"payment_method"`
	Fee                          string     `schema:"fee"`
	Earnings                     string     `schema:"earnings"`
	PSignature                   string     `schema:"p_signature"`
	EventTime                    customTime `schema:"event_time"`
}
