package constants

import (
	"fmt"
)

// HTTPCode simply represents the http request code
type HTTPCode uint16

// APPErrorCode represents the error code for a particular error
type APPErrorCode uint16

// APPHttpStatus represents the complete http status of a request, along with the errors
type APPHttpStatus struct {
	HTTPStatusCode HTTPCode   `json:"httpStatusCode"`
	Success        bool       `json:"success"`
	Errors         []AppError `json:"errors"`
}

//AppError represents the complete error with status, generic message and the developer message
type AppError struct {
	Code             APPErrorCode `json:"code"`
	Message          string       `json:"message"`
	DeveloperMessage string       `json:"developerMessage"`
}

func (e AppError) Error() string { return e.Message }

//AppErrors just contains a List of AppError
type AppErrors struct {
	Errors []AppError
}

func (e AppErrors) Error() string {
	var s string
	for _, appError := range e.Errors {
		s = fmt.Sprintf("%s\n%s", s, appError.Message)
	}
	return s
}

const (
	// ParamsInSufficientErrorCode represents the error code if necessary or expected params are not present in the request
	ParamsInSufficientErrorCode APPErrorCode = 1401
	// ParamsInValidErrorCode represents the error code if invalid params are present in the request
	ParamsInValidErrorCode APPErrorCode = 1402

	// IncorrectDataErrorCode is the error code if invalid params are present in the request
	IncorrectDataErrorCode APPErrorCode = 1403

	// InvalidURLKeyErrorCode is the error code if url contains an invalid key
	InvalidURLKeyErrorCode APPErrorCode = 1404

	//RequestValidationFailedCode is the error code if request validation fails
	RequestValidationFailedCode = 1405

	// InvalidURLKeyErrorCode is the error code if url contains an invalid key
	ResourceErrorCode APPErrorCode = 1501

	// DbErrorCode is the error code if any DB related error occurs
	DbErrorCode APPErrorCode = 1502

	IndexErrorCode APPErrorCode = 1503

	// CacheErrorCode is the error code if any cache related error occurs
	CacheErrorCode APPErrorCode = 1504

	InvalidRequestURI APPErrorCode = 1601

	InconsistantDataStateErrorCode       APPErrorCode = 1407
	NotAuthenticatedErrorCode            APPErrorCode = 1409
	NotAuthorizedErrorCode               APPErrorCode = 1411
	FunctionalityNotImplementedErrorCode APPErrorCode = 1505
	InternalServerErrorCode              APPErrorCode = 1500

	InvalidErrorCode = 2501

	// Rate limiting errors
	RateLimiterInternalError = 3501
	RateLimitExceeded        = 1429
)

const (
	HTTPStatusSuccessCode             HTTPCode = 200
	HTTPStatusBadRequestCode          HTTPCode = 400
	HTTPStatusInternalServerErrorCode HTTPCode = 500
	HTTPFatalErrorCode                HTTPCode = 501
	HTTPStatusNotFound                HTTPCode = 404
	HTTPRateLimitExceeded             HTTPCode = 429
	HTTPForbidden                     HTTPCode = 403
)

var appErrorCodeToHTTPCodeMap = map[APPErrorCode]HTTPCode{

	ResourceErrorCode:              HTTPStatusInternalServerErrorCode,
	DbErrorCode:                    HTTPStatusInternalServerErrorCode,
	IndexErrorCode:                 HTTPStatusInternalServerErrorCode,
	CacheErrorCode:                 HTTPStatusInternalServerErrorCode,
	RateLimiterInternalError:       HTTPStatusInternalServerErrorCode,
	InconsistantDataStateErrorCode: HTTPStatusInternalServerErrorCode,
	InternalServerErrorCode:        HTTPStatusInternalServerErrorCode,

	NotAuthenticatedErrorCode: HTTPForbidden,
	NotAuthorizedErrorCode:    HTTPForbidden,

	ParamsInSufficientErrorCode: HTTPStatusBadRequestCode,
	ParamsInValidErrorCode:      HTTPStatusBadRequestCode,
	IncorrectDataErrorCode:      HTTPStatusBadRequestCode,
	InvalidURLKeyErrorCode:      HTTPStatusBadRequestCode,
	RequestValidationFailedCode: HTTPStatusBadRequestCode,

	InvalidRequestURI: HTTPStatusNotFound,

	FunctionalityNotImplementedErrorCode: HTTPFatalErrorCode,
	InvalidErrorCode:                     HTTPFatalErrorCode,

	RateLimitExceeded: HTTPRateLimitExceeded,
}

func GetAppHTTPError(appErrors AppErrors) *APPHttpStatus {
	httpCode := HTTPStatusSuccessCode

	//Only considering the last app error to generate the http code
	if appErrors.Errors != nil && len(appErrors.Errors) > 0 {
		lastAppError := appErrors.Errors[len(appErrors.Errors)-1]
		v, found := appErrorCodeToHTTPCodeMap[lastAppError.Code]
		if !found {
			httpCode = InvalidErrorCode
		}
		httpCode = v
	}

	return getAppErrStatus(httpCode, appErrors)
}

// getAppErrStatus returns the complete httpStatus containing errors and success/failure status
func getAppErrStatus(status HTTPCode, appErrors AppErrors) *APPHttpStatus {
	var httpStatus = &APPHttpStatus{HTTPStatusCode: status}
	var apiErrors []AppError

	apiErrors = appErrors.Errors
	httpStatus.Errors = apiErrors
	if status == HTTPStatusSuccessCode && len(apiErrors) == 0 {
		httpStatus.Success = true
	} else {
		httpStatus.Success = false
	}

	return httpStatus
}

// UpdateAppHTTPError updates the map with error code to http code
func UpdateAppHTTPError(appErrorCodeMap map[APPErrorCode]HTTPCode) {
	for k, v := range appErrorCodeMap {
		appErrorCodeToHTTPCodeMap[k] = v
	}
}
