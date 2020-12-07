package error_custom

import "fmt"

const (
	//Error Common
	Success                   ErrorType = 200
	Error                     ErrorType = 500
	InvalidParams             ErrorType = 400
	ErrorBadRequest           ErrorType = 421
	ErrorNoPermission         ErrorType = 403
	ErrorNotFound             ErrorType = 404
	ErrorMethodNotAllow       ErrorType = 405
	ErrorInvalidParent        ErrorType = 409
	ErrorAllowDeleteWithChild ErrorType = 410
	ErrorNotAllowDelete       ErrorType = 411
	ErrorUserDisabled         ErrorType = 412

	ErrorLoginFailed           ErrorType = 422
	ErrorInvalidOldPass        ErrorType = 423
	ErrorPasswordRequired      ErrorType = 424
	ErrorTooManyRequest        ErrorType = 429
	ErrorInternalServer        ErrorType = 512
	ErrorAuthCheckTokenFail    ErrorType = 401
	ErrorAuthCheckTokenTimeout ErrorType = 402
	ErrorAuthToken             ErrorType = 408
	ErrorAuth                  ErrorType = 407
	ErrorTokenExpired          ErrorType = 461
	ErrorTokenInvalid          ErrorType = 462
	ErrorTokenMalformed        ErrorType = 463

	//Error Exist
	ErrorExistRole     ErrorType = 414
	ErrorExistRoleUser ErrorType = 415
	ErrorNotExistUser  ErrorType = 416
	ErrorExistUser     ErrorType = 417
	ErrorExistEmail    ErrorType = 430
	ErrorNotExistRole  ErrorType = 431

	//Error working with database
	ErrorInsertDocument  ErrorType = 1001
	ErrorDeleteDocument  ErrorType = 1002
	ErrorFindOneDocument ErrorType = 1003
)

type ErrorType int

func (errType ErrorType) New() error {
	return CustomError{
		errorType:    errType,
		wrappedError: fmt.Errorf(GetMessageError(int(errType))),
	}
}
