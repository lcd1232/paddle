package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseInvoicePaidWebhook(form url.Values) (InvoicePaid, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return InvoicePaid{}, errors.WithStack(err)
	}
	var ipw invoicePaidWebhook
	if err := decoder.Decode(&ipw, form); err != nil {
		return InvoicePaid{}, errors.WithStack(err)
	}
	ip := InvoicePaid{
		AlertName:                    Alert(ipw.AlertName),
		AlertID:                      ipw.AlertID,
		PaymentID:                    ipw.PaymentID,
		Amount:                       ipw.Amount,
		SaleGross:                    ipw.SaleGross,
		TermDays:                     ipw.TermDays,
		Status:                       Status(ipw.Status),
		PurchaseOrderNumber:          ipw.PurchaseOrderNumber,
		InvoicedAt:                   ipw.InvoicedAt.Time(),
		Currency:                     ipw.Currency,
		ProductID:                    ipw.ProductID,
		ProductName:                  ipw.ProductName,
		ProductAdditionalInformation: ipw.ProductAdditionalInformation,
		CustomerID:                   ipw.CustomerID,
		CustomerName:                 ipw.CustomerName,
		Email:                        ipw.Email,
		CustomerVatNumber:            ipw.CustomerVatNumber,
		CustomerCompanyNumber:        ipw.CustomerCompanyNumber,
		CustomerAddress:              ipw.CustomerAddress,
		CustomerCity:                 ipw.CustomerCity,
		CustomerState:                ipw.CustomerState,
		CustomerZipcode:              ipw.CustomerZipcode,
		Country:                      ipw.Country,
		ContractID:                   ipw.ContractID,
		ContractStartDate:            ipw.ContractStartDate.Time(),
		ContractEndDate:              ipw.ContractEndDate.Time(),
		Passthrough:                  ipw.Passthrough,
		DateCreated:                  ipw.DateCreated.Time(),
		BalanceCurrency:              ipw.BalanceCurrency,
		PaymentTax:                   ipw.PaymentTax,
		PaymentMethod:                PaymentMethod(ipw.PaymentMethod),
		Fee:                          ipw.Fee,
		Earnings:                     ipw.Earnings,
		BalanceEarnings:              ipw.BalanceEarnings,
		BalanceFee:                   ipw.BalanceFee,
		BalanceTax:                   ipw.BalanceTax,
		BalanceGross:                 ipw.BalanceGross,
		DateReconciled:               ipw.DateReconciled.Time(),
		EventTime:                    ipw.EventTime.Time(),
	}
	return ip, nil
}

type InvoicePaid struct {
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
	BalanceEarnings              string
	BalanceFee                   string
	BalanceTax                   string
	BalanceGross                 string
	DateReconciled               time.Time
	EventTime                    time.Time
}

type invoicePaidWebhook struct {
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
	BalanceEarnings              string     `schema:"balance_earnings"`
	BalanceFee                   string     `schema:"balance_fee"`
	BalanceTax                   string     `schema:"balance_tax"`
	BalanceGross                 string     `schema:"balance_gross"`
	DateReconciled               customDate `schema:"date_reconciled"`
	PSignature                   string     `schema:"p_signature"`
	EventTime                    customTime `schema:"event_time"`
}
