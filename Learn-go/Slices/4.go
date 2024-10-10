package main

import "fmt"

func sum(nums ...int) int {

	var theSum int
	for i := 0; i < len(nums); i++ {
		theSum += nums[i]
	}
	return theSum

}

func main4() {

	// nums := sum(1, 2, 3, 3, 3, 4, 5)
	nums := []int{1, 2, 3, 3, 3, 4, 5}
	res := sum(nums...)
	fmt.Println(res)

}
