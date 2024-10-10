package main

import (
	"fmt"
)

func printPrimes(max int) {

	for i := 2; i < max; i++ {

		if i == 2 {
			fmt.Println(i)
		}

		if i%2 == 0 {
			continue
		}
		isprime := true
		for j := 3; j*j <= i; j += 2 {
			if i%j == 0 {
				isprime = false
				break
			}

		}
		if isprime {

			fmt.Println(i)

		}
	}

}

func main5() {

	printPrimes(22)

}
