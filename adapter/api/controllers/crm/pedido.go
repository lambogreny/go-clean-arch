package controllersCrm

import (
	"github.com/augusto/imersao5-esquenta-go/services/crm/erp_crm"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
