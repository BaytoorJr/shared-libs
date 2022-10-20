package errors

// 5xx errors variables
var (
	InternalServerError = &ArgError{ErrSystem, 503, "internal server error", "internal server error"}
	DBConnectError      = &ArgError{ErrSystem, 503, "db connect error", "db connect error"}
	DBReadError         = &ArgError{ErrSystem, 503, "db read error", "db read error"}
	DBWriteError        = &ArgError{ErrSystem, 503, "db write error", "db write error"}
	ENVReadError        = &ArgError{ErrSystem, 503, "env variables reading error", "env variables reading error"}
	FilesystemReadError = &ArgError{ErrSystem, 503, "filesystem reading error", "filesystem reading error"}
	RPCError            = &ArgError{ErrSystem, 503, "RPC error", "RPC error"}
	SerializeError      = &ArgError{ErrSystem, 503, "serialization error", "serialization error"}
)
