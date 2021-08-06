package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type error struct {
	errorMessage string `json:"error_message,omitempty"`
	errorCode    int    `json:"error_code,omitempty"`
}

func RecoverOfPanic(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		error := error{
			errorMessage: err,
			errorCode:    http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, gin.
			H{"error_message": error.errorMessage, "error_code": error.errorCode})
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
