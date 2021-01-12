package models

import (
	"errors"
	"payment_app/src/utils"
)

// Account struct that will be created for storage {"account": {"": , "": }, violations: []}
type Account struct {
	Accounts   Accounts `json:"account"`
	Violations []string `json:"violations"`
}

// Accounts struct that will be created for storage {"active-card": , "available-limit": }
type Accounts struct {
	ID             int     `json:"-"`
	ActiveCard     bool    `json:"active-card"`
	AvailableLimit float64 `json:"available-limit"`
	ViolationsCode []int   `json:"-"`
}

// Repository violations account and Transactions
var (
	ViolationsAccount = []string{
		"account-already-initialized",
		"account-not-initialized",
		"card-not-active",
		"insufficient-limit",
	}
	ViolationsTransactions = []string{
		"high-frequency-small-interval",
		"doubled-transaction",
	}
)

// PathFileC save the absolute path of the temporary file to account
const (
	PathFileC = `./src/files/account.json`
)

// AccountValidate validates data when creating a account
func (account Account) AccountValidate() error {

	if account.Accounts.AvailableLimit <= 0 {
		return errors.New("Invalid Account limit. The limit cannot be equal to or less than zero")
	}

	return nil
}

//AccountViolations verification if violation in the transaction and returns to account
func (account Account) AccountViolations() ([]string, bool) {
	allViolations := append(ViolationsAccount, ViolationsTransactions...)

	allTransactions, hasViolation := utils.SeeksViolations(allViolations, account.Accounts.ViolationsCode)
	if hasViolation {
		return allTransactions, hasViolation
	}
	return nil, hasViolation
}

//GetAccountViolations return create account violations
func (account Account) GetAccountViolations(code int) string {
	return ViolationsAccount[code]
}
