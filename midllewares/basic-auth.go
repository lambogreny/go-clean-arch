package midllewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	})
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

/*
	Função que é acionada pelo Middleware e verifica se na requisição existe um token do cliente
*/
func CheckClientToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		requiredToken := c.Request.Header.Get("x-token")

		if requiredToken == "" {
			respondWithError(c, 401, "Client key is necessary to use this api")
			return
		}

		c.Next()
	}
}

//Retorna um id de transação
func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}
