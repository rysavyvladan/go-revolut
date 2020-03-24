package business

import (
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

type Client struct {
	clientId string
	sandbox  bool
}

func NewClient(clientId string, sandbox bool) *Client {
	return &Client{
		clientId: clientId,
		sandbox:  sandbox,
	}
}

func (b *Client) OAuth(privateKeyFilename, issuer string) *OAuthService {
	privateKeyFile, err := ioutil.ReadFile(privateKeyFilename)
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		panic(err)
	}

	return &OAuthService{
		issuer:     issuer,
		clientId:   b.clientId,
		privateKey: privateKey,
		sandbox:    b.sandbox,
	}
}

func (b *Client) Account(accessToken string) *AccountService {
	return &AccountService{
		accessToken: accessToken,
		sandbox:     b.sandbox,
	}
}

func (b *Client) Counterparty(accessToken string) *CounterpartyService {
	return &CounterpartyService{
		accessToken: accessToken,
		sandbox:     b.sandbox,
	}
}

//func (b *Client) Transfer(accessToken string) *Transfer {
//	return &AccountService{
//		accessToken:   accessToken,
//		sandbox: b.sandbox,
//	}
//}
//
//func (b *Client) Payment(accessToken string) *Payment {
//	return &AccountService{
//		accessToken:   accessToken,
//		sandbox: b.sandbox,
//	}
//}
//
