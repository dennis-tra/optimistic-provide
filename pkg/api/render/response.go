package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is the structured type for every API response. The type contains two fields, 'result' for the
// actual data and 'response' for the API response information.
type Response struct {
	Error *Error      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

type Error struct {
	Status  int         `json:"-"`
	Code    ErrorCode   `json:"code"`
	Message string      `json:"msg"`
	Details interface{} `json:"details,omitempty"`
}

func NewError(status int, code ErrorCode, msg string, err error) *Error {
	return &Error{
		Status:  status,
		Code:    code,
		Message: msg,
		Details: getDetailsFromErr(err),
	}
}

func NewErrBadRequest(code ErrorCode, msg string, err error) *Error {
	return &Error{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: msg,
		Details: getDetailsFromErr(err),
	}
}

func NewErrInternalServerError(code ErrorCode, msg string, err error) *Error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Code:    code,
		Message: msg,
		Details: getDetailsFromErr(err),
	}
}

func NewErrNotFound(code ErrorCode, msg string, err error) *Error {
	return &Error{
		Status:  http.StatusNotFound,
		Code:    code,
		Message: msg,
		Details: getDetailsFromErr(err),
	}
}

func OK(c *gin.Context, data interface{}) {
	resp := Response{
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func Err(c *gin.Context, err *Error) {
	resp := Response{
		Error: err,
	}
	c.JSON(err.Status, resp)
}

func ErrInternal(c *gin.Context, msg string, err error) {
	resp := Response{
		Error: NewErrInternalServerError(ErrorCodeInternal, msg, err),
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func getDetailsFromErr(err error) interface{} {
	if err == nil {
		return nil
	} else {
		return err.Error()
	}
}
