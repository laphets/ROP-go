package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	NoEnoughAuth		= &Errno{Code: 10003, Message: "You have not enough auth to access this resource."}
	RemoteError			= &Errno{Code: 10004, Message: "Error occurred when requesting remoter server."}


	DBError				= &Errno{Code: 20001, Message: "Error occurred when processing database."}
	ErrToken      		= &Errno{Code: 20002, Message: "Error occurred while signing the JSON web token."}
	ErrMissingHeader	= &Errno{Code: 20003, Message: "The length of the `Authorization` header is zero."}
	ErrTokenInvalid     = &Errno{Code: 20104, Message: "The token was invalid."}


		// user errors
	ErrUserNotFound 	= &Errno{Code: 20102, Message: "The user was not found."}
)
