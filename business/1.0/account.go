package business

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rysavyvladan/go-revolut/business/1.0/request"
	"net/http"
	"time"
)

type AccountService struct {
	accessToken string
	sandbox     bool

	err error
}

type AccountState string

const (
	AccountState_ACTIVE   AccountState = "active"
	AccountState_INACTIVE AccountState = "inactive"
)

type AccountResp struct {
	// the account ID
	Id string `json:"id"`
	// the account name
	Name string `json:"name"`
	// the available balance
	Balance float64 `json:"balance"`
	// the account currency
	Currency string `json:"currency"`
	// the account state, one of active, inactive
	State AccountState `json:"state"`
	// determines if the account is visible to other businesses on Revolut
	Public bool `json:"public"`
	// the instant when the account was created
	UpdatedAt time.Time `json:"updated_at"`
	// the instant when the account was last updated
	CreatedAt time.Time `json:"created_at"`
}

type AccountSchema string

const (
	AccountSchema_CHAPS           AccountSchema = "chaps"
	AccountSchema_BACS            AccountSchema = "bacs"
	AccountSchema_FASTER_PAYMENTS AccountSchema = "faster_payments"
	AccountSchema_SEPA            AccountSchema = "sepa"
	AccountSchema_SWIFT           AccountSchema = "swift"
	AccountSchema_ACH             AccountSchema = "ach"
)

type AccountDetailResp struct {
	// IBAN
	Iban string `json:"iban"`
	// BIC
	Bic string `json:"bic"`
	// the account number
	AccountNo string `json:"account_no"`
	// the sort code
	SortCode string `json:"sort_code"`
	// the routing number
	RoutingNumber string `json:"routing_number"`
	// the beneficiary name
	Beneficiary        string             `json:"beneficiary"`
	BeneficiaryAddress BeneficiaryAddress `json:"beneficiary_address"`
	// the country of the bank
	BankCountry string `json:"bank_country"`
	// determines if this account address is pooled or unique
	Pooled bool `json:"pooled"`
	// the reference of the pooled account
	UniqueReference string `json:"unique_reference"`
	// the list of supported schemes, possible values: chaps, bacs, faster_payments, sepa, swift, ach
	Schemes       []AccountSchema `json:"schemes"`
	EstimatedTime EstimatedTime   `json:"estimated_time"`
}

type AccountUnit string

const (
	AccountUnit_DAYS  AccountUnit = "days"
	AccountUnit_HOURS AccountUnit = "hours"
)

type EstimatedTime struct {
	// the unit of the inbound transfer time estimate, possible values: days, hours
	Unit AccountUnit `json:"unit"`
	// the maximum estimate
	Min int `json:"min"`
	// the minimum estimate
	Max int `json:"max"`
}

type BeneficiaryAddress struct {
	// the address line 1 of the beneficiary
	StreetLine1 string `json:"street_line1"`
	// the address line 2 of the beneficiary
	StreetLine2 string `json:"street_line2"`
	// the region of the beneficiary
	Region string `json:"region"`
	// the city of the beneficiary
	City string `json:"city"`
	// the country of the beneficiary
	Country string `json:"country"`
	// the postal code of the beneficiary
	Postcode string `json:"postcode"`
}

// List: This endpoint retrieves your accounts.
// doc: https://revolut-engineering.github.io/api-docs/#business-api-business-api-accounts-get-accounts
func (a *AccountService) List() ([]*AccountResp, error) {
	if a.err != nil {
		return nil, a.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodGet,
		Url:         "https://b2b.revolut.com/api/1.0/accounts",
		AccessToken: a.accessToken,
		Sandbox:     a.sandbox,
		Body:        nil,
	})
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	var r []*AccountResp
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// WithId: This endpoint retrieves one of your accounts by ID.
// doc: https://revolut-engineering.github.io/api-docs/#business-api-business-api-accounts-get-account
func (a *AccountService) WithId(id string) (*AccountResp, error) {
	if a.err != nil {
		return nil, a.err
	}

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodGet,
		Url:         fmt.Sprintf("https://b2b.revolut.com/api/1.0/accounts/%s", id),
		AccessToken: a.accessToken,
		Sandbox:     a.sandbox,
		Body:        nil,
	})
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := &AccountResp{}
	if err := json.Unmarshal(resp, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DetailWithId: This endpoint retrieves individual account details.
// doc: https://revolut-engineering.github.io/api-docs/business-api/#accounts-get-account-details
func (a *AccountService) DetailWithId(id string) ([]*AccountDetailResp, error) {
	if a.err != nil {
		return nil, a.err
	}
	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodGet,
		Url:         fmt.Sprintf("https://b2b.revolut.com/api/1.0/accounts/%s/bank-details", id),
		AccessToken: a.accessToken,
		Sandbox:     a.sandbox,
		Body:        nil,
	})
	if err != nil {
		return []*AccountDetailResp{}, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := []*AccountDetailResp{}
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return r, nil
}
