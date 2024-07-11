package apperrors

import (
	"fmt"
	"net/http"
)

// CustomError represents a custom error type as a string.
type CustomError string

// Error implements the error interface for CustomError.
// It converts the CustomError to a string and returns it.
func (e CustomError) Error() string {
	return string(e)
}

// Custome errors with errormessage
const (
	InvalidId                   = CustomError("Invalid id")
	InternalServerError         = CustomError("Internal server error")
	JSONParsingErrorReq         = CustomError("error in parsing request in json")
	JSONParsingErrorResp        = CustomError("error in parsing response in json")
	OutOfRange                  = CustomError("request value is out of range")
	OrganizationNotFound        = CustomError("organization of given id not found")
	InvalidContactEmail         = CustomError("Contact email is already present")
	InvalidDomainName           = CustomError("Domain name is already present")
	InvalidCoreValueData        = CustomError("Invalid corevalue data")
	TextFieldBlank              = CustomError("Text field cannot be blank")
	DescFieldBlank              = CustomError("Description cannot be blank")
	InvalidParentValue          = CustomError("Invalid parent core value")
	InvalidOrgId                = CustomError("Invalid organisation")
	UniqueCoreValue             = CustomError("Choose a unique coreValue name")
	InvalidAuthToken            = CustomError("Invalid Auth token")
	IntranetValidationFailed    = CustomError("Intranet Validation Failed")
	UserNotFound                = CustomError("User not found")
	InvalidIntranetData         = CustomError("Invalid data recieved from intranet")
	GradeNotFound               = CustomError("Grade not found")
	AppreciationNotFound        = CustomError("appreciation not found")
	RoleUnathorized             = CustomError("Role unauthorized")
	PageParamNotFound           = CustomError("Page parameter not found")
	RepeatedUser                = CustomError("Repeated user")
	BadRequest                  = CustomError("Bad request")
	InternalServer              = CustomError("Internal Server")
	FailedToCreateDriver        = CustomError("failure to create driver obj")
	MigrationFailure            = CustomError("migrate failure")
	SelfAppreciationError       = CustomError("user cannot give appreciation to ourself")
	CannotReportOwnAppreciation = CustomError("You cannot report your own appreciations")
	RepeatedReport              = CustomError("You cannot report an appreciation twice")
)

// ErrKeyNotSet - Returns error object specific to the key value passed in
func ErrKeyNotSet(key string) (err error) {
	return fmt.Errorf("key not set: %s", key)
}

// GetHTTPStatusCode returns status code according to customerror and default returns InternalServer error
func GetHTTPStatusCode(err error) int {
	switch err {
	case InternalServerError, JSONParsingErrorResp:
		return http.StatusInternalServerError
	case OrganizationNotFound, InvalidCoreValueData, InvalidParentValue, InvalidOrgId, GradeNotFound, AppreciationNotFound, PageParamNotFound, InvalidIntranetData:
		return http.StatusNotFound
	case BadRequest, InvalidId, JSONParsingErrorReq, TextFieldBlank, DescFieldBlank, UniqueCoreValue, IntranetValidationFailed, RepeatedUser, SelfAppreciationError, CannotReportOwnAppreciation, RepeatedReport:
		return http.StatusBadRequest
	case InvalidContactEmail, InvalidDomainName:
		return http.StatusConflict
	case InvalidAuthToken, RoleUnathorized:
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}
