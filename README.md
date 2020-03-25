# go-revolut
go-revolut is a Go client library for the [Revolut API](https://developers.revolut.com/)

### Features
* Business API
    * OAuth
    * Accounts
    * Counterparties
    * Payments
    * Transfers
    * Exchanges
    * Payment Drafts
    * Webhooks
    
### Install
```
    go get github.com/rysavyvladan/go-revolut
```

## Business API
### Usage
for setup business api visit [official documentation](https://developers.revolut.com/docs/#business-api-business-api-authentication-setting-up-access-to-your-business-account) 

#### Create client
Every access token is valid for 40 minutes, after which is automatically refresh.

> For businesses on the freelancer plan: You can do this for 90 days, after which the refresh token will not be valid anymore. You will then need to repeat the API authorisation process, as required by the PSD2 regulations.


```go
	clientId := "pOoEBEmp8CwpBDgf3opC7aPnSe9OaSCC-fvvoti_RJU"
	issuer := "webhook.site"
	privateKeyFilename := "privatekey.pem"
	sandbox := true
	refreshToken := "oa_sand_mYSDtsl9SXjEEOy7maxO_ISrAOeqji_Eo30y6GSCRnc"

	privateKeyFile, err := ioutil.ReadFile(privateKeyFilename)
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		panic(err)
	}

	bC,err := business.NewClient(clientId, refreshToken, privateKey, issuer, sandbox)
	if err != nil {
		panic(err)
	}
```

#### Accounts
```go
    a := bC.Account()
    
    // retrieve all accounts
	accounts, err := a.GetAccounts()
	if err != nil {
		panic(err)
	}

	for _, account := range accounts {
		fmt.Println(account)
	}

    // retrieve account by id
	account, err := a.GetAccount(accounts[0].Id)
	if err != nil {
		panic(err)
	}
```

More examples you can find in [main.go](https://github.com/rysavyvladan/go-revolut/blob/master/cmd/go-revolut/main.go)
