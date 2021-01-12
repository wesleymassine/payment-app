package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"payment_app/src/models"
	"payment_app/src/repositories"
	"payment_app/src/responses"
	"time"
)

// CreateTransaction handles the transaction creation endpoint
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var transaction models.Transaction
	if err = json.Unmarshal(requestBody, &transaction); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	account, err := repositories.GetAccount()
	if err != nil {
		fmt.Println(account, err)
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = transaction.TransactionValidate(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var hasViolation bool
	if account.Accounts.ViolationsCode, hasViolation = transaction.TransactionViolations(account); hasViolation {
		account.Violations, hasViolation = account.AccountViolations()
	}

	transaction.Transactions.Time = time.Now()
	if !hasViolation {
		transaction.Transactions.AccountID = account.Accounts.ID
		if account.Accounts.AvailableLimit, err = transaction.ProcessTransactionAccount(account); err == nil {
			repositories.UpdateLimitAccount(account.Accounts.AvailableLimit)
		}
		transaction.Transactions.ID, err = repositories.CreateTransaction(transaction.Transactions)

		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	responses.JSON(w, http.StatusCreated, account)
}

// GetTransaction handles the account get
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction

	account, err := repositories.GetTransaction(transaction.Transactions)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, account)
}
