package controllers

import (
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/adapter/repository"
	"github.com/augusto/imersao5-esquenta-go/usecase/process_transaction"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController struct {
}

func (t TransactionController) NewTransaction(c *gin.Context) {

	DB := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database

	//Header do request
	//fmt.Println(c.Request.Header.Get("x-token"))

	var inputData process_transaction.TransactionDtoInput

	//Validando a entrada
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Data input validation",
			"error":   err.Error(),
		})
		return
	}

	repo := repository.NewTransactionRepositoryDb(DB)
	usecase := process_transaction.NewProcessTransaction(repo)
	output, err := usecase.Execute(inputData)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(output)

	c.JSON(http.StatusOK, output)
}
