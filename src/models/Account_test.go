package models_test

import (
	"errors"
	. "payment_app/src/models"
	"testing"
)

var (
	accountAvailableLimit error = errors.New("Invalid Account limit. The limit cannot be equal to or less than zero")
)

type accountTestScenario struct {
	violations     []string
	expectedOutput []string
}

var accountTestScenarioViolations = []accountTestScenario{
	{[]string{"account-already-initialized"}, []string{"account-already-initialized"}},
	{[]string{"account-not-initialized"}, []string{"account-not-initialized"}},
	{[]string{"card-not-active"}, []string{"card-not-active"}},
	{[]string{"insufficient-limit"}, []string{"insufficient-limit"}},
	{[]string{"high-frequency-small-interval"}, []string{"high-frequency-small-interval"}},
	{[]string{"doubled-transaction"}, []string{"doubled-transaction"}},
	{[]string{""}, []string{""}},
	//{[]string{""}, []string{"doubled-transaction"}},
}

type accountsTestScenarios struct {
	accountID      int
	activeCard     bool
	availableLimit float64
	ViolationsCode []int
	expectedOutput []int
	expectedError  error
}

var accountCreateTestScenario = []accountsTestScenarios{
	{1, true, 0, []int{}, []int{}, accountAvailableLimit},
	{1, true, 100, []int{AccountAlreadyInitialized}, []int{0}, nil},
	{1, true, 0, []int{AccountAlreadyInitialized}, []int{0}, nil},
	{1, true, 100, []int{AccountNotInitialized}, []int{1}, nil},
	{1, false, 100, []int{CardNotActive}, []int{2}, nil},
	{1, true, 0, []int{InsufficientLimit}, []int{3}, nil},
	{1, true, 100, []int{HighFrequencySmallInterval}, []int{4}, nil},
	{1, true, 100, []int{DoubledTransaction}, []int{5}, nil},
	//{1, true, 100, []int{DoubledTransaction}, []int{4}, nil},
}

func TestValidateAccount(t *testing.T) {
	Accounts := Accounts{}
	var hasViolation bool

	for _, AccountScenario := range accountCreateTestScenario {

		Accounts.ID = AccountScenario.accountID
		Accounts.ActiveCard = AccountScenario.activeCard
		Accounts.AvailableLimit = AccountScenario.availableLimit
		Accounts.ViolationsCode = AccountScenario.ViolationsCode

		account := Account{
			Accounts: Accounts,
		}

		t.Run("AccountValidateData", func(t *testing.T) {
			if err := account.AccountValidate(); err == nil {
				if AccountScenario.expectedError != nil {
					t.Errorf("Expected error %s but got error %v", AccountScenario.expectedError.Error(), nil)
				}

				for i, v := range Accounts.ViolationsCode {
					//fmt.Println(v, AccountScenario.expectedOutput[i])
					if v != AccountScenario.expectedOutput[i] {
						t.Errorf("Expected error %d but got error %v", AccountScenario.expectedOutput, Accounts.ViolationsCode)
					}
				}
			}
		})

		t.Run("AccountViolations", func(t *testing.T) {
			if _, hasViolation = account.AccountViolations(); hasViolation {
				account := Account{
					Violations: account.Violations,
				}

				for _, scenarioViolations := range accountTestScenarioViolations {
					account.Violations = scenarioViolations.violations

					for i, v := range account.Violations {
						//fmt.Println(v, scenarioViolations.expectedOutput[i])
						if v != scenarioViolations.expectedOutput[i] {
							t.Errorf("Expected error %d but got error %v", AccountScenario.expectedOutput, Accounts.ViolationsCode)
						}

						if accountViolations := account.GetAccountViolations(0); accountViolations == scenarioViolations.expectedOutput[i] {
							//fmt.Println(accountViolations, scenarioViolations.expectedOutput[i])
							if Accounts.ID == 0 {
								t.Errorf("Expected error %s but got error %v", account.GetAccountViolations(0), account.Violations[i])
							}

						}

					}
				}
			}
		})

	}
}
