package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rysavyvladan/go-revolut/business/1.0"
	"io/ioutil"
)

func main() {
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

	bC, err := business.NewClient(clientId, refreshToken, privateKey, issuer, sandbox)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n--- ACCOUNTS ---")

	a := bC.Account()

	fmt.Println("\nList of all accounts:")

	accounts, err := a.GetAccounts()
	if err != nil {
		panic(err)
	}

	for _, account := range accounts {
		fmt.Println(account)
	}

	fmt.Println("\nGet AccountService by id:")

	account, err := a.GetAccount(accounts[0].Id)
	if err != nil {
		panic(err)
	}

	fmt.Println(account)

	fmt.Println("\nGet AccountService detail by id:")
	fmt.Println("todo")
	//da, err := account.GetAccountDetail(accounts[0].Id)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(da)

	fmt.Println("\n--- COUNTERARTIES ---")
	cp := bC.Counterparty()

	fmt.Println("\nCreate revolut counterparty:")
	revolutCounterparty, err := cp.AddRevolutCounterparty(&business.RevolutCounterpartyReq{
		ProfileType: business.CounterpartyProfileType_PERSONAL,
		Name:        "John Smith",
		Phone:       "+4412345678900",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(revolutCounterparty)
	fmt.Println("\nGet CounterpartyService by id:")
	revolutCounterparty, err = cp.GetCounterparty(revolutCounterparty.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(revolutCounterparty)

	fmt.Println("\nDelete CounterpartyService by id:")
	err = cp.DeleteCounterparty(revolutCounterparty.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(revolutCounterparty)

	fmt.Println("\nCreate non-revolut counterparty:")
	nonRevolutCounterparty, err := cp.AddNonRevolutCounterparty(&business.NonRevolutCounterpartyReq{
		CompanyName: "John Smith Co.",
		BankCountry: "GB",
		Currency:    "GBP",
		AccountNo:   "12345678",
		SortCode:    "223344",
		Email:       "john@smith.co",
		Phone:       "+447771234455",
		Address: business.NonRevolutCounterpartyReqAddress{
			StreetLine1: "1 Canada Square",
			StreetLine2: "Canary Wharf",
			Region:      "East End",
			Postcode:    "E115AB",
			City:        "London",
			Country:     "GB",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(nonRevolutCounterparty)
	fmt.Println("\nGet CounterpartyService by id:")
	nonRevolutCounterparty, err = cp.GetCounterparty(nonRevolutCounterparty.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(nonRevolutCounterparty)

	fmt.Println("\nDelete CounterpartyService by id:")
	err = cp.DeleteCounterparty(nonRevolutCounterparty.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(nonRevolutCounterparty)

	fmt.Println("\nList of all counterparties:")
	counterparties, err := cp.GetCounterparties()
	if err != nil {
		panic(err)
	}
	for _, counterparty := range counterparties {
		fmt.Println(counterparty)
	}

	//fmt.Println("\n--- TRANSFERS ---")
	//t := bC.Transfer()
	//transfer, err := t.CreateTransfer(&business.TransferReq{
	//	RequestId:       "e0cbf84637264ee082a848c",
	//	SourceAccountId: "af7b7bec-fa83-4528-84ff-5203d97cdc1c",
	//	TargetAccountId: "aa430e82-be4d-4880-a59b-a568c0f10043",
	//	Amount:          1,
	//	Currency:        "GBP",
	//	Reference:       "Test reference payment",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(transfer)

	fmt.Println("\n--- PAYMENT ---")
	t := bC.Payment()
	transactions, err := t.GetTransactions(&business.TransactionReq{
		//From:         "2017-06-01",
		//To:           "2017-06-10",
		//Counterparty: "",
		//Count:        20,
		//Type:         "",
	})
	if err != nil {
		panic(err)
	}
	for _, transaction := range transactions {
		fmt.Println(transaction)
	}

	fmt.Println("\n--- EXCHANGE ---")
	e := bC.Exchange()
	rate, err := e.GetExchangeRates(&business.ExchangeRateReq{
		From:   "USD",
		To:     "EUR",
		Amount: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rate)

	//exchange, err := e.ExchangeCurrency(&business.ExchangeReq{
	//	From: business.ExchangeAmount{
	//		AccountId: "aa430e82-be4d-4880-a59b-a568c0f10043",
	//		Amount:    2,
	//		Currency:  "GBP",
	//	},
	//	To: business.ExchangeAmount{
	//		AccountId: "fcdfc950-46c8-4279-9765-4985a92e5ac0",
	//		Currency:  "USD",
	//	},
	//	Reference: "Test Exchange",
	//	RequestId: "0",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(exchange)
}
