# go-revolut
go-revolut is a Go client library for the [Revolut API](https://developers.revolut.com/)

### Features
* Business API
    * OAuth
    * Accounts
    * Counterparties
    
### Install
```
    go get github.com/rysavyvladan/go-revolut
```

## Business API
### Usage
for setup business api visit [official documentation](https://developers.revolut.com/docs/#business-api-business-api-authentication-setting-up-access-to-your-business-account) 

#### Create client
```go
    clientId := "pOoEBEmp8CwpBDgf3opC7aPnSe9OaSCC-fvvoti_RJU"
    sandbox := true 

    bC := business.NewClient(clientId, sandbox)
```

#### OAuth
```go
	oa := bC.OAuth("privatekey.pem", issuer)

    // exchange an authorisation code with an access token.
	token, err := oa.ExchangeAuthorisationCode("oa_sand_lQRYFO66KjP1HHMOqj1Bfo_PZLThUi2onuyBaBfWTaE")
	if err != nil {
		panic(err)
	}

    // request new access token
	token, err := oa.RefreshAccessToken("oa_sand_mYSDtsl9SXjEEOy7maxO_ISrAOeqji_Eo30y6GSCRnc")
	if err != nil {
		panic(err)
	}
```

#### Accounts
```go
    a := bC.Account(token.AccessToken)

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

More examples you can find in main.go
