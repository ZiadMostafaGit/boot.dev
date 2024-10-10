package main

import "fmt"

func reformat(massage string, formatter func(string) string) string {
	massage = formatter(massage)
	massage = formatter(massage)
	massage = formatter(massage)
	massage = "TEXTIO: " + massage
	return massage

}

func addPeriod(s string) string {
	return s + "."
}

func main3() {
	var base string = "ziad"
	base = reformat(base, addPeriod)
	fmt.Printf("%v", base)

}
