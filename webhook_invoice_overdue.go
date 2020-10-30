package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *Client) ParseInvoiceOverdueWebhook(form url.Values) (InvoiceOverdue, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return InvoiceOverdue{}, errors.WithStack(err)
	}
	var iow invoiceOverdueWebhook
	if err := decoder.Decode(&iow, form); err != nil {
		return InvoiceOverdue{}, errors.WithStack(err)
	}
	io := InvoiceOverdue{
		AlertName:                    Alert(iow.AlertName),
		AlertID:                      iow.AlertID,
		PaymentID:                    iow.PaymentID,
		Amount:                       iow.Amount,
		SaleGross:                    iow.SaleGross,
		TermDays:                     iow.TermDays,
		Status:                       Status(iow.Status),
		PurchaseOrderNumber:          iow.PurchaseOrderNumber,
		InvoicedAt:                   time.Time(iow.InvoicedAt),
		Currency:                     iow.Currency,
		ProductID:                    iow.ProductID,
		ProductName:                  iow.ProductName,
		ProductAdditionalInformation: iow.ProductAdditionalInformation,
		CustomerID:                   iow.CustomerID,
		CustomerName:                 iow.CustomerName,
		Email:                        iow.Email,
		CustomerVatNumber:            iow.CustomerVatNumber,
		CustomerCompanyNumber:        iow.CustomerCompanyNumber,
		CustomerAddress:              iow.CustomerAddress,
		CustomerCity:                 iow.CustomerCity,
		CustomerState:                iow.CustomerState,
		CustomerZipcode:              iow.CustomerZipcode,
		Country:                      iow.Country,
		ContractID:                   iow.ContractID,
		ContractStartDate:            time.Time(iow.ContractStartDate),
		ContractEndDate:              time.Time(iow.ContractEndDate),
		Passthrough:                  iow.Passthrough,
		DateCreated:                  time.Time(iow.DateCreated),
		BalanceCurrency:              iow.BalanceCurrency,
		PaymentTax:                   iow.PaymentTax,
		PaymentMethod:                PaymentMethod(iow.PaymentMethod),
		Fee:                          iow.Fee,
		Earnings:                     iow.Earnings,
		PSignature:                   iow.PSignature,
		EventTime:                    time.Time(iow.EventTime),
	}
	return io, nil
}

type InvoiceOverdue struct {
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
	PSignature                   string
	EventTime                    time.Time
}

type invoiceOverdueWebhook struct {
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
