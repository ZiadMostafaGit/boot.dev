package main

// func (e email) cost() int {
// 	if e.isSubscribed {
// 		return len(e.body) * 2
// 	}
// 	return len(e.body) * 5
// }

// func (e email) format() string {
// 	if e.isSubscribed {
// 		return "'" + e.body + "' | Subscribed"
// 	}
// 	return "'" + e.body + "' | Not Subscribed"
// }

// type expense interface {
// 	cost() int
// }

// type formatter interface {
// 	format() string
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// }

// func showinfo(info interface{}) {

// 	if exp, ok := info.(expense); ok {
// 		fmt.Printf("Cost: %v \n", exp.cost())
// 	}
// 	if form, ok := info.(formatter); ok {
// 		fmt.Println(form.format())
// 	}

// }

// func main() {

// myemail := email{
// 	isSubscribed: true,
// 	body:         "hello iam ziad mostafa elsaid",
// }
// 	showinfo(myemail)

// }

/////////////////////////////////////////

// THERE IS TWO WAYS TO USE ONE STRUCT WITH MORE THAN ONE INTERFACE EATHER USING EMPTY INTERFACE AT THE FUNC WE USE INTERFACES WITH OR USE THE INTERFACES AND CALL THE STRUCT AT THE MAIN MULTIPLETIMES

/////////////////////////////////////////

// func showinfo(exp expense, form formatter) {

// 	fmt.Printf("Cost:%v\n", exp.cost())
// 	fmt.Println(form.format())
// }

// func main() {

// 	myemail := email{
// 		isSubscribed: true,
// 		body:         "ziad mostafa",
// 	}

// 	showinfo(myemail, myemail)

// }
