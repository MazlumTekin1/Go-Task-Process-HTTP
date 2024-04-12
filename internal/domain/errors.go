package domain

import "fmt"

type DatabaseError struct {
	Operation string
	Err       error
}

type UserIdIsRequiredError struct {
	Operation string
	Message   string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s: %v", e.Operation, e.Err)
}

func (e *UserIdIsRequiredError) Error() string {
	return fmt.Sprintf("user id is required during %s: %s", e.Operation, e.Message)
}

func NewDatabaseError(operation string, err error) error {
	return &DatabaseError{
		Operation: operation,
		Err:       err,
	}
}

func NewUserIdIsRequiredError(operation string, message string) error {
	return &UserIdIsRequiredError{
		Operation: operation,
		Message:   message,
	}
}

func NewIdLessThanZeroError(operation string, message string) error {
	return &UserIdIsRequiredError{
		Operation: operation,
		Message:   message,
	}
}

func NewTaskFieldIsRequiredError(operation string, message string) error {
	return &UserIdIsRequiredError{
		Operation: operation,
		Message:   message,
	}
}
