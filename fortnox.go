package main

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// Function that returns LedgerTransactions from Fortnox
func getLedgerTransactions() []LedgerTransaction {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")
	if err != nil {
		println(err.Error())
	}

	println(resp.String())

	//unmarsal json to LedgerTransactions
	ledgerTransactions := []LedgerTransaction{}
	err = json.Unmarshal([]byte(resp.String()), &ledgerTransactions)
	if err != nil {
		println(err.Error())
	}
}
