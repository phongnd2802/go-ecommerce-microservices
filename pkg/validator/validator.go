package validator

import (
	"fmt"
	"net/mail"
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateInt(value int64, minValue int64, maxValue int64) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("must be a number from %d to %d", minValue, maxValue)
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 255); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}



func ValidatePassword(password string) error {
	return ValidateString(password, 6, 100)
}