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
	DB := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database
	defer DB.Close()
	var StringToReturn string = fmt.Sprintf("Working! Client for this request : %s", c.Request.Header.Get("x-token"))
	c.String(http.StatusOK, StringToReturn)
}
