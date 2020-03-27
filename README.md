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
* Merchant API
    * Orders
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

### Examples
#### Accounts
##### Get all accounts
```go
	accounts, err := bC.Account().List()
	if err != nil {
		panic(err)
	}

	for _, account := range accounts {
		fmt.Println(account)
	}
```
##### Get Account by Id
```go
	account, err := bC.Account().WithId("8b8be318-e81a-4dee-97b5-35399628814f")
	if err != nil {
		panic(err)
	}

	fmt.Println(account)
```

### Counterparties
#### Get all counterparties
```go
	counterparties, err := bC.Counterparty().List()
	if err != nil {
		panic(err)
	}

	for _, counterparty := range counterparties {
		fmt.Println(counterparty)
	}
```

#### Retrieve counterparty by id
```go
	counterparty, err := bC.Counterparty().WithId("2af1d943-a6ee-4ab0-b8b1-67f7d92aa330")
	if err != nil {
		panic(err)
	}
	fmt.Println(counterparty)
```

#### Delete counterparty
```go
	if err := bC.Counterparty().Delete("2af1d943-a6ee-4ab0-b8b1-67f7d92aa330"); err != nil {
		panic(err)
	}
```

### Transfers
#### Create transfer
```go
	transfer, err := bC.Transfer().Create(&business.TransferReq{
		RequestId:       "e0cbf84637264ee082a848c",
		SourceAccountId: "af7b7bec-fa83-4528-84ff-5203d97cdc1c",
		TargetAccountId: "aa430e82-be4d-4880-a59b-a568c0f10043",
		Amount:          1,
		Currency:        "GBP",
		Reference:       "Test reference payment",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(transfer)
```


### Exchanges
#### Get rates
```go
	rate, err := bC.Exchange().Rate(&business.ExchangeRateReq{
		From:   "USD",
		To:     "EUR",
		Amount: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rate)
```

#### Exchange currency
```go
	exchange, err := bC.Exchange().Exchange(&business.ExchangeReq{
		From: business.ExchangeAmount{
			AccountId: "aa430e82-be4d-4880-a59b-a568c0f10043",
			Amount:    2,
			Currency:  "GBP",
		},
		To: business.ExchangeAmount{
			AccountId: "fcdfc950-46c8-4279-9765-4985a92e5ac0",
			Currency:  "USD",
		},
		Reference: "Test Exchange",
		RequestId: "0",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(exchange)
```
