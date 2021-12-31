package controllersCrm

import (
	"net/http"

	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/services/crm/erp_crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/prd"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
)

type PrdControllerErp struct {
}

/*
	Chama o service de integração
*/
func (t PrdControllerErp) CallPrdService(c *gin.Context) {

	resp := erp_crm.PrdService(c.Request.Header.Get("x-token"))

	if resp != nil {
		utils.LogFile("CRM/PRD", " SERVER_ERROR", "CRITICAL ", resp.Error(), "Erro na manipulação do banco no service")
		c.JSON(http.StatusConflict, utils.Error{
			StatusCode:  http.StatusConflict,
			Message:     resp.Error(),
			Description: "Falha ao realizar o processamento de integração",
		})
		return
	}

	c.String(http.StatusOK, "OK")
}

/*
	Devolve todos os produtos do ERP que devem ser integrados
*/
func (t PrdControllerErp) GetErp(c *gin.Context) {
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error{
			StatusCode:  http.StatusBadRequest,
			Message:     err.Error(),
			Description: "Credenciais do cliente inválidas!",
		})
		return
	}

	//Criando o repositório
	repo := crmRepository.NewPrdRepositoryDbErp(DB)

	//Destinando o caso de uso
	usecase := prd.NewProcessPrd(repo)

	output, err := usecase.UseCaseSelect()

	//Só testando o método de checagem de update
	//teste, err := usecase.CheckUpdateCrm("adss")
	//fmt.Println(teste)

	if err != nil {
		utils.LogFile("CRM/PRD", " SERVER_ERROR", "CRITICAL ", err.Error(), "Erro na manipulação do banco")
		c.JSON(http.StatusInternalServerError, utils.Error{
			StatusCode:  http.StatusInternalServerError,
			Message:     err.Error(),
			Description: "Erro de processamento",
		})
		return
	}

	c.JSON(http.StatusOK, output)
}
