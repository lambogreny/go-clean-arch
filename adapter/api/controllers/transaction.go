package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"

	"github.com/augusto/imersao5-esquenta-go/adapter/repository"
	"github.com/augusto/imersao5-esquenta-go/usecase/process_transaction"
	"github.com/augusto/imersao5-esquenta-go/utils"
)

type TransactionController struct {
}

func (t TransactionController) DeleteTransaction(c *gin.Context) {
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error{
			StatusCode:  http.StatusBadRequest,
			Message:     err.Error(),
			Description: "Credenciais do cliente inválidas!",
		})
		return
	}

	repo := repository.NewTransactionRepositoryDb(DB)

	usecase := process_transaction.NewProcessTransaction(repo)

	resp := usecase.DeleteTransaction()

	if resp != nil {
		c.JSON(http.StatusInternalServerError, utils.Error{
			StatusCode:  http.StatusInternalServerError,
			Message:     err.Error(),
			Description: "Erro de processamento",
		})
		return
	}

	c.JSON(http.StatusOK, process_transaction.DeleteTransactionOutput{Message: "Transação finalizada com sucesso!"})

}

func (t TransactionController) GetTransaction(c *gin.Context) {
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database
	fmt.Println("O ERRO é : ", err)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error{
			StatusCode:  http.StatusBadRequest,
			Message:     err.Error(),
			Description: "Credenciais do cliente inválidas!",
		})
		return
	}

	fmt.Println(DB)

	//id, has := c.GetQuery("id")
	id, _ := c.GetQuery("id")

	//Criando o repositório
	repo := repository.NewTransactionRepositoryDb(DB)

	//Destinando o caso de uso
	usecase := process_transaction.NewProcessTransaction(repo)

	//Executando
	output, err := usecase.GetAll(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error{
			StatusCode:  http.StatusInternalServerError,
			Message:     err.Error(),
			Description: "Erro de processamento",
		})
		return
	}

	//fmt.Println("o tipo da saída é : ", reflect.TypeOf(output))

	outputLength := len(output)
	if outputLength == 0 {
		utils.LogFile("INFO", " transaction", "INFO ", "O banco está sem dados de transação", "")
		c.JSON(http.StatusNoContent, nil)
	}

	c.JSON(http.StatusOK, output)

}

func (t TransactionController) NewTransaction(c *gin.Context) {

	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database

	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				Description: "Credenciais do cliente inválidas!",
			})
			return
		}
	}

	var inputData process_transaction.TransactionDtoInput

	//Validando a entrada
	if err := c.ShouldBindJSON(&inputData); err != nil {
		fmt.Println(reflect.TypeOf(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Data input validation",
			"error":   err.Error(),
		})
		return
	}

	//Criando o repositório
	repo := repository.NewTransactionRepositoryDb(DB)

	//Destinando o caso de uso
	usecase := process_transaction.NewProcessTransaction(repo)

	//Executando
	output, err := usecase.Execute(inputData)

	//Tratando se a transação for rejeitada
	if output.ErrorMessage != "" {
		fmt.Println("Aqui o cenário de transação que não deu certo!")
		c.JSON(http.StatusConflict, output)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, output)
}
