package main

import "fmt"

func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

func converter(x int) int {

	x = x * 10 / 5
	return x

}

func main4() {

	// x, y, z := conversions(converter, 5, 10, 15)
	x, y, z := conversions(func(i int) int {
		i = i * 10 / 5
		return i
	}, 5, 10, 15)
	fmt.Printf("X=%v Y =%v and z=%v", x, y, z)

}
