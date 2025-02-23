package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	containsUpperCase := false
	containsDigit := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			containsUpperCase = true
		}
		if char >= '0' && char <= '9' {
			containsDigit = true
		}
	}

	if containsDigit && containsUpperCase {
		return true
	}
	return false
}
