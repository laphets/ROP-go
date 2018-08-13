package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	NoEnoughAuth		= &Errno{Code: 10003, Message: "You have not enough auth to access this resource."}
	RemoteError			= &Errno{Code: 10004, Message: "Error occurred when requesting remoter server."}
	ErrParam            = &Errno{Code: 10005, Message: "The param has some error."}
	NoCookie			= &Errno{Code: 10006, Message: "There is no cookie in the request."}

	DBError				= &Errno{Code: 20001, Message: "Error occurred when processing database."}
	ErrToken      		= &Errno{Code: 20002, Message: "Error occurred while signing the JSON web token."}
	ErrMissingHeader	= &Errno{Code: 20003, Message: "The length of the `Authorization` header is zero."}
	ErrTokenInvalid     = &Errno{Code: 20004, Message: "The token was invalid."}
	DuplicateKey		= &Errno{Code: 20005, Message: "Duplicate key for database."}
	ErrSMS				= &Errno{Code: 20006, Message: "Error occurred when sending SMS."}


	// user errors
	ErrUserNotFound 	= &Errno{Code: 20102, Message: "The user was not found."}
	ErrInstanceNotFound = &Errno{Code: 20302, Message: "The instance was not found."}
	ErrFormCantEdit		= &Errno{Code: 20201, Message: "This form can't be edited."}
	ErrFormNotFound		= &Errno{Code: 20202, Message: "The form was not found."}

)
