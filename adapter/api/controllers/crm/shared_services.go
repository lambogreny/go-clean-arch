package controllersCrm

import (
	"github.com/gin-gonic/gin"
)

type SharedServicesController struct {
}

func (t SharedServicesController) LogCsv(c *gin.Context) {
	// DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token"))

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, utils.Error{
	// 		StatusCode:  http.StatusBadRequest,
	// 		Message:     err.Error(),
	// 		Description: "Credenciais do cliente inválidas!",
	// 	})
	// 	return
	// }

	// //Criando o repositório
	// repo := crmRepository.NewPrdRepositoryDbErp(DB)

	// //Destinando o caso de uso
	// usecase := prd.NewProcessPrd(repo)

	// output, err := usecase.UseCaseSelect()

	// //Só testando o método de checagem de update
	// //teste, err := usecase.CheckUpdateCrm("adss")
	// //fmt.Println(teste)

	// if err != nil {
	// 	utils.LogFile("CRM/PRD", " SERVER_ERROR", "CRITICAL ", err.Error(), "Erro na manipulação do banco")
	// 	c.JSON(http.StatusInternalServerError, utils.Error{
	// 		StatusCode:  http.StatusInternalServerError,
	// 		Message:     err.Error(),
	// 		Description: "Erro de processamento",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, output)
}
