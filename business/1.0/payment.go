package business

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rysavyvladan/go-revolut/business/1.0/request"
	"net/http"
	"net/url"
	"time"
)

type PaymentService struct {
	accessToken string
	sandbox     bool

	err error
}

type PaymentReq struct {
	// the client provided ID of the transaction (40 characters max)
	RequestId string `json:"request_id"`
	// the ID of the account to pay from
	AccountId string          `json:"account_id"`
	Receiver  PaymentReceiver `json:"receiver"`
	// the transaction amount
	Amount float64 `json:"amount"`
	// the transaction currency
	Currency string `json:"currency"`
	// an optional textual reference shown on the transaction
	Reference string `json:"reference,omitempty"`
	// a future date/time
	// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-schedule-payment
	ScheduleFor string `json:"schedule_for,omitempty"`
}

type PaymentReceiver struct {
	// the ID of the receiving counterparty
	CounterpartyId string `json:"counterparty_id"`
	// the ID of the receiving counterparty's account, provide only for payments to business counterparties,
	//can be own account (only for internal counterparties)
	AccountId string `json:"account_id"`
}

type PaymentState string

const (
	PaymentState_PENDING  PaymentState = "pending"
	PaymentState_COMPLETE PaymentState = "completed"
	PaymentState_DECLINE  PaymentState = "declined"
	PaymentState_FAILED   PaymentState = "failed"
)

type PaymentType string

const (
	PaymentType_ATM             PaymentType = "atm"
	PaymentType_CARD_PAYMENT    PaymentType = "card_payment"
	PaymentType_CARD_REFUND     PaymentType = "card_refund"
	PaymentType_CARD_CHARGEBACK PaymentType = "card_chargeback"
	PaymentType_CARD_CREDIT     PaymentType = "card_credit"
	PaymentType_EXCHANGE        PaymentType = "exchange"
	PaymentType_TRANSFER        PaymentType = "transfer"
	PaymentType_LOAN            PaymentType = "loan"
	PaymentType_FEE             PaymentType = "fee"
	PaymentType_REFUND          PaymentType = "refund"
	PaymentType_TOPUP           PaymentType = "topup"
	PaymentType_TOPUP_RETURN    PaymentType = "topup_return"
	PaymentType_TAX             PaymentType = "tax"
	PaymentType_TAX_REFUND      PaymentType = "tax_refund"
)

type TransactionResp struct {
	// the ID of transaction
	Id string `json:"id"`
	// he transaction type, one of atm, card_payment, card_refund, card_chargeback,
	//card_credit, exchange, transfer, loan, fee, refund, topup, topup_return, tax, tax_refund
	Type PaymentType `json:"type"`
	// the client provided request ID
	RequestId string `json:"request_id,omitempty"`
	// the transction state: pending, completed, declined or failed
	State PaymentState `json:"state"`
	// the instant when the transaction was created
	CreatedAt time.Time `json:"created_at"`
	// the instant when the transaction was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// the instant when the transaction was completed, mandatory for completed state only
	CompletedAt time.Time `json:"completed_at,omitempty"`
	// an optional date when the transaction was scheduled for
	ScheduledFor string `json:"scheduled_for"`
	// a user provided payment reference
	Reference string `json:"reference,omitempty"`
	// the legs of transaction, there'll be 2 legs between your Revolut accounts and 1 leg in other cases
	Legs []TransactionLeg `json:"legs"`
	// reason code for declined or failed transaction state
	ReasonCode string `json:"reason_code,omitempty"`
	// the merchant info (only for card payments)
	Merchant TransactionMerchant `json:"merchant,omitempty"`
	// the card information (only for card payments)
	Card TransactionCard `json:"card,omitempty"`
	// the ID of the original transaction which has been refunded (only for refunds)
	RelatedTransactionId string `json:"related_transaction_id,omitempty"`
}

type TransactionLeg struct {
	// the ID of the leg
	LegId string `json:"leg_id"`
	// the ID of the account the transaction is associated with
	AccountId    string          `json:"account_id"`
	Counterparty LegCounterparty `json:"counterparty"`
	// the transaction amount
	Amount float64 `json:"amount"`
	// the transaction currency
	Currency string `json:"currency"`
	// the billing amount for cross-currency payments
	BillAmount float64 `json:"bill_amount"`
	// the billing currency for cross-currency payments
	BillCurrency string `json:"bill_currency"`
	// the transaction leg purpose
	Description string `json:"description"`
	// a total balance of the account the transaction is associated with (optional)
	Balance float64 `json:"balance,omitempty"`
}

type LegCounterparty struct {
	// the counterparty ID
	Id string `json:"id"`
	// the type of account: self, revolut, external
	Type CounterpartyType `json:"type"`
	// the counterparty account ID
	AccountId string `json:"account_id"`
}

type TransactionMerchant struct {
	// the merchant name
	Name string `json:"name"`
	// the merchant city
	City string `json:"city"`
	// the merchant category code
	CategoryCode string `json:"category_code"`
	// 3-letter ISO bankCountry code
	Country string `json:"country"`
}

type TransactionCard struct {
	// the masked card number
	CardNumber string `json:"card_number"`
	// the cardholder's first name
	FirstName string `json:"first_name"`
	// the cardholder's last name
	LastName string `json:"last_name"`
	// the cardholder's phone number
	Phone string `json:"phone"`
}

type TransactionReq struct {
	// an optional timestamp to query from, filtering on the created_at field
	From string
	// an optional timestamp to query to, filtering on the created_at field. Default is now
	To string
	// an optional counterparty id
	Counterparty string
	// an optional number of records to return (1000 max, default is 100)
	Count int32
	// the transaction type, one of atm, card_payment, card_refund, card_chargeback, card_credit,
	//exchange, transfer, loan, fee, refund, topup, topup_return, tax, tax_refund
	Type PaymentType
}

// Create: This endpoint creates a new payment. If the payment is for another Revolut account,
// business or personal, the transaction may be processed synchronously.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-create-payment
func (p *PaymentService) Create(paymentReq *PaymentReq) (*TransactionResp, error) {
	if p.err != nil {
		return nil, p.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodPost,
		Url:         "https://b2b.revolut.com/api/1.0/pay",
		AccessToken: p.accessToken,
		Sandbox:     p.sandbox,
		Body:        paymentReq,
		ContentType: request.ContentType_APPLICATION_JSON,
	})
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := &TransactionResp{}
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// WithId: To retrieve a transaction by ID
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) WithId(id string) (*TransactionResp, error) {
	if p.err != nil {
		return nil, p.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodGet,
		Url:         fmt.Sprintf("https://b2b.revolut.com/api/1.0/transaction/%s", id),
		AccessToken: p.accessToken,
		Sandbox:     p.sandbox,
	})
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := &TransactionResp{}
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// WithRequestId: To retrieve a transaction by request ID
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) WithRequestId(requestId string) (*TransactionResp, error) {
	if p.err != nil {
		return nil, p.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodGet,
		Url:         fmt.Sprintf("https://b2b.revolut.com/api/1.0/transaction/%s?id_type=request_id", requestId),
		AccessToken: p.accessToken,
		Sandbox:     p.sandbox,
	})
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := &TransactionResp{}
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// Cancel: This endpoint allows to cancel a scheduled transaction that was initiated by you, via API.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) Cancel(id string) error {
	if p.err != nil {
		return p.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodDelete,
		Url:         fmt.Sprintf("https://b2b.revolut.com/api/1.0/transaction/%s", id),
		AccessToken: p.accessToken,
		Sandbox:     p.sandbox,
	})
	if err != nil {
		return err
	}
	if statusCode != http.StatusNoContent {
		return errors.New(string(resp))
	}

	return nil
}

// List: This endpoint retrieves historical transactions based on the provided query criteria.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) List(transactionReq *TransactionReq) ([]*TransactionResp, error) {
	if p.err != nil {
		return nil, p.err
	}

	params := url.Values{}
	if transactionReq.From != "" {
		params.Add("from", transactionReq.From)
	}
	if transactionReq.To != "" {
		params.Add("to", transactionReq.To)
	}
	if transactionReq.Counterparty != "" {
		params.Add("counterparty", transactionReq.Counterparty)
	}
	if transactionReq.Count != 0 {
		params.Add("count", fmt.Sprintf("%d", transactionReq.Count))
	}
	if transactionReq.Type != "" {
		params.Add("type", string(transactionReq.Type))
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodGet,
		Url:         fmt.Sprintf("https://b2b.revolut.com/api/1.0/transactions?%s", params.Encode()),
		AccessToken: p.accessToken,
		Sandbox:     p.sandbox,
	})
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := []*TransactionResp{}
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return r, nil
}
