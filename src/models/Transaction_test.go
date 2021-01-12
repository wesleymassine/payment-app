package models_test

import (
	"errors"
	"testing"

	. "payment_app/src/models"
)

var (
	transactionsAmount  error = errors.New("Invalid transaction amount. The transaction value must not be zero")
	transactionMerchant error = errors.New("Invalid transaction merchant. The transaction merchant must not be empty")
)

type transactionsTestScenario struct {
	accountID     int
	merchant      string
	amount        float64
	expectedError error
}

var transactionTestScenarios = []transactionsTestScenario{
	{1, "Habbib's", 0, transactionsAmount},
	{1, "", 20, transactionMerchant},
	{0, "Habbib's", 20, nil},
	{1, "Burger King", 20, nil},
	{1, "Burger Bobs", 20, nil},
}

type accountsTestScenario struct {
	accountID      int
	activeCard     bool
	availableLimit float64
	ViolationsCode []int
	expectedOutput []int
}

var accountTestTransactionScenarios = []accountsTestScenario{
	{1, true, 100, []int{AccountAlreadyInitialized}, []int{0}},

	{1, false, 100, []int{CardNotActive}, []int{2}},
	{1, true, 0, []int{InsufficientLimit}, []int{3}},
	{1, true, 100, []int{HighFrequencySmallInterval}, []int{4}},
	{1, true, 100, []int{DoubledTransaction}, []int{5}},
	{0, false, 0, []int{AccountNotInitialized}, []int{1}},
	{1, true, 100, []int{}, []int{}},
}

func TestValidateTransaction(t *testing.T) {
	Transactions := Transactions{}
	Accounts := Accounts{}
	var hasViolation bool

	for _, scenario := range transactionTestScenarios {

		Transactions.AccountID = scenario.accountID
		Transactions.Merchant = scenario.merchant
		Transactions.Amount = scenario.amount

		transaction := Transaction{
			Transactions: Transactions,
		}

		t.Run("TransactionValidate", func(t *testing.T) {
			if err := transaction.TransactionValidate(); err == nil {
				if scenario.expectedError != nil {
					t.Errorf("Expected error %s but got error %v", scenario.expectedError.Error(), nil)
				}

			} else if err.Error() != scenario.expectedError.Error() {
				t.Errorf("Expected error %s but got error %s", scenario.expectedError, err)
			}
		})

		t.Run("TransactionViolations", func(t *testing.T) {
			for _, scenarioAccount := range accountTestTransactionScenarios {

				Accounts.ActiveCard = scenarioAccount.activeCard
				Accounts.AvailableLimit = scenarioAccount.availableLimit
				Accounts.ViolationsCode = scenarioAccount.ViolationsCode

				account := Account{
					Accounts: Accounts,
				}

				if _, hasViolation = transaction.TransactionViolations(account); hasViolation {
					for i, v := range Accounts.ViolationsCode {
						if v != scenarioAccount.expectedOutput[i] {
							t.Errorf("Expected error %d but got error code violations %v", scenarioAccount.expectedOutput, Accounts.ViolationsCode)
						}
					}
				}

				t.Run("ProcessTransactionAccount", func(t *testing.T) {
					if _, err := transaction.ProcessTransactionAccount(account); err == nil {
						if !hasViolation {
							t.Errorf("Expected error %s but got error %s", scenario.expectedError, err)
						}
					}
				})
			}

		})
	}
}
