package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Details []interface{} `json:"details"`

	HTTPStatus int `json:"-"`
}

var DefaultErrorResponse = ErrorResponse{
	HTTPStatus: http.StatusInternalServerError,

	Code:    500,
	Message: "internal server error",
}

func RespondSuccess(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		SuccessResponse{
			Code:    0,
			Message: "successfully",
			Data:    data,
		},
	)
}

func RespondFailure(c *gin.Context, err error) {

	response := responseFromErr(err)

	c.JSON(
		response.HTTPStatus,
		response,
	)
}

func responseFromErr(err error) ErrorResponse {
	switch err {
	default:
		return DefaultErrorResponse
	}
}
