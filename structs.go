package paddle

type Status string

const (
	StatusActive   Status = "active"
	StatusTrialing Status = "trialing"
	StatusPastDue  Status = "past_due"
	StatusPaused   Status = "paused"
	StatusDeleted  Status = "deleted"
	StatusOverdue  Status = "overdue"
)

type Alert string

const (
	AlertSubscriptionCreated          Alert = "subscription_created"
	AlertSubscriptionUpdated          Alert = "subscription_updated"
	AlertSubscriptionCancelled        Alert = "subscription_cancelled"
	AlertSubscriptionPaymentSucceeded Alert = "subscription_payment_succeeded"
	AlertSubscriptionPaymentFailed    Alert = "subscription_payment_failed"
	AlertSubscriptionPaymentRefunded  Alert = "subscription_payment_refunded"

	AlertPaymentSucceeded Alert = "payment_succeeded"
	AlertPaymentRefunded  Alert = "payment_refunded"
	AlertLockerProcessed  Alert = "locker_processed"

	AlertPaymentDisputeCreated      Alert = "payment_dispute_created"
	AlertPaymentDisputeClosed       Alert = "payment_dispute_closed"
	AlertHighRiskTransactionCreated Alert = "high_risk_transaction_created"
	AlertHighRiskTransactionUpdated Alert = "high_risk_transaction_updated"

	AlertTransferCreated Alert = "transfer_created"
	AlertTransferPaid    Alert = "transfer_paid"

	AlertNewAudienceMember    Alert = "new_audience_member"
	AlertUpdateAudienceMember Alert = "update_audience_member"

	AlertInvoicePaid    Alert = "invoice_paid"
	AlertInvoiceSent    Alert = "invoice_sent"
	AlertInvoiceOverdue Alert = "invoice_overdue"
)

type PausedReason string

const (
	PausedReasonDelinquent PausedReason = "delinquent"
	PausedReasonVoluntary  PausedReason = "voluntary"
)

// PaymentMethod defines possible payment method type.
type PaymentMethod string

const (
	PaymentMethodCard         PaymentMethod = "card"
	PaymentMethodPayPal       PaymentMethod = "paypal"
	PaymentMethodFree         PaymentMethod = "free"
	PaymentMethodApplePay     PaymentMethod = "apple-pay"
	PaymentMethodWireTransfer PaymentMethod = "wire-transfer"
)

type alertName struct {
	AlertName string `schema:"alert_name"`
}

// RefundType defines possible reason for refund.
type RefundType string

const (
	RefundTypeFull    RefundType = "full"
	RefundTypeVat     RefundType = "vat"
	RefundTypePartial RefundType = "partial"
)
