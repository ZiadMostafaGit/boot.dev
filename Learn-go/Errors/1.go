package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	lenOfMess, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	lenOfMess2, _error := sendSMS(msgToSpouse)
	if _error != nil {
		return 0, _error
	}
	return lenOfMess + lenOfMess2, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {

	msgCustomer := "hi iam customer"
	msgspouse := "hi iam customers spouseaaa"

	num, err := sendSMSToCouple(msgCustomer, msgspouse)
	if err == nil {

		fmt.Println(num)

	}

}
