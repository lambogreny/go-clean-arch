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
	//router.Use(midllewares.ErrorHandle())         // Nao esta sendo utilziado

	//docs2.SwaggerInfo.BasePath = "/v1"

	v1 := router.Group("/v1")

	health := new(controllers.HeathController)
	transaction := new(controllers.TransactionController)
	approval := new(controllers.ApprovalController)

	v1.GET("/health", health.Status)
	v1.POST("/transaction", transaction.NewTransaction)
	v1.GET("/transaction", transaction.GetTransaction)
	v1.DELETE("/transaction", transaction.DeleteTransaction)

	v1.GET("/approval", approval.GetApproval)

	//v1.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
