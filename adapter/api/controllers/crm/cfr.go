package controllersCrm

import (
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/services/crm/erp_crm"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CfrControllerErp struct {
}

//Cfr para p Ti9
func (t CfrControllerErp) CallCfrService(c *gin.Context) {
	resp := erp_crm.CfrService(c.Request.Header.Get("x-token"))

	if resp != nil {
		//utils.LogFile("CRM/CFR", " CONTROLLER", "CRITICAL ", resp.Error(), "Erro na manipulação do banco no service")
		c.JSON(http.StatusConflict, utils.Error{
			StatusCode:  http.StatusConflict,
			Message:     resp.Error(),
			Description: "Falha ao realizar o processamento de integração",
		})
		return
	}

	c.String(http.StatusOK, "OK")
}

func (t CfrControllerErp) CallAccountService(c *gin.Context) {

	fmt.Println("Levar os dados do TI9 para o CRM")

	resp := erp_crm.AccountService(c.Request.Header.Get("x-token"))

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
