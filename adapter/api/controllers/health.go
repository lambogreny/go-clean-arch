package controllers

import (
	"fmt"
	"net/http"

	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
)

type HeathController struct {
}

func (h HeathController) Status(c *gin.Context) {
	//Apenas abre e fecha conexão com o banco
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database

	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     err.Error(),
				Description: "Credenciais do cliente inválidas!",
			})
			return
		}
	}
	defer DB.Close()
	var StringToReturn string = fmt.Sprintf("Working! Client for this request : %s", c.Request.Header.Get("x-token"))
	c.String(http.StatusOK, StringToReturn)
}
