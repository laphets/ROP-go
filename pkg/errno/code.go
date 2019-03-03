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
	ErrLinkInvaild		= &Errno{Code: 10007, Message: "The link is invaild."}

	DBError				= &Errno{Code: 20001, Message: "Error occurred when processing database."}
	ErrToken      		= &Errno{Code: 20002, Message: "Error occurred while signing the JSON web token."}
	ErrMissingHeader	= &Errno{Code: 20003, Message: "The length of the `Authorization` header is zero."}
	ErrTokenInvalid     = &Errno{Code: 20004, Message: "The token was invalid."}
	DuplicateKey		= &Errno{Code: 20005, Message: "Duplicate key for database."}
	ErrSMS				= &Errno{Code: 20006, Message: "Error occurred when sending SMS."}
	ErrTime				= &Errno{Code: 20007, Message: "Time error in your submission."}
	ErrOperation		= &Errno{Code: 20008, Message: "Unsupported operation."}
	ErrEncrypt      	= &Errno{Code: 20009, Message: "Encrypt error."}
	ErrPrase			= &Errno{Code: 20010, Message: "Prase error."}
	ErrFileType			= &Errno{Code: 20011, Message: "File type not match."}
	ErrFileRead			= &Errno{Code: 20012, Message: "Read.io error."}
	ErrOSS				= &Errno{Code: 20013, Message: "OSS File Upload Error."}
	ErrFileSize			= &Errno{Code: 20014, Message: "File size exceeding."}
	ErrUUID				= &Errno{Code: 20015, Message: "UUID generates fail."}
	ErrDecrypt      	= &Errno{Code: 20016, Message: "Decrypt error."}

	// user errors
	ErrUserNotFound 	= &Errno{Code: 20102, Message: "The user was not found."}
	ErrPasswordWrong 	= &Errno{Code: 20103, Message: "Password is wrong."}
	ErrPasswordNotSet 	= &Errno{Code: 20103, Message: "Password is not set."}



	ErrInstanceNotFound 	= &Errno{Code: 20302, Message: "The instance was not found."}
	ErrInstanceNotSelected  = &Errno{Code: 20303, Message: "请先在查看纳新页面选择一个纳新实例."}

	ErrFormCantEdit		= &Errno{Code: 20201, Message: "This form can't be edited."}
	ErrFormNotFound		= &Errno{Code: 20202, Message: "The form was not found."}
	ErrNoProperGroup	= &Errno{Code: 20301, Message: "You have no proper group to join."}
	ErrInstanceNotOpen  = &Errno{Code: 20401, Message: "The instance is not available."}

	// submission errors
	ErrFieldEmpty		= &Errno{Code: 20401, Message: "Some field is required."}
	ErrTypeNotMatch		= &Errno{Code: 20402, Message: "Error type not match."}
	ErrAnsNotMatch		= &Errno{Code: 20404, Message: "Some Ans not match."}
	TooMuchIntent		= &Errno{Code: 20405, Message: "Too much intents."}

	ErrInterviewFull		= &Errno{Code: 20501, Message: "Interview is full"}
	ErrInterviewNotAble		= &Errno{Code: 20502, Message: "Interview is not available"}

	ErrAssociationNotExist	= &Errno{Code: 20601, Message: "Association does not exist"}
)
