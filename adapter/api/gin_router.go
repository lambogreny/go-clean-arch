package api

import (
	"github.com/augusto/imersao5-esquenta-go/adapter/api/controllers"
	"github.com/augusto/imersao5-esquenta-go/midllewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(midllewares.BasicAuth())           //Basic auth
	router.Use(midllewares.RequestIdMiddleware()) //Id do request
	router.Use(midllewares.CheckClientToken())    //Valida se há api key
	//router.Use(midllewares.ErrorHandle()) // Nao esta sendo utilziado

	v1 := router.Group("/v1")

	health := new(controllers.HeathController)
	transaction := new(controllers.TransactionController)

	v1.GET("/health", health.Status)
	v1.POST("/transaction", transaction.NewTransaction)
	v1.GET("/transaction", transaction.GetTransaction)
	//router.GET("/transaction",Tr)

	return router
}
