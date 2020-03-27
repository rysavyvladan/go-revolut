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
	return &AccountService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
		err:         b.refreshAccessToken(),
	}
}

func (b *Client) Counterparty() *CounterpartyService {
	return &CounterpartyService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
		err:         b.refreshAccessToken(),
	}
}

func (b *Client) Transfer() *TransferService {
	return &TransferService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
		err:         b.refreshAccessToken(),
	}
}

func (b *Client) Payment() *PaymentService {
	return &PaymentService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
		err:         b.refreshAccessToken(),
	}
}

func (b *Client) PaymentDraft() *PaymentDraftService {
	return &PaymentDraftService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
		err:         b.refreshAccessToken(),
	}
}

func (b *Client) Exchange() *ExchangeService {
	return &ExchangeService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
		err:         b.refreshAccessToken(),
	}
}

func (b *Client) Webhook() *WebhookService {
	return &WebhookService{
		accessToken: b.accessToken,
		sandbox:     b.sandbox,
		err:         b.refreshAccessToken(),
	}
}

func (b *Client) refreshAccessToken() error {
	if b.accessTokenExpiration > time.Now().Unix() {
		return nil
	}

	expirationOfAccessToken := time.Now().Unix()
	accessToken, err := b.oa.RefreshAccessToken(b.refreshToken)
	if err != nil {
		return err
	}
	b.accessTokenExpiration = expirationOfAccessToken + int64(accessToken.ExpiresIn)
	b.accessToken = accessToken.AccessToken

	return nil
}
