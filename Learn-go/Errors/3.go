package main

import (
	"errors"
)

// type devideError struct {
// 	devided float64
// }

// func (de devideError) Error() string {
// 	return fmt.Sprintf("no dividing by 0")
// }

// func divide(x, y float64) (float64, error) {
// 	if y == 0 {
// 		return y, devideError{devided: x}
// 		// ?
// 	}
// 	// if x == 0 {
// 	// 	return x, devideError{devided: y}
// 	// }
// 	return x / y, nil
// }

func divide(x, y float64) (float64, error) {

	err := errors.New("no dividing by 0")

	if y == 0 {
		return 0.0, err

	}
	return x / y, nil
}
