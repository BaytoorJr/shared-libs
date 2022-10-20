package errors

// 4xx errors variables
var (
	AccessDenied     = &ArgError{System: ErrSystem, Status: 403, Message: "access denied", DeveloperMessage: "access denied"}
	InvalidCharacter = &ArgError{ErrSystem, 400, "incorrect input", "incorrect input"}
	IncorrectRequest = &ArgError{ErrSystem, 400, "incorrect request", "incorrect request"}
	NotFound         = &ArgError{ErrSystem, 404, "not found", "not found"}
)
