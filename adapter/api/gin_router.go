package api

import (
	"github.com/augusto/imersao5-esquenta-go/adapter/api/controllers"
	controllersCrm "github.com/augusto/imersao5-esquenta-go/adapter/api/controllers/crm"
	"github.com/augusto/imersao5-esquenta-go/midllewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(midllewares.BasicAuth()) //Basic auth
	// router.Use(midllewares.RequestIdMiddleware()) //Id do request
	router.Use(midllewares.CheckClientToken()) //Valida se há api key
	router.Use(midllewares.ErrorHandle())      // Nao esta sendo utilziado

	v1 := router.Group("/v1")

	health := new(controllers.HeathController)
	transaction := new(controllers.TransactionController)
	approval := new(controllers.ApprovalController)
	queryInfo := new(controllers.QueryInfoController)

	crmPrd := new(controllersCrm.PrdControllerErp)
	Cfr := new(controllersCrm.CfrControllerErp)
	pedido := new(controllersCrm.PedidoControllerErp)

	v1.GET("/health", health.Status)
	v1.POST("/transaction", transaction.NewTransaction)
	v1.GET("/transaction", transaction.GetTransaction)
	v1.DELETE("/transaction", transaction.DeleteTransaction)

	v1.GET("/approval", approval.GetApproval)
	v1.POST("/approval", approval.InteractApproval)

	v1.POST("/queryInfo/cards", queryInfo.GetCards)

	//CRM
	v1.GET("/crm/erp/prd", crmPrd.GetErp)

	//Crm Services
	v1.GET("/crm/erp/prd/service", crmPrd.CallPrdService)       //Leva os dados para a PRd
	v1.GET("/crm/erp/cfr/service", Cfr.CallCfrService)          //Leva os dados para a CFR
	v1.GET("/crm/erp/account/service", Cfr.CallAccountService)  //Leva os dados para a Account
	v1.GET("/crm/erp/pedido/service", pedido.CallPedidoService) //Leva os dados para a CPV e IPV

	return router
}
