package main

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

type car struct {
	make  string
	model string
}

type truck struct {
	bedsize int
	car
}

// func main() {

// 	ziad := sender{
// 		rateLimit: 1000,
// 		user: user{
// 			name:   "ziad",
// 			number: 12212121,
// 		},
// 	}

// 	theTruck := truck{
// 		bedsize: 10,
// 		car: car{
// 			make:  "toyota",
// 			model: "camry",
// 		},
// 	}

// 	fmt.Println(theTruck.bedsize)
// 	fmt.Println(theTruck.make)
// 	fmt.Println(theTruck.model)
// 	fmt.Println(ziad.rateLimit)
// 	fmt.Println(ziad.name)
// 	fmt.Println(ziad.number)

// }
