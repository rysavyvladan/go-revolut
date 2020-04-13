package merchant

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rysavyvladan/go-revolut/merchant/1.0/request"
	"net/http"
)

type OrderService struct {
	apiKey string
}

type OrderType string

const (
	OrderType_PAYMENT    OrderType = "PAYMENT"
	OrderType_REFUND     OrderType = "REFUND"
	OrderType_CHARGEBACK OrderType = "CHARGEBACK"
)

type Funding string

const (
	Funding_CREDIT  Funding = "CREDIT"
	Funding_DEBIT   Funding = "DEBIT"
	Funding_PREPAID Funding = "PREPAID"
)

type TreeDsState string

const (
	TreeDSState_VERIFIED  TreeDsState = "VERIFIED"
	TreeDSState_FAILED    TreeDsState = "FAILED"
	TreeDSState_CHALLENGE TreeDsState = "CHALLENGE"
)

type CardType string

const (
	CardType_VISA       CardType = "VISA"
	CardType_MASTERCARD CardType = "MASTERCARD"
)

type RiskLevel string

const (
	RiskLevel_LOW  RiskLevel = "LOW"
	RiskLevel_HIGH RiskLevel = "HIGH"
)

type CvvVerification string

const (
	CvvVerification_MATCH         CvvVerification = "MATCH"
	CvvVerification_NOT_MATCH     CvvVerification = "NOT_MATCH"
	CvvVerification_INCORRECT     CvvVerification = "INCORRECT"
	CvvVerification_NOT_PROCESSED CvvVerification = "NOT_PROCESSED"
)

type CheckResult string

const (
	CheckResult_MATCH     CheckResult = "MATCH"
	CheckResult_NOT_MATCH CheckResult = "NOT_MATCH"
	CheckResult_N_A       CheckResult = "N_A"
	CheckResult_INVALID   CheckResult = "INVALID"
)

type OrderState string

const (
	OrderState_PENDING    OrderState = "PENDING"
	OrderState_PROCESSING OrderState = "PROCESSING"
	OrderState_AUTHORISED OrderState = "AUTHORISED"
	OrderState_COMPLETED  OrderState = "COMPLETED"
	OrderState_FAILED     OrderState = "FAILED"
)

type Amount struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

type FeeType string

const (
	FeeType_FX        FeeType = "FX"
	FeeType_ACQUIRING FeeType = "ACQUIRING"
)

type Fee struct {
	// Fee amount
	Value int `json:"value"`
	// Fee currency
	Currency string `json:"currency"`
	// Fee type
	Type FeeType `json:"type"`
}

type Payment struct {
	Type          string `json:"type"`
	Amount        Amount `json:"amount"`
	CreatedDate   int64  `json:"created_date"`
	UpdatedDate   int64  `json:"updated_date"`
	CompletedDate int    `json:"completed_date"`
	Card          Card   `json:"card"`
}

type ThreeDs struct {
	// 3DS check result
	State TreeDsState `json:"state"`
	// 3DS version
	Version int `json:"version"`
}

type Check struct {
	// Confirms whether a proxy number is used or not
	Proxy bool `json:"proxy"`
	// Confirms whether a VPN connection is used or not
	Vpn bool `json:"vpn"`
	// Country name associated with the IP address of the card used
	CountryByIp string  `json:"country_by_ip"`
	ThreeDs     ThreeDs `json:"three_ds"`
	// Authorization code returned by the Processor
	AuthorizationCode string `json:"authorization_code"`
	// CVV verification
	CvvVerification CvvVerification `json:"cvv_verification"`
	// Address verification
	Address CheckResult `json:"address"`
	// Postal code verification
	PostalCode CheckResult `json:"postal_code"`
	// Cardholder verification
	CardHolder CheckResult `json:"card_holder"`
}

type Card struct {
	// Card type
	CardType CardType `json:"card_type"`
	// Card funding
	Funding Funding `json:"funding"`
	// Card BIN
	CardBin string `json:"card_bin"`
	// Card last four digits
	CardLastFour string `json:"card_last_four"`
	// Card expiry date in the format of MM/YY
	CardExpiry string `json:"card_expiry"`
	// Cardholder name
	CardholderName string         `json:"cardholder_name"`
	Checks         Check          `json:"checks"`
	RiskLevel      RiskLevel      `json:"risk_level"`
	BillingAddress BillingAddress `json:"billing_address"`
}

type BillingAddress struct {
	// Street line 1 information
	StreetLine1 string `json:"street_line_1"`
	// Street line 2 information
	StreetLine2 string `json:"street_line_2"`
	// Region name
	Region string `json:"region"`
	// City name
	City string `json:"city"`
	// Country associated with the address
	CountryCode string `json:"country_code"`
	// Postcode associated with the address
	Postcode string `json:"postcode"`
}

type OrderResp struct {
	// Order ID for a merchant
	Id string `json:"id"`
	// Temporary ID for a customer
	PublicId string `json:"public_id"`
	// Temporary ID for a customer
	Type OrderType `json:"type"`
	// Order state
	State OrderState `json:"state"`
	// Order creation date, measured in ms since the Unix epoch (UTC)
	CreatedDate int64 `json:"created_date"`
	// Last update date, measured in ms since the Unix epoch (UTC)
	UpdatedDate int64 `json:"updated_date"`
	// Order completion date, measured in ms since the Unix epoch (UTC)
	CompletedDate int64  `json:"completed_date"`
	OrderAmount   Amount `json:"order_amount"`
	// Merchant order ID
	MerchantOrderExtRef string `json:"merchant_order_ext_ref"`
	// Merchant customer ID
	MerchantCustomerExtRef string `json:"merchant_customer_ext_ref"`
	// Customer e-mail
	Email           string           `json:"email"`
	SettledAmount   Amount           `json:"settled_amount"`
	RefundedAmount  Amount           `json:"refunded_amount"`
	Fees            []Fee            `json:"fees"`
	Payments        []Payment        `json:"payments"`
	Attempts        []AttemptRelated `json:"attempts"`
	Related         []AttemptRelated `json:"related"`
	ShippingAddress ShippingAddress  `json:"shipping_address"`
	Phone           string           `json:"phone"`
}

type AttemptRelated struct {
	Id     string    `json:"id"`
	Type   OrderType `json:"type"`
	Amount Amount    `json:"amount"`
}

type ShippingAddress struct {
	// Shipping address: Street line 1 information
	StreetLine1 string `json:"street_line_1"`
	// Shipping address: Street line 2 information
	StreetLine2 string `json:"street_line_2"`
	// Shipping address: Region name
	Region string `json:"region"`
	// Shipping address: City name
	City string `json:"city"`
	// Shipping address: Country associated with the address
	CountryCode string `json:"country_code"`
	// Shipping address: Postcode associated with the address
	Postcode string `json:"postcode"`
}

type CaptureMode string

const (
	CaptureMode_MANUAL    CaptureMode = "MANUAL"
	CaptureMode_AUTOMATIC CaptureMode = "AUTOMATIC"
)

type OrderReq struct {
	// Minor amount
	Amount int `json:"amount"`
	// Capture mode. If it is equal to null then AUTOMATIC is used
	CaptureMode CaptureMode `json:"capture_mode"`
	// Merchant order ID
	MerchantOrderID string `json:"merchant_order_id"`
	// Customer e-mail
	CustomerEmail string `json:"customer_email"`
	// Order description
	Description string `json:"description"`
	// Currency code
	Currency string `json:"currency"`
	// Settlement currency. If it is equal to null then the payment is settled in transaction currency.
	SettlementCurrency string `json:"settlement_currency"`
	// Merchant customer ID
	MerchantCustomerID string `json:"merchant_customer_id"`
}

type RefundReq struct {
	// Minor amount
	Amount int `json:"amount"`
	// Merchant order ID
	MerchantOrderID string `json:"merchant_order_id"`
	// Order description
	Description string `json:"description"`
	// Currency code
	Currency string `json:"currency"`
}

type RefundResp struct {
	// Order ID for a merchant
	Id string `json:"id"`
	// Order type
	Type OrderType `json:"type"`
	// Order state
	State OrderState `json:"state"`
	// Order creation date, measured in ms since the Unix epoch (UTC)
	CreatedDate int64 `json:"created_date"`
	// Last update date, measured in ms since the Unix epoch (UTC)
	UpdatedDate int64 `json:"updated_date"`
	// Order completion date, measured in ms since the Unix epoch (UTC)
	CompletedDate int64  `json:"completed_date"`
	OrderAmount   Amount `json:"order_amount"`
	// Merchant customer ID
	MerchantCustomerExtRef string `json:"merchant_customer_ext_ref"`
	// Customer e-mail
	Email   string           `json:"email"`
	Related []AttemptRelated `json:"related"`
}

// Create:
// doc: https://revolut-engineering.github.io/api-docs/merchant-api/#backend-api-backend-api-order-object-create-payment-order
func (a *OrderService) Create(orderReq *OrderReq) (*OrderResp, error) {
	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodPost,
		Url:         "https://merchant.revolut.com/api/1.0/orders",
		ApiKey:      a.apiKey,
		Body:        orderReq,
		ContentType: request.ContentType_APPLICATION_JSON,
	})
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	var r *OrderResp
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// WithId: If you would like to get information about the created order, please use the following request.
// doc: https://revolut-engineering.github.io/api-docs/merchant-api/#backend-api-backend-api-order-object-retrieve-order
func (a *OrderService) WithId(id string) (*OrderResp, error) {
	resp, statusCode, err := request.New(request.Config{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("https://merchant.revolut.com/api/1.0/orders/%s", id),
		ApiKey: a.apiKey,
	})
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	var r *OrderResp
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// Capture: Once the payment is authorised, the merchant needs to
// capture it in order for it to be sent into the processing stage.
// doc: https://revolut-engineering.github.io/api-docs/merchant-api/#backend-api-backend-api-order-object-capture-order
func (a *OrderService) Capture(id string) (*OrderResp, error) {
	resp, statusCode, err := request.New(request.Config{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("https://merchant.revolut.com/api/1.0/orders/%s/capture", id),
		ApiKey: a.apiKey,
	})
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	var r *OrderResp
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// Cancel: In case the payment has not been captured yet and the merchant decides
// to not proceed with the order, the order can be cancelled manually.
// doc: https://revolut-engineering.github.io/api-docs/merchant-api/#backend-api-backend-api-order-object-cancel-order
func (a *OrderService) Cancel(id string) (*OrderResp, error) {
	resp, statusCode, err := request.New(request.Config{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("https://merchant.revolut.com/api/1.0/orders/%s/cancel", id),
		ApiKey: a.apiKey,
	})
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	var r *OrderResp
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// Refund: In case the customer requires a refund for a payment that has been already captured,
// the merchant can always issue a full or partial refund for a particular payment.
// doc: https://revolut-engineering.github.io/api-docs/merchant-api/#backend-api-backend-api-order-object-refund-order
func (a *OrderService) Refund(id string, refundReq *RefundReq) (*RefundResp, error) {
	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodPost,
		Url:         fmt.Sprintf("https://merchant.revolut.com/api/1.0/orders/%s/refund", id),
		ApiKey:      a.apiKey,
		Body:        refundReq,
		ContentType: request.ContentType_APPLICATION_JSON,
	})
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	var r *RefundResp
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}
