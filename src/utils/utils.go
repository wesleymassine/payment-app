package utils

import (
	"os"
	"time"
)

const (
	pathFileC = `./src/files/account.json`
	pathFileT = `./src/files/transaction.json`
)

// HasLimit check if account has limit and returns a Boolean value
func HasLimit(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return false
		}
	}
	return true
}

//SeeksViolations check if there was any violation and return two values a string array and a boolean
func SeeksViolations(slice []string, violation []int) ([]string, bool) {
	AccountViolations := []string{}
	for j := 0; j < len(violation); j++ {
		AccountViolations = append(AccountViolations, slice[violation[j]])
	}
	return AccountViolations, len(violation) > 0
}

//CheckTempFile checks if a temporary file exists and returns a variable of type error or null
func CheckTempFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

// AccountAlreadyInitialized checks if there is an initialized account and returns a Boolean value
func AccountAlreadyInitialized() bool {
	_, err := os.Stat(pathFileC)
	if err != nil {
		return false
	}
	return true
}

// InitializesState initializes state app
func InitializesState() bool {
	c := os.Remove(pathFileC)
	t := os.Remove(pathFileT)
	if t != nil || c != nil {
		return false
	}
	return true
}

//CheckAmount checks if there are equal amounts and returns a Boolean value
func CheckAmount(arr []float64, val float64) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

//CheckMerchant checks if there are equal merchant and returns a Boolean value
func CheckMerchant(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// CheckIntervalTransaction checks if transactions in the expected interval and returns interval diff
func CheckIntervalTransaction(secondInterval time.Duration, start time.Time) bool {
	loc, _ := time.LoadLocation("UTC")
	createdAt := time.Now().In(loc).Add(1 * time.Second)
	expiresAt := time.Now().In(loc).Add(secondInterval * time.Second)
	diffInterval := expiresAt.Sub(createdAt)

	t := time.Now()
	StartProcess := t.Sub(start)

	return diffInterval > StartProcess
}
