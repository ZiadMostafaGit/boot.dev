package main

import "fmt"

func bulkSend(numMessages int) float64 {
	fee := 0.00
	totalcost := 0.0
	for i := 0; i < numMessages; i++ {
		totalcost += 1 + fee
		fee += 0.01

	}
	return totalcost

}

func main2() {

	total := bulkSend(6)
	fmt.Printf("%.2f", total)

}
