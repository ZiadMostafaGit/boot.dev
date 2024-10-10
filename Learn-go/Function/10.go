package main

import "fmt"

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	res := calculateBaseBill(costPerMessage, numMessages) * (1 - calculateDiscountRate(numMessages))
	return res
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.2
	}

	if messagesSent > 1000 {
		return 0.1
	}

	return 0.0
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

func main7() {

	resulte := calculateFinalBill(3.00, 7421)
	fmt.Printf("resulte is %v ", resulte)

}
