package main

/* Add an interface check here for the custom errors we have */

type FileCouldNotBeOpenedError struct {
	message string
}

func NewFileCouldNotBeOpenedError(message string) FileCouldNotBeOpenedError {
	return FileCouldNotBeOpenedError{message}
}

func (f FileCouldNotBeOpenedError) Error() string {
	return f.message
}

/* Create a new type FileWasEmptyError that returns a useful error message */
