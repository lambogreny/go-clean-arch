package controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/augusto/imersao5-esquenta-go/adapter/repository"
	"github.com/augusto/imersao5-esquenta-go/usecase/process_transaction"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
}

func (t TransactionController) GetTransaction(c *gin.Context) {
	DB := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database

	id, has := c.GetQuery("id")

	if has {
		fmt.Println("Tem id!", id)
	} else {
		fmt.Println("Não tem id, porque o o has é :", has)
	}

	//Criando o repositório
	repo := repository.NewTransactionRepositoryDb(DB)

	//Destinando o caso de uso
	usecase := process_transaction.NewProcessTransaction(repo)

	//Executando
	output, err := usecase.GetAll()

	if err != nil {
		log.Fatal(err)
		return
	}

	//fmt.Println("o tipo da saída é : ", reflect.TypeOf(output))
	//fmt.Println("O tamanho do array é :", len(output))

	outputLength := len(output)
	if outputLength == 0 {
		c.JSON(http.StatusNoContent, nil)
	}

	c.JSON(http.StatusOK, output)

}

func (t TransactionController) NewTransaction(c *gin.Context) {

	DB := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database

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
