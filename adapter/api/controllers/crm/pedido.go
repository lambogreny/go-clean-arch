package controllersCrm

import (
	"fmt"
	"net/http"

	"github.com/augusto/imersao5-esquenta-go/services/crm/erp_crm"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
)

type PedidoControllerErp struct {
}

func (t PedidoControllerErp) CallPedidoService(c *gin.Context) {

	resp := erp_crm.PedidoService(c.Request.Header.Get("x-token"))

	if resp != nil {
		utils.LogFile("CRM/PEDIDO", " SERVER_ERROR", "CRITICAL ", resp.Error(), "Erro na manipulação do banco no service")
		c.JSON(http.StatusConflict, utils.Error{
			StatusCode:  http.StatusConflict,
			Message:     resp.Error(),
			Description: "Falha ao realizar o processamento de integração",
		})
		return
	}

	c.String(http.StatusOK, "OK")

}

func (t PedidoControllerErp) CallPedidoQuoteErpCrm(c *gin.Context) {
	fmt.Println("Serviço que leva dados do ERP para o CRM da quote")

	resp := erp_crm.QuoteService(c.Request.Header.Get("x-token"))

	if resp != nil {
		c.JSON(http.StatusConflict, utils.Error{
			StatusCode:  http.StatusConflict,
			Message:     resp.Error(),
			Description: "Falha ao realizar o processamento de integração",
		})
		return
	}

	c.String(http.StatusOK, "OK")
}

func (t PedidoControllerErp) CallQuoteItemErpCrm(c *gin.Context) {
	fmt.Println("Serviço que leva dados do ERP para o CRM da quote item")

	resp := erp_crm.QuoteItemService(c.Request.Header.Get("x-token"))

	if resp != nil {
		c.JSON(http.StatusConflict, utils.Error{
			StatusCode:  http.StatusConflict,
			Message:     resp.Error(),
			Description: "Falha ao realizar o processamento de integração",
		})
		return
	}

	c.String(http.StatusOK, "OK")
}
