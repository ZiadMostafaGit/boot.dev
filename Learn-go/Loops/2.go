package main

import "fmt"

func maxMessages(thresh int) int {
	fee := 0
	for i := 0; ; i++ {
		thresh -= (100 + fee)
		fee += 1
		if thresh < 0 {
			return i
		}

	}

}

func main3() {

	thresh := 505
	res := maxMessages(thresh)
	fmt.Println(res)

}
