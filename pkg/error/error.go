package errorhandler

import (
	"fmt"
	resultType "github.com/fnxr21/item-list/pkg/type"
	"github.com/labstack/echo/v4"
)

// lama
func ErrorHandler(c echo.Context, err error, message string, httpStatus int) error {
	c.Logger().Errorf("Error: %v, Message: %s, HTTP Status: %d", err, message, httpStatus)

	return c.JSON(httpStatus, resultType.ErrorResult{
		Status:  httpStatus,
		Message: message,
	})
}

func HttpErrorResponse(c echo.Context, err error, message string, httpStatus int) error {
	if message == "" {
		message = "success"
	}

	// Return JSON response with structured error details
	return c.JSON(httpStatus, resultType.ErrorResultV2{
		Code:    httpStatus,
		Message: message,
		// Error:   err,
	})
}

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

func NewServiceError(code int, err error) *ServiceError {

	return &ServiceError{
		Code:    code,
		Message: err.Error(),
	}
}

func HandlerValidationError(c echo.Context, code int, err string) error {

	return c.JSON(code, resultType.ErrorResult{
		Status:  code,
		Message: err,
	})

}

type check struct {
}
