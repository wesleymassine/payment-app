package utils_test

import (
	. "payment_app/src/utils"
	"testing"
)

type accountTestScenario struct {
	violations     []string
	expectedOutput []int
	expectedFound  bool
}

var (
	accountAlreadyInitialized  = 0
	accountNotInitialized      = 1
	cardNotActive              = 2
	insufficientLimit          = 3
	highFrequencySmallInterval = 4
	doubledTransaction         = 5
)

var (
	ViolationsCode = []int{
		accountAlreadyInitialized,
		accountNotInitialized,
		cardNotActive,
		insufficientLimit,
		highFrequencySmallInterval,
		doubledTransaction}
)
var (
	ViolationsName = []string{
		"account-already-initialized",
		"account-not-initialized",
		"card-not-active",
		"insufficient-limit",
		"high-frequency-small-interval",
		"doubled-transaction",
	}
)

var accountTestScenarioViolations = []accountTestScenario{
	{[]string{"account-already-initialized"}, []int{accountAlreadyInitialized}, true},
	{[]string{"account-not-initialized"}, []int{accountNotInitialized}, true},
	{[]string{"card-not-active"}, []int{cardNotActive}, true},
	{[]string{"insufficient-limit"}, []int{insufficientLimit}, true},
	{[]string{"high-frequency-small-interval"}, []int{highFrequencySmallInterval}, true},
	{[]string{"doubled-transaction"}, []int{doubledTransaction}, true},
	{[]string{}, []int{}, true},
	//{[]string{"doubled-transaction"}, []int{highFrequencySmallInterval}, false},
}

func TestSeeksViolations(t *testing.T) {
	for _, scenario := range accountTestScenarioViolations {
		if _, found := SeeksViolations(ViolationsName, ViolationsCode); found != scenario.expectedFound {
			t.Errorf("Expected found to be %t but got %t", scenario.expectedFound, found)
		}
	}
}
