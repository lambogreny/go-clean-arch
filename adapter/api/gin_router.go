package api

import (
	"github.com/augusto/imersao5-esquenta-go/adapter/api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")

	health := new(controllers.HeathController)
	transaction := new(controllers.TransactionController)

	v1.GET("/health", health.Status)
	v1.POST("/transaction", transaction.NewTransaction)
	//router.GET("/transaction",Tr)

	return router
}
