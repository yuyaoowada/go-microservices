package dberrors

type ConflictError struct {
}

func (e *ConflictError) Error() string {
	return "atttempted to create a record with an existing key"
}
