package business

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rysavyvladan/go-revolut/business/1.0/request"
	"net/http"
	"net/url"
	"strconv"
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
	Type        PaymentType  `json:"type"`
	RequestId   string       `json:"request_id,omitempty"`
	State       PaymentState `json:"state"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	CompletedAt time.Time    `json:"completed_at,omitempty"`
	Reference   string       `json:"reference,omitempty"`
	Legs        []struct {
		LegId        string `json:"leg_id"`
		AccountId    string `json:"account_id"`
		Counterparty struct {
			Type      string `json:"type"`
			AccountId string `json:"account_id"`
		} `json:"counterparty"`
		Amount      float64 `json:"amount"`
		Currency    string  `json:"currency"`
		Description string  `json:"description"`
		Balance     int     `json:"balance"`
	} `json:"legs"`
	ReasonCode string `json:"reason_code,omitempty"`
	Merchant   struct {
		Name         string `json:"name"`
		City         string `json:"city"`
		CategoryCode string `json:"category_code"`
		Country      string `json:"country"`
	} `json:"merchant,omitempty"`
	Card struct {
		CardNumber string `json:"card_number"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Phone      string `json:"phone"`
	} `json:"card,omitempty"`
	RelatedTransactionId string `json:"related_transaction_id,omitempty"`
}

type TransactionReq struct {
	From         string
	To           string
	Counterparty string
	Count        int32
	Type         PaymentType
}

// CreatePayment: This endpoint creates a new payment. If the payment is for another Revolut account,
// business or personal, the transaction may be processed synchronously.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-create-payment
func (p *PaymentService) CreatePayment(paymentReq *PaymentReq) (*TransactionResp, error) {
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

// GetTransactionById: To retrieve a transaction by ID
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) GetTransactionById(id string) (*TransactionResp, error) {
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

// GetTransactionByRequestId: To retrieve a transaction by request ID
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) GetTransactionByRequestId(requestId string) (*TransactionResp, error) {
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

// CancelPayment: This endpoint allows to cancel a scheduled transaction that was initiated by you, via API.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) CancelPayment(id string) (*TransactionResp, error) {
	if p.err != nil {
		return nil, p.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodDelete,
		Url:         fmt.Sprintf("https://b2b.revolut.com/api/1.0/transaction/%s", id),
		AccessToken: p.accessToken,
		Sandbox:     p.sandbox,
	})
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusNoContent {
		return nil, errors.New(string(resp))
	}

	r := &TransactionResp{}
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetTransactions: This endpoint retrieves historical transactions based on the provided query criteria.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#payments-get-transaction
func (p *PaymentService) GetTransactions(transactionReq *TransactionReq) ([]*TransactionResp, error) {
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
		params.Add("count", strconv.Itoa(int(transactionReq.Count)))
	}
	if transactionReq.Type != "" {
		params.Add("type", string(transactionReq.Type))
	}
	fmt.Println(params.Encode())

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
