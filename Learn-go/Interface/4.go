package main

// func getExpenseReport(e expense) (theTypeOfStruct string, theCost float64) {

// 	if em, ok := e.(email); ok {
// 		theTypeOfStruct = em.toAddress
// 		return theTypeOfStruct, em.cost()
// 	}
// 	if sm, ok := e.(sms); ok {
// 		theTypeOfStruct = sm.toPhoneNumber
// 		return theTypeOfStruct, sm.cost()
// 	}
// 	return "", 0.0
// }

// // don't touch below this line

// type expense interface {
// 	cost() float64
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// 	toAddress    string
// }

// func (e email) cost() float64 {
// 	if !e.isSubscribed {
// 		return float64(len(e.body)) * .05
// 	}
// 	return float64(len(e.body)) * .01
// }

// type sms struct {
// 	isSubscribed  bool
// 	body          string
// 	toPhoneNumber string
// }

// func (s sms) cost() float64 {
// 	if !s.isSubscribed {
// 		return float64(len(s.body)) * .1
// 	}
// 	return float64(len(s.body)) * .03
// }

// type invalid struct{}

// func (i invalid) cost() float64 {
// 	return 0.0
// }
