package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/gin-gonic/gin"
)

type HeathController struct {
}

func (h HeathController) Status(c *gin.Context) {
	//Apenas abre e fecha conexão com o banco
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database
	fmt.Println("Variavel de ambiente teste : ", os.Getenv("TESTE"))

	//Logando no database
	utils.LogDatabase("health", "health", "I", "health", false, "health")
	utils.LogDatabaseDetails("health", "health", "health", "health", "health")

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
	var StringToReturn string = fmt.Sprintf("Working! Client for this request : %s Variavel de ambiente teste : %s", c.Request.Header.Get("x-token"), os.Getenv("TESTE"))
	c.String(http.StatusOK, StringToReturn)
}

func (h HeathController) GenerateLogs(c *gin.Context) {
	//Apenas abre e fecha conexão com o banco
	DB, err := utils.DatabaseConnection(c.Request.Header.Get("x-token")) //Função de database
	fmt.Println("Variavel de ambiente teste : ", os.Getenv("TESTE"))

	//Logando no database
	for i := 0; i < 40; i++ {
		utils.LogDatabase("fakeLogs", "fakeLogs", "I", "fakeLogs", false, "fakeLogs")
		utils.LogDatabase("fakeLogs", "fakeLogs", "I", "fakeLogs", true, "fakeLogs")
		utils.LogDatabaseDetails("fakeLogs", "fakeLogs", "fakeLogs", "fakeLogs", "fakeLogs")
	}

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
	var StringToReturn string = fmt.Sprintf("Logs done! Client for this request : %s Variavel de ambiente teste : %s", c.Request.Header.Get("x-token"), os.Getenv("TESTE"))
	c.String(http.StatusOK, StringToReturn)
}
