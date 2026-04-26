package main

import (
	"errors"
	"fmt"
)

type Aoba struct {
	Employee  string
	Something string
}

// You had:
// type ValidationError struct { Missing string; Message string }
// That became a more specific error type because the exercise wants an
// "empty field" error that carries the field name directly.
type EmptyFieldError struct {
	Field string
}

func (e EmptyFieldError) Error() string {
	return fmt.Sprintf("%s is empty", e.Field)
}

// You had:
// var ErrInvalidId = errors.New("Invalid Id")
// Same sentinel idea, just named and formatted conventionally.
var ErrInvalidID = errors.New("invalid ID")

func DoSomething(mode string) (string, error) {
	if mode == "error" {
		return "", ErrInvalidID
	}

	return "ok", nil
}

// You had:
// func Validate(aoba Aoba) (bool, error)
// Returning only error is enough here: nil means valid, non-nil means invalid.
func Validate(aoba Aoba) error {
	var errs []error

	if aoba.Employee == "" {
		errs = append(errs, EmptyFieldError{Field: "Employee"})
	}

	if aoba.Something == "" {
		errs = append(errs, EmptyFieldError{Field: "Something"})
	}

	// You had:
	// return false, ValidationError{...}
	// That returned the first error found. The exercise wants one error value
	// that still contains every validation problem.
	return errors.Join(errs...)
}

func main() {
	// You had:
	// if errors.Is(ErrInvalidId, err)
	// Same logic; errors.Is just needs the returned error first so it also
	// works if the sentinel is wrapped later.
	something, err := DoSomething("error")
	if errors.Is(err, ErrInvalidID) {
		fmt.Println("found invalid ID error")
	}
	fmt.Println(something)

	otherSomething, err := DoSomething("good")
	if errors.Is(err, ErrInvalidID) {
		fmt.Println("found invalid ID error and should not")
	}
	fmt.Println(otherSomething)

	err = Validate(Aoba{})
	if err != nil {
		fmt.Println("validation errors:")

		// You had:
		// var validationError ValidationError
		// if errors.As(err, &validationError) { ... }
		// errors.Join creates an error with Unwrap() []error, so iterate over the
		// contained errors instead of reporting only the first one.
		type unwrapper interface {
			Unwrap() []error
		}

		var joined unwrapper
		if errors.As(err, &joined) {
			for _, validationErr := range joined.Unwrap() {
				var emptyFieldErr EmptyFieldError
				if errors.As(validationErr, &emptyFieldErr) {
					fmt.Printf("- %s field is empty\n", emptyFieldErr.Field)
					continue
				}
				fmt.Printf("- %v\n", validationErr)
			}
		}
	}
}
