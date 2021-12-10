package paddle

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func (c *WebhookClient) ParseTransferCreatedWebhook(form url.Values) (TransferCreated, error) {
	signature := form.Get(signatureKey)
	if err := c.verifier.Verify(c.publicKey, signature, form); err != nil {
		return TransferCreated{}, errors.WithStack(err)
	}
	var tcw transferCreatedWebhook
	if err := decoder.Decode(&tcw, form); err != nil {
		return TransferCreated{}, errors.WithStack(err)
	}
	tc := TransferCreated{
		AlertName: Alert(tcw.AlertName),
		AlertID:   tcw.AlertID,
		Amount:    tcw.Amount,
		Currency:  tcw.Currency,
		EventTime: tcw.EventTime.Time(),
		PayoutID:  tcw.PayoutID,
		Status:    Status(tcw.Status),
	}
	return tc, nil
}

type TransferCreated struct {
	AlertName Alert
	AlertID   string
	Amount    string
	Currency  string
	EventTime time.Time
	PayoutID  string
	Status    Status
}

type transferCreatedWebhook struct {
	AlertName string     `schema:"alert_name"`
	AlertID   string     `schema:"alert_id"`
	Amount    string     `schema:"amount"`
	Currency  string     `schema:"currency"`
	EventTime customTime `schema:"event_time"`
	PayoutID  string     `schema:"payout_id"`
	Status    string     `schema:"status"`
	Signature string     `schema:"p_signature"`
}
