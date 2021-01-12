package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"payment_app/src/models"
	"payment_app/src/repositories"
	"payment_app/src/responses"
	"payment_app/src/utils"
)

// CreateAccount handles the accounts POST endpoint
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var account models.Account
	if err = json.Unmarshal(requestBody, &account); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	account.Violations = make([]string, 0)
	if err = account.AccountValidate(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	accountInitialized := utils.AccountAlreadyInitialized()
	if !accountInitialized {
		err = repositories.CreateAccount(account)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusCreated, account)
		return
	}

	account, err = repositories.GetAccount()
	account.Violations = []string{account.GetAccountViolations(0)}
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, account)
}

// GetAccount handles the account get
func GetAccount(w http.ResponseWriter, r *http.Request) {
	account, err := repositories.GetAccount()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, account)
}
