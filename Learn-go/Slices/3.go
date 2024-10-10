package main

import "fmt"

func getMessageCosts(messages []string) []float64 {

	messagesCost := make([]float64, len(messages), cap(messages))
	for i := 0; i < len(messagesCost); i++ {

		messagesCost[i] = float64(len(messages[i])) * 0.01

	}
	return messagesCost

}
func main3() {

	myslic := make([]int, 5, 10)
	fmt.Println(len(myslic))
	fmt.Println(cap(myslic))

	slic := []string{"ziad", "mostafa", "elsaid"}

	for name := 0; name < len(slic); name++ {
		fmt.Println(slic[name])
	}

}
