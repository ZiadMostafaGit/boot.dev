package main

func getInsuranceAmount(status insuranceStatus) int {
	amount := 0
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 1000
	}
	if status.isDented() {
		return 160
	}
	return amount

}

// Define the insuranceStatus struct
type insuranceStatus struct {
	hasInsuranceStatus bool
	isTotaledStatus    bool
	isDentedStatus     bool
	isBigDentStatus    bool
}

// Method to check if the vehicle has insurance
func (i insuranceStatus) hasInsurance() bool {
	return i.hasInsuranceStatus
}

// Method to check if the vehicle is totaled
func (i insuranceStatus) isTotaled() bool {
	return i.isTotaledStatus
}

// Method to check if the vehicle is dented
func (i insuranceStatus) isDented() bool {
	return i.isDentedStatus
}

// Method to check if the vehicle has a big dent
func (i insuranceStatus) isBigDent() bool {
	return i.isBigDentStatus
}

func main2() {
	// Example of how to create an insuranceStatus and use it
	status := insuranceStatus{
		hasInsuranceStatus: true,
		isTotaledStatus:    false,
		isDentedStatus:     true,
		isBigDentStatus:    false,
	}

	insuranceAmount := getInsuranceAmount(status)
	println(insuranceAmount)
}
