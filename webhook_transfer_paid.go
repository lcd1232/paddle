package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseTransferPaidWebhook(form url.Values) (TransferPaid, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return TransferPaid{}, errors.WithStack(err)
	}
	var tpw transferPaidWebhook
	if err := decoder.Decode(&tpw, form); err != nil {
		return TransferPaid{}, errors.WithStack(err)
	}
	tp := TransferPaid{
		AlertName: Alert(tpw.AlertName),
		AlertID:   tpw.AlertID,
		Amount:    tpw.Amount,
		Currency:  tpw.Currency,
		EventTime: tpw.EventTime.Time(),
		PayoutID:  tpw.PayoutID,
		Status:    Status(tpw.Status),
	}
	return tp, nil
}

type TransferPaid struct {
	AlertName Alert
	AlertID   string
	Amount    string
	Currency  string
	EventTime time.Time
	PayoutID  string
	Status    Status
}

type transferPaidWebhook struct {
	AlertName string     `schema:"alert_name"`
	AlertID   string     `schema:"alert_id"`
	Amount    string     `schema:"amount"`
	Currency  string     `schema:"currency"`
	EventTime customTime `schema:"event_time"`
	PayoutID  string     `schema:"payout_id"`
	Status    string     `schema:"status"`
	Signature string     `schema:"p_signature"`
}
