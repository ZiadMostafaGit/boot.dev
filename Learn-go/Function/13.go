package main

import "fmt"

func adder() func(int) int {
	var num int
	return func(num1 int) int {
		num += num1
		return num
	}
}

func main() {

	theAdder := adder()
	fmt.Println(theAdder(10))
	fmt.Println(theAdder(10))
	fmt.Println(theAdder(10))
	fmt.Println(theAdder(10))

}
