package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {

	if plan == planPro {
		messagesSlice := messages[:]
		return messagesSlice, nil
	}
	if plan == planFree {

		messagesSlice := messages[:2]

		return messagesSlice, nil

	}
	return nil, errors.New("unsupported plan")
}

func main2() {
	messages := [3]string{"Message 1", "Message 2", "Message 3"}

	proMessages, err := getMessageWithRetriesForPlan("pro", messages)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pro Plan Messages:", proMessages)
	}

	freeMessages, err := getMessageWithRetriesForPlan("free", messages)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Free Plan Messages:", freeMessages)
	}

	unsupportedMessages, err := getMessageWithRetriesForPlan("premium", messages)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Unsupported Plan Messages:", unsupportedMessages)
	}
}
