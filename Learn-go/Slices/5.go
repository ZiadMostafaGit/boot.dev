package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {

	var res []float64

	for _, c := range costs {

		for len(res) <= c.day {

			res = append(res, 0.0)
		}
		res[c.day] += c.value

	}
	return res
}

// 1,2,3.1

func main5() {

	arr := []cost{
		{0, 4.0},
		{1, 2.1},
		{5, 2.5},
		{1, 3.1},
	}

	res := getCostsByDay(arr)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

}
