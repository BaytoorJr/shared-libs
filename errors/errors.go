package errors

import "fmt"

// ErrSystem
const ErrSystem = "CC Marketplace"

type ArgError struct {
	System           string `json:"system"`
	Status           int    `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

// Error to formatted string
func (e *ArgError) Error() string {
	return fmt.Sprintf("%d %s", e.Status, e.DeveloperMessage)
}

// Set Developer message and return
func (e *ArgError) SetDevMessage(developerMessage string) *ArgError {
	e.DeveloperMessage = developerMessage
	return e
}
