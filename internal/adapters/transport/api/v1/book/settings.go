package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// httpResponseError provides a base error type for all errors.
type httpResponseError struct {
	Type          httpErrType `json:"-"`
	Message       string      `json:"message"`
	Code          string      `json:"code,omitempty"`
	Details       interface{} `json:"details,omitempty"`
	InvalidFields interface{} `json:"invalidFields,omitempty"`
}

// httpErrType is used to define error type.
type httpErrType string

const (
	// ErrorTypeServer is an "unexpected" internal server error.
	ErrorTypeServer httpErrType = "server"
	// ErrorTypeClient is an "expected" business error.
	ErrorTypeClient httpErrType = "client"
)

// wrapHandler provides unified error handling for all handlers.
func wrapHandler(handler func(c *gin.Context) (interface{}, *httpResponseError)) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := handler(c)

		// handle panics
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			}
		}()

		// check if middleware
		if body == nil && err == nil {
			return
		}

		if err != nil {
			switch err.Type {
			case ErrorTypeServer:
				c.AbortWithStatusJSON(http.StatusInternalServerError, body)
			case ErrorTypeClient:
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, body)
			default:
				return
			}
		}

		c.JSON(http.StatusOK, body)
	}
}
