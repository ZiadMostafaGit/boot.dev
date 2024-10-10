package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {

	stringArray := [3]string{primary, secondary, tertiary}
	var intArray [3]int

	intArray[0] = len(stringArray[0])

	for i := 1; i < len(intArray); i++ {
		intArray[i] = len(stringArray[i]) + intArray[i-1]
	}

	return stringArray, intArray

}
