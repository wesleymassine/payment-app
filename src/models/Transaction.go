package models

import (
	"errors"
	"fmt"
	"payment_app/src/utils"
	"time"
)

// Transaction struct that will be created for storage {"transaction": {"merchant": "", "amount": , "time": ""}}
type Transaction struct {
	Transactions `json:"transaction"`
}

// Transactions struct that will be created for storage {"merchant": "", "amount": , "time": ""}
type Transactions struct {
	ID        int       `json:"id"`
	AccountID int       `json:"-"`
	Merchant  string    `json:"merchant"`
	Amount    float64   `json:"amount"`
	Time      time.Time `json:"time"`
}

// Const codes violations repository
const (
	AccountAlreadyInitialized  = 0
	AccountNotInitialized      = 1
	CardNotActive              = 2
	InsufficientLimit          = 3
	HighFrequencySmallInterval = 4
	DoubledTransaction         = 5
)

var (
	transactionNumbers = 0
	countTransactions  = 0
	merchants          = []string{}
	amounts            = []float64{}
	startTransaction   = time.Now()
)

//Define time high-frequency-small-interval and doubled-transaction in seconds
const (
	secondFrequencyInterval = 180 // 3min
)

// TransactionValidate validates data when creating a transaction
func (transaction Transaction) TransactionValidate() error {

	if transaction.Transactions.Amount <= 0 {
		return errors.New("Invalid transaction amount. The transaction value must not be zero")
	}

	if transaction.Transactions.Merchant == "" {
		return errors.New("Invalid transaction merchant. The transaction merchant must not be empty")
	}

	return nil
}

//TransactionViolations checks if there was a violation at the time of the transaction according to business rules
func (transaction Transaction) TransactionViolations(account Account) ([]int, bool) {
	allViolations := []int{}

	if !utils.AccountAlreadyInitialized() {
		allViolations = append(allViolations, AccountNotInitialized)
		return allViolations, len(allViolations) > 0
	}

	if !account.Accounts.ActiveCard {
		allViolations = append(allViolations, CardNotActive)
		return allViolations, len(allViolations) > 0
	}

	if transaction.Transactions.Amount > account.Accounts.AvailableLimit {
		allViolations = append(allViolations, InsufficientLimit)
	}

	if higtFrequecyTransaction() {
		allViolations = append(allViolations, HighFrequencySmallInterval)
	}

	if doubledTransactionFrequency(transaction) {
		allViolations = append(allViolations, DoubledTransaction)
	}

	return allViolations, len(allViolations) > 0
}

//ProcessTransactionAccount process amount of a transaction by checking limit in the account
func (transaction Transaction) ProcessTransactionAccount(account Account) (float64, error) {
	hasLimit := utils.HasLimit(account.Accounts.ViolationsCode, InsufficientLimit)

	if hasLimit && transaction.Transactions.Amount > 0 {
		transactionNumbers++
		fmt.Println("Total transactions account:", transactionNumbers)
		return (account.Accounts.AvailableLimit - transaction.Transactions.Amount), nil
	}
	return account.Accounts.AvailableLimit, nil
}

func higtFrequecyTransaction() bool {
	countTransactions = transactionNumbers
	isInterval := utils.CheckIntervalTransaction(secondFrequencyInterval, startTransaction)

	if countTransactions >= 3 && isInterval {
		countTransactions = 0
		startTransaction = time.Now()
		return true
	}
	return false
}

func doubledTransactionFrequency(transaction Transaction) bool {
	CheckMerchant := utils.CheckMerchant(merchants, transaction.Transactions.Merchant)
	merchants = append(merchants, transaction.Transactions.Merchant)

	CheckAmount := utils.CheckAmount(amounts, transaction.Transactions.Amount)
	amounts = append(amounts, transaction.Transactions.Amount)

	if (CheckMerchant && CheckAmount) && (utils.CheckIntervalTransaction(secondFrequencyInterval, startTransaction)) {
		startTransaction = time.Now()
		return true
	}
	fmt.Println()
	fmt.Println("Merchants Transactions: ", merchants)
	fmt.Println("Amounts Transactions: ", amounts)
	return false
}
