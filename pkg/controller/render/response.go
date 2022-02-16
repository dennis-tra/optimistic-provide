package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is the structured type for every API response. The type contains two fields, 'result' for the
// actual data and 'response' for the API response information.
type Response struct {
	Error *Error      `json:"error"`
	Data  interface{} `json:"data"`
}

type Error struct {
	// Status  string `json:"status"`
	Message string `json:"msg"`
	Error   string `json:"err"`
}

func OK(c *gin.Context, data interface{}) {
	resp := Response{
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func BadRequest(c *gin.Context, err error, msg string) {
	resp := Response{
		Error: &Error{
			Message: msg,
			Error:   err.Error(),
		},
	}
	c.JSON(http.StatusBadRequest, resp)
}

func InternalServerError(c *gin.Context, err error, msg string) {
	resp := Response{
		Error: &Error{
			Message: msg,
			Error:   err.Error(),
		},
	}
	c.JSON(http.StatusInternalServerError, resp)
}
