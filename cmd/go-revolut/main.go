package main

import (
	"fmt"
	"github.com/rysavyvladan/go-revolut/business/1.0"
)

func main() {
	clientId := "pOoEBEmp8CwpBDgf3opC7aPnSe9OaSCC-fvvoti_RJU"
	issuer := "webhook.site"

	bC := business.NewClient(clientId, true)

	fmt.Println("--- CLIENTS ---")

	oa := bC.OAuth("privatekey.pem", issuer)

	//token, err := oauth.ExchangeAuthorisationCode("oa_sand_lQRYFO66KjP1HHMOqj1Bfo_PZLThUi2onuyBaBfWTaE")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(token)

	fmt.Println("\nRefresh access token:")

	token, err := oa.RefreshAccessToken("oa_sand_mYSDtsl9SXjEEOy7maxO_ISrAOeqji_Eo30y6GSCRnc")
	if err != nil {
		panic(err)
	}

	fmt.Println(token.AccessToken)

	fmt.Println("\n--- ACCOUNTS ---")

	a := bC.Account(token.AccessToken)

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
	cp := bC.Counterparty(token.AccessToken)

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
}
