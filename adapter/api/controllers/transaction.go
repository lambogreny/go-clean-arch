package controllers

import (
	"database/sql"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/adapter/repository"
	"github.com/augusto/imersao5-esquenta-go/usecase/process_transaction"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TransactionController struct {
}

func (t TransactionController) NewTransaction(c *gin.Context) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)
	output, err := usecase.Execute(inputData)

	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(output)

	//c.String(http.StatusOK, "Working!")
	c.JSON(http.StatusOK, output)
}
