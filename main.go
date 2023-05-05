package main

// struct representing a ledger transaction
type LedgerTransaction struct {
	VoucherNumber string `json:"voucherNumber"`
	Date          string `json:"date"`
	Amount        int    `json:"amount"`
	AccountNumber string `json:"accountNumber"`
}

// Function that returns true if its a cost account
func isCostAccount(accountNumber string) bool {
	return accountNumber == "1211"
}

// Function that returns true if its a sales account
func isSalesAccount(accountNumber string) bool {
	return accountNumber == "3001"
}

// main function
func main() {
	ledgerTransactions := []LedgerTransaction{}
	//Customer invoice
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"1", "2023-01-05", 1250, "1511"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"1", "2023-01-05", -250, "2611"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"1", "2023-01-05", -1000, "3001"})

	//Customer invoice
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"1", "2023-02-05", 1250, "1511"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"1", "2023-02-05", -250, "2611"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"1", "2023-02-05", -1000, "3001"})

	//Supplier invoice
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"2", "2023-02-05", -125, "2441"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"2", "2023-02-05", 25, "2641"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"2", "2023-02-05", 100, "1211"})

	//Supplier invoice
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"2", "2023-03-05", -125, "2441"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"2", "2023-03-05", 25, "2641"})
	ledgerTransactions = append(ledgerTransactions, LedgerTransaction{"2", "2023-03-05", 100, "1211"})

	sumSales := 0
	sumCosts := 0

	for _, ledgerTransaction := range ledgerTransactions {
		if isSalesAccount(ledgerTransaction.AccountNumber) {
			sumSales += -1 * ledgerTransaction.Amount
		}

		if isCostAccount(ledgerTransaction.AccountNumber) {
			sumCosts += ledgerTransaction.Amount
		}
	}

	println("Sales: ", sumSales)
	println("Costs: ", sumCosts)
}
