package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"payment_app/src/models"
	"payment_app/src/utils"
)

// PathFileC save the absolute path of the temporary file to account
const (
	PathFileC = `./src/files/account.json`
)

// CreateAccount adds an account to temporary json files
func CreateAccount(account models.Account) error {

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(PathFileC, accountJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

//GetAccount query a created account stored in a temporary file
func GetAccount() (models.Account, error) {
	var account models.Account

	err := utils.CheckTempFile(PathFileC)
	if err != nil {
		return account, nil
	}

	jsonFile, err := os.Open(PathFileC)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	accountJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		fmt.Println(err)
	}

	return account, nil
}

//UpdateLimitAccount update account limit and return value float
func UpdateLimitAccount(currentLimit float64) {

	err := utils.CheckTempFile(PathFileC)
	if err != nil {
		return
	}

	jsonFile, err := os.Open(PathFileC)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	accountJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	var account models.Account

	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		fmt.Println(err)
	}

	account.Accounts.AvailableLimit = currentLimit

	accountJSON, err = json.Marshal(account)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(PathFileC, accountJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
