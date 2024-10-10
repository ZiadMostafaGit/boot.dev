package main

func getMonthlyPrice(tier string) int {
	var priceInPennies int = 0
	switch tier {
	case "basic":
		priceInPennies = 100 * 100
	case "premium":
		priceInPennies = 150 * 100
	case "enterprise":
		priceInPennies = 500 * 100
	default:
		priceInPennies = 0

	}
	return priceInPennies
}
