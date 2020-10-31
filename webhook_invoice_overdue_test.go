package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInvoiceOverdueWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    InvoiceOverdue
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":                     {"invoice_overdue"},
					"alert_id":                       {"27120763"},
					"payment_id":                     {"1234"},
					"amount":                         {"9.99"},
					"sale_gross":                     {"0.10"},
					"term_days":                      {"1"},
					"status":                         {"overdue"},
					"purchase_order_number":          {"1234"},
					"invoiced_at":                    {"2020-10-29 00:00:00"},
					"currency":                       {"USD"},
					"product_id":                     {"1234"},
					"product_name":                   {"name"},
					"product_additional_information": {"info"},
					"customer_id":                    {"1234"},
					"customer_name":                  {"john doe"},
					"email":                          {"test@example.org"},
					"customer_vat_number":            {"1234"},
					"customer_company_number":        {"1234"},
					"customer_address":               {"address"},
					"customer_city":                  {"city"},
					"customer_state":                 {"state"},
					"customer_zipcode":               {"1234"},
					"country":                        {"US"},
					"contract_id":                    {"1234"},
					"contract_start_date":            {"2020-01-01"},
					"contract_end_date":              {"2021-01-01"},
					"passthrough":                    {"some data"},
					"date_created":                   {"2020-10-28"},
					"balance_currency":               {"USD"},
					"payment_tax":                    {"1234"},
					"payment_method":                 {"card"},
					"fee":                            {"1234"},
					"earnings":                       {"1234"},
					"p_signature":                    {"signature"},
					"event_time":                     {"2020-10-30 00:00:00"},
				},
			},
			want: InvoiceOverdue{
				AlertName:                    AlertInvoiceOverdue,
				AlertID:                      "27120763",
				PaymentID:                    "1234",
				Amount:                       "9.99",
				SaleGross:                    "0.10",
				TermDays:                     "1",
				Status:                       StatusOverdue,
				PurchaseOrderNumber:          "1234",
				InvoicedAt:                   time.Date(2020, 10, 29, 0, 0, 0, 0, time.UTC),
				Currency:                     "USD",
				ProductID:                    "1234",
				ProductName:                  "name",
				ProductAdditionalInformation: "info",
				CustomerID:                   "1234",
				CustomerName:                 "john doe",
				Email:                        "test@example.org",
				CustomerVatNumber:            "1234",
				CustomerCompanyNumber:        "1234",
				CustomerAddress:              "address",
				CustomerCity:                 "city",
				CustomerState:                "state",
				CustomerZipcode:              "1234",
				Country:                      "US",
				ContractID:                   "1234",
				ContractStartDate:            time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				ContractEndDate:              time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Passthrough:                  "some data",
				DateCreated:                  time.Date(2020, 10, 28, 0, 0, 0, 0, time.UTC),
				BalanceCurrency:              "USD",
				PaymentTax:                   "1234",
				PaymentMethod:                PaymentMethodCard,
				Fee:                          "1234",
				Earnings:                     "1234",
				EventTime:                    time.Date(2020, 10, 30, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, vm := NewTestClient()
			vm.On("Verify",
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(
				nil,
			).Once()
			got, err := c.ParseInvoiceOverdueWebhook(tt.args.form)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
			vm.AssertExpectations(t)
		})
	}
}
