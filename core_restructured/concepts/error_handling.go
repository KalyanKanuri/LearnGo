package concepts

import (
	"errors"
	"fmt"
)

// sentinel Errors
var (
	ErrUnauthorized = errors.New("Unauthorized Access")
)

// Custom Errors
type ValidationError struct {
	Field string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Validation Failed %s", v.Field)
}

// returning Errors
func validateData() error {
	return fmt.Errorf("Invalid data")
}

// wrapping errors
func wrapErrorExample() error {
	err := validateData()
	return fmt.Errorf("Error in wrapping %w", err)
}

// errors.Is() usage
func errorsIsExample() error {
	return ErrUnauthorized
}

// errors.As() usage
func errorsAsExample() error {
	return &ValidationError{
		Field: "Email",
	}
}

func ErrorHandlingInGo() {
	err := validateData()
	if err != nil {
		fmt.Printf("Error Processing Data %+v\n", err)
	}

	err = wrapErrorExample()
	if err != nil {
		fmt.Printf("Wrap example %+v\n", err)
	}

	err = errorsIsExample()
	if errors.Is(err, ErrUnauthorized) {
		fmt.Printf("UnAuthorized, please try re login %+v\n", err)
	}

	err = errorsAsExample()
	if validationError, ok := errors.AsType[*ValidationError](err); ok {
		fmt.Printf("Error while validation, %+v\n", validationError.Field)
	}
}
