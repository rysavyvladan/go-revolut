package business

import (
	"encoding/json"
	"errors"
	"github.com/rysavyvladan/go-revolut/business/1.0/request"
	"net/http"
	"time"
)

type TransferService struct {
	accessToken string
	sandbox     bool

	err error
}

type TransferReq struct {
	// a unique value used to handle duplicates submitted as a
	// result of lost connection or another client error (40 characters max)
	RequestId string `json:"request_id"`
	// the ID of a source account
	SourceAccountId string `json:"source_account_id"`
	// the ID of a target account
	TargetAccountId string `json:"target_account_id"`
	// the transaction amount
	Amount float64 `json:"amount"`
	// the transaction currency, both source and target accounts should be in this currency
	Currency string `json:"currency"`
	// an optional textual reference shown on the transaction
	Reference string `json:"reference,omitempty"`
}

type TransferState string

const (
	TransferState_PENDING  TransferState = "pending"
	TransferState_COMPLETE TransferState = "completed"
	TransferState_DECLINE  TransferState = "declined"
	TransferState_FAILED   TransferState = "failed"
)

type TransferResp struct {
	// the ID of the created transaction
	Id string `json:"id"`
	// the transction state: pending, completed, declined or failed
	State string `json:"state"`
	// the instant when the transaction was created
	CreatedAt time.Time `json:"created_at"`
	// the instant when the transaction was completed
	CompletedAt time.Time `json:"completed_at"`
}

// CreateTransfer: This endpoint processes transfers between accounts of the business with the same currency.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#transfers-create-transfer
func (t *TransferService) CreateTransfer(transferReq *TransferReq) (*TransferResp, error) {
	if t.err != nil {
		return nil, t.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodPost,
		Url:         "https://b2b.revolut.com/api/1.0/transfer",
		AccessToken: t.accessToken,
		Sandbox:     t.sandbox,
		Body:        transferReq,
		ContentType: request.ContentType_APPLICATION_JSON,
	})
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := &TransferResp{}
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}
