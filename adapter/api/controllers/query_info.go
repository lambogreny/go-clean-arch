package controllers

import (
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueryInfoController struct {
}

func (t QueryInfoController) GetCards(c *gin.Context) {
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error{
			StatusCode:  http.StatusBadRequest,
			Message:     err.Error(),
			Description: "Credenciais do cliente inválidas!",
		})
		return
	}
	fmt.Println(DB)

	c.String(http.StatusOK, "Contruindo!!")
}
