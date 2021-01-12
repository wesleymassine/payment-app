package repositories

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"payment_app/src/models"
	"payment_app/src/utils"
)

// PathFileT save the absolute path of the temporary file to transaction
const (
	PathFileT = `./src/files/transaction.json`
)

var (
	lastTransaction = 0
)

// CreateTransaction adds an account to temporary json files
func CreateTransaction(transaction models.Transactions) (int, error) {
	err := utils.CheckTempFile(PathFileT)

	if err != nil {
		_, err := os.Create(PathFileT)
		if err != nil {
			println(err)
		}
	}

	file, err := ioutil.ReadFile(PathFileT)
	if err != nil {
		println(err)
	}

	data := []models.Transaction{}

	json.Unmarshal(file, &data)
	lastTransaction++

	transaction.ID = lastTransaction

	newStruct := &models.Transaction{
		Transactions: transaction,
	}

	data = append(data, *newStruct)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		println(err)
	}

	err = ioutil.WriteFile(PathFileT, dataBytes, 0644)
	if err != nil {
		println(err)
	}

	return transaction.ID, nil
}

//GetTransaction query a created account stored in a temporary file
func GetTransaction(transaction models.Transactions) ([]models.Transaction, error) {

	err := utils.CheckTempFile(PathFileT)
	if err != nil {
		data := []models.Transaction{}
		return data, nil
	}

	file, err := ioutil.ReadFile(PathFileT)
	if err != nil {
		println("Error: There are currently no transactions...")
	}

	data := []models.Transaction{}
	json.Unmarshal(file, &data)

	return data, err
}
