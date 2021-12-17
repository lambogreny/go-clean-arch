package controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/augusto/imersao5-esquenta-go/adapter/repository"
	"github.com/augusto/imersao5-esquenta-go/usecase/process_approval"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
)

type ApprovalController struct {
}

func (t ApprovalController) GetApproval(c *gin.Context) {
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

	var inputData process_approval.ApprovalDtoInput
	fmt.Println(c.BindQuery(inputData))

	//if err := c.BindQuery(&inputData); err != nil {
	if err := c.ShouldBind(&inputData); err != nil {
		fmt.Println(reflect.TypeOf(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Data input validation",
			"error":   err.Error(),
		})
		return
	}

	//Criando o repositório
	repo := repository.NewApprovalRepositoryDb(DB)

	//Destinando o caso de uso
	usecase := process_approval.NewApprovalTransaction(repo)

	//Executando
	output, err := usecase.GetAll()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.Error{
			StatusCode:  http.StatusInternalServerError,
			Message:     err.Error(),
			Description: "Erro de processamento",
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (t ApprovalController) InteractApproval(c *gin.Context) {
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token"))
	fmt.Println(DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error{
			StatusCode:  http.StatusBadRequest,
			Message:     err.Error(),
			Description: "Credenciais do cliente inválidas!",
		})
		return
	}

	var inputData process_approval.ApprovalDtoInteractionInput
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

	repo := repository.NewApprovalRepositoryDb(DB)

	usecase := process_approval.NewApprovalTransaction(repo)

	outputErr := usecase.Interact(inputData)

	if outputErr != nil {
		c.JSON(http.StatusConflict, utils.Error{
			StatusCode:  http.StatusConflict,
			Message:     outputErr.Error(),
			Description: "Erro no processamento da transação!",
		})
		return
	}
	c.String(http.StatusOK, "Transação realizada com sucesso!")
}
