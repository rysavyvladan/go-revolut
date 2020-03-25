package business

import (
	"crypto/rsa"
	"time"
)

type Client struct {
	clientId     string
	sandbox      bool
	privateKey   *rsa.PrivateKey
	issuer       string
	refreshToken string

	accessToken           string
	accessTokenExpiration int64
	oa                    *OAuthService
}

func NewClient(clientId, refreshToken string, privateKey *rsa.PrivateKey, issuer string, sandbox bool) (*Client, error) {
	oa := &OAuthService{
		clientId:   clientId,
		privateKey: privateKey,
		issuer:     issuer,
		sandbox:    sandbox}

	accessTokenExpiration := time.Now().Unix()
	accessToken, err := oa.RefreshAccessToken(refreshToken)
	if err != nil {
		return nil, err
	}
	accessTokenExpiration += int64(accessToken.ExpiresIn)

	return &Client{
		clientId:     clientId,
		sandbox:      sandbox,
		privateKey:   privateKey,
		issuer:       issuer,
		refreshToken: refreshToken,

		accessToken:           accessToken.AccessToken,
		accessTokenExpiration: accessTokenExpiration,
		oa:                    oa,
	}, nil
}

func (b *Client) Account() *AccountService {
	if b.accessTokenExpiration < time.Now().Unix() {
		expirationOfAccessToken := time.Now().Unix()
		accessToken, err := b.oa.RefreshAccessToken(b.refreshToken)
		if err != nil {
			return &AccountService{
				err: err,
			}
		}
		b.accessTokenExpiration = expirationOfAccessToken + int64(accessToken.ExpiresIn)
	}

	return &AccountService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
	}
}

func (b *Client) Counterparty() *CounterpartyService {
	if b.accessTokenExpiration < time.Now().Unix() {
		expirationOfAccessToken := time.Now().Unix()
		accessToken, err := b.oa.RefreshAccessToken(b.refreshToken)
		if err != nil {
			return &CounterpartyService{
				err: err,
			}
		}
		b.accessTokenExpiration = expirationOfAccessToken + int64(accessToken.ExpiresIn)
	}

	return &CounterpartyService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
	}
}

func (b *Client) Transfer() *TransferService {
	if b.accessTokenExpiration < time.Now().Unix() {
		expirationOfAccessToken := time.Now().Unix()
		accessToken, err := b.oa.RefreshAccessToken(b.refreshToken)
		if err != nil {
			return &TransferService{
				err: err,
			}
		}
		b.accessTokenExpiration = expirationOfAccessToken + int64(accessToken.ExpiresIn)
	}

	return &TransferService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
	}
}

func (b *Client) Payment() *PaymentService {
	if b.accessTokenExpiration < time.Now().Unix() {
		expirationOfAccessToken := time.Now().Unix()
		accessToken, err := b.oa.RefreshAccessToken(b.refreshToken)
		if err != nil {
			return &PaymentService{
				err: err,
			}
		}
		b.accessTokenExpiration = expirationOfAccessToken + int64(accessToken.ExpiresIn)
	}

	return &PaymentService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
	}
}
