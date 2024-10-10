package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	messageContent := msg.getMessage()
	messagelen := len(messageContent) * 3
	return messageContent, messagelen

}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main2() {

	sReport := sendingReport{
		reportName:    "ziad mostafa",
		numberOfSends: 200,
	}
	birthday := birthdayMessage{
		recipientName: "John",
		birthdayTime:  time.Date(1990, time.September, 15, 0, 0, 0, 0, time.UTC), // Example birthday
	}
	str1, int1 := sendMessage(sReport)
	str2, int2 := sendMessage(birthday)
	fmt.Println(str1)
	fmt.Println(int1)
	fmt.Println(str2)
	fmt.Println(int2)

}
