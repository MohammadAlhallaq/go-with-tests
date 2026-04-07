package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("field '%s': %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	if len(email) == 0 {
		return ValidationError{
			Field:   "email",
			Message: "can't be empty",
		}
	}
	return nil
}

func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("divisor cant be 0")
	}
	return num1 / num2, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result) // 5
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Error: cannot divide by zero
	}

	// Test validateEmail
	err = validateEmail("")
	if err != nil {
		fmt.Println(err)
	}

	var ve ValidationError
	if errors.As(err, &ve) {
		fmt.Println("Bad field:", ve.Field) // Bad field: email
	}
}
