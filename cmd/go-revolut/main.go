package main

import (
	"fmt"
	merchant "github.com/rysavyvladan/go-revolut/merchant/1.0"
)

func main() {
	//clientId := "pOoEBEmp8CwpBDgf3opC7aPnSe9OaSCC-fvvoti_RJU"
	//issuer := "webhook.site"
	//privateKeyFilename := "privatekey.pem"
	//sandbox := true
	//refreshToken := "oa_sand_mYSDtsl9SXjEEOy7maxO_ISrAOeqji_Eo30y6GSCRnc"
	//
	//privateKeyFile, err := ioutil.ReadFile(privateKeyFilename)
	//if err != nil {
	//	panic(err)
	//}
	//
	//privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	//if err != nil {
	//	panic(err)
	//}
	//
	//bC, err := business.NewClient(clientId, refreshToken, privateKey, issuer, sandbox)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("\n--- ACCOUNTS ---")
	//
	//a := bC.Account()
	//
	//fmt.Println("\nList of all accounts:")
	//
	//accounts, err := a.List()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, account := range accounts {
	//	fmt.Println(account)
	//}
	//
	//fmt.Println("\nGet AccountService by id:")
	//
	//account, err := a.WithId(accounts[0].Id)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(account)
	//
	//fmt.Println("\nGet AccountService detail by id:")
	//fmt.Println("todo")
	////da, err := account.GetAccountDetail(accounts[0].Id)
	////if err != nil {
	////	panic(err)
	////}
	////
	////fmt.Println(da)
	//
	//fmt.Println("\n--- COUNTERARTIES ---")
	//cp := bC.Counterparty()
	//
	//fmt.Println("\nCreate revolut counterparty:")
	//revolutCounterparty, err := cp.AddRevolut(&business.RevolutCounterpartyReq{
	//	ProfileType: business.CounterpartyProfileType_PERSONAL,
	//	Name:        "John Smith",
	//	Phone:       "+4412345678900",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(revolutCounterparty)
	//fmt.Println("\nGet CounterpartyService by id:")
	//revolutCounterparty, err = cp.WithId(revolutCounterparty.Id)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(revolutCounterparty)
	//
	//fmt.Println("\nDelete CounterpartyService by id:")
	//err = cp.Delete(revolutCounterparty.Id)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(revolutCounterparty)
	//
	//fmt.Println("\nCreate non-revolut counterparty:")
	//nonRevolutCounterparty, err := cp.AddNonRevolut(&business.NonRevolutCounterpartyReq{
	//	CompanyName: "John Smith Co.",
	//	BankCountry: "GB",
	//	Currency:    "GBP",
	//	AccountNo:   "12345678",
	//	SortCode:    "223344",
	//	Email:       "john@smith.co",
	//	Phone:       "+447771234455",
	//	Address: business.NonRevolutCounterpartyReqAddress{
	//		StreetLine1: "1 Canada Square",
	//		StreetLine2: "Canary Wharf",
	//		Region:      "East End",
	//		Postcode:    "E115AB",
	//		City:        "London",
	//		Country:     "GB",
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(nonRevolutCounterparty)
	//fmt.Println("\nGet CounterpartyService by id:")
	//nonRevolutCounterparty, err = cp.WithId(nonRevolutCounterparty.Id)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(nonRevolutCounterparty)
	//
	//fmt.Println("\nDelete CounterpartyService by id:")
	//err = cp.Delete(nonRevolutCounterparty.Id)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(nonRevolutCounterparty)
	//
	//fmt.Println("\nList of all counterparties:")
	//counterparties, err := cp.List()
	//if err != nil {
	//	panic(err)
	//}
	//for _, counterparty := range counterparties {
	//	fmt.Println(counterparty)
	//}
	//
	//fmt.Println("\n--- EXCHANGE ---")
	//rate, err := bC.Exchange().Rate(&business.ExchangeRateReq{
	//	From:   "USD",
	//	To:     "EUR",
	//	Amount: 100,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rate)

	apiKey := "Yo1ysB60s8GHEFmrConnG450iK6diAw7cDf-wExzbMdqKi8buyWB5CSEETw6hI_Z"

	order, err := merchant.NewClient(apiKey).Order().Create(&merchant.OrderReq{
		Amount:             200,
		CaptureMode:        merchant.CaptureMode_MANUAL,
		MerchantOrderID:    "00122",
		CustomerEmail:      "rysavy.vladan@gmail.com",
		Description:        "description",
		Currency:           "GBP",
		SettlementCurrency: "USD",
		MerchantCustomerID: "sally01",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(order)

	webhooks, err := merchant.NewClient(apiKey).Webhook().List()
	if err != nil {
		panic(err)
	}
	for _, webhook := range webhooks {
		fmt.Println(webhook)
	}

	// Acme Corporation

	err = merchant.NewClient(apiKey).Webhook().Set(&merchant.WebhookUrl{Url: "https://webhook.site/#!/9d886b86-2880-493f-abf9-2b0a04a77df5"})
	if err != nil {
		panic(err)
	}
	//
	//webhooks, err = mC.Webhook().List()
	//if err != nil {
	//	panic(err)
	//}
	//for _, webhook := range webhooks {
	//	fmt.Println(webhook)
	//}
}
