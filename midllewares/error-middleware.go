package midllewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type appError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		// Use reflect.TypeOf(err.Err) to known the type of your error
		if error, ok := errors.Cause(err.Err).(*json.SyntaxError); ok {
			fmt.Println("AQUI O MID!!")
			c.JSON(400, gin.H{
				"error": error,
			})
			return
		}
	}
}
