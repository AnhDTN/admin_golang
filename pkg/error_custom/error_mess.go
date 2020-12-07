package error_custom

var MessageMap = map[ErrorType]string{
	Success:                    "OK",
	InvalidParams:              "Request parameter error - %s",
	ErrorAuthCheckTokenFail:    "Token authentication failed",
	ErrorAuthCheckTokenTimeout: "Token time out",
	ErrorAuthToken:             "Token build failed",
	ErrorAuth:                  "Token error",
	Error:                      "Error occurred",
	ErrorInternalServer:        "Server error",
	ErrorExistEmail:            "The Email Address entered already exists in the system",
	ErrorBadRequest:            "Request error",
	ErrorInvalidParent:         "Invalid parent node",
	ErrorAllowDeleteWithChild:  "Contains children, cannot be deleted",
	ErrorNotAllowDelete:        "Resources are not allowed to be deleted",
	ErrorInvalidOldPass:        "Old password is incorrect",
	ErrorNotFound:              "Resource does not exist",
	ErrorPasswordRequired:      "Password is required",
	ErrorUserDisabled:          "User is disabled, please contact administrator",
	ErrorNoPermission:          "No access",
	ErrorMethodNotAllow:        "Method is not allowed",
	ErrorTooManyRequest:        "Requests are too frequent",
	ErrorLoginFailed:           "Email or password is invalid",
	ErrorExistRole:             "Role name already exists",
	ErrorNotExistUser:          "Account is invalid",
	ErrorExistUser:             "User is exist",
	ErrorExistRoleUser:         "The role has been given to the user and is not allowed to be deleted",
	ErrorNotExistRole:          "Role user is disabled, please contact administrator",
	ErrorTokenExpired:          "Token is expired",
	ErrorTokenInvalid:          "Token is invalid",
	ErrorTokenMalformed:        "That's not even a token",
	ErrorInsertDocument:        "Insert data to database error",
	ErrorDeleteDocument:        "Delete data from database error",
	ErrorFindOneDocument:       "Find document from database error",
}

func GetMessageError(status int) string {
	mess, ok := MessageMap[ErrorType(status)]
	if ok {
		return mess
	}
	return MessageMap[Error]
}
