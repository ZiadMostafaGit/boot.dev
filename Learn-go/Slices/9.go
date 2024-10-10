package main

func isValidPassword(password string) bool {
	hasUppercase := false
	hasnumber := false

	size := len(password)

	if size < 5 {
		return false
	}

	if size > 12 {
		return false
	}

	for _, chr := range password {
		if chr >= 65 && chr <= 90 {
			hasUppercase = true
		}
		if chr >= 48 && chr <= 57 {
			hasnumber = true
		}

	}
	if !hasUppercase {
		return false
	}
	if !hasnumber {
		return false
	}

	return true

}
