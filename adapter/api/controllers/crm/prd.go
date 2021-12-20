package controllersCrm

import (
	"fmt"
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/prd"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PrdControllerErp struct {
}

func (t PrdControllerErp) Get(c *gin.Context) {
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error{
			StatusCode:  http.StatusBadRequest,
			Message:     err.Error(),
			Description: "Credenciais do cliente inválidas!",
		})
		return
	}

	//utils.LogFile("CRM/PRD", " prd", "CRITICAL ", "err.Error()", "")

	//Criando o repositório
	repo := crmRepository.NewPrdRepositoryDbErp(DB)

	//Destinando o caso de uso
	usecase := prd.NewProcessPrd(repo)

	output, err := usecase.Repository.Select()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.Error{
			StatusCode:  http.StatusInternalServerError,
			Message:     err.Error(),
			Description: "Erro de processamento",
		})
		return
	}

	fmt.Println("O resultado do caso de uso é", output)

	c.JSON(http.StatusOK, output)
}
