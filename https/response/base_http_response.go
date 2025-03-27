package https

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHttpResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type ListResponse struct {
	Total int   `json:"total"`
	List  []any `json:"list"`
}

func SuccessResponse(c *gin.Context, message string, data any) {
	response := BaseHttpResponse{
		StatusCode: http.StatusOK, Message: message, Data: data}
	c.JSON(response.StatusCode, response)
}
func ErrorResponse(c *gin.Context, statusCode int, message string, data any) {
	response := BaseHttpResponse{
		StatusCode: statusCode, Message: message, Data: data}
	c.JSON(statusCode, response)
}
