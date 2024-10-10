package main

import "fmt"

func printReports(intro, body, outro string) {

	printCostReport(func(s string) int {
		res := len(s) * 2
		return res
	}, intro)

	printCostReport(func(s string) int {
		res := len(s) * 3
		return res
	}, body)

	printCostReport(func(s string) int {
		res := len(s) * 4
		return res
	}, outro)

}

// don't touch below this line

func main5() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
