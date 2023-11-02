package helpers

import (
	"strings"
	"unicode"
)

func ValidEmail(email string) bool {
	// Simple email validation example
	// More complex validation may be required
	// In this example, we're checking if the email is in a basic format
	// Example: user@example.com
	// For more complex validation, you can use regex or email validation libraries
	if len(email) < 3 || len(email) > 320 {
		return false
	}

	atSign := strings.LastIndex(email, "@")
	if atSign == -1 {
		return false
	}

	// You can perform length checks for parts before and after "@" sign
	usernamePart := email[:atSign]
	domainPart := email[atSign+1:]

	if len(usernamePart) < 1 || len(domainPart) < 3 {
		return false
	}

	return true
}

func StrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSymbol := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
		} else if unicode.IsLower(char) {
			hasLowercase = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else {
			// Custom logic for special character checks can be added here
			hasSymbol = true
		}
	}

	// Password strength may require the following conditions to be met
	// 1. Must have at least one uppercase letter (hasUppercase)
	// 2. Must have at least one lowercase letter (hasLowercase)
	// 3. Must have at least one digit (hasDigit)
	// 4. Must have at least one special character (hasSymbol)
	// 5. Must have a minimum length (e.g., at least 8 characters)

	return hasUppercase && hasLowercase && hasDigit && hasSymbol
}
