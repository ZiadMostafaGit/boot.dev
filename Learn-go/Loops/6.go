package main

import "fmt"

func countConnections(groupSize int) int {

	var res int = 0
	for i := 1; i < groupSize; i++ {
		res += i

	}
	return res
}

func main() {
	fmt.Println(countConnections(5))

}
