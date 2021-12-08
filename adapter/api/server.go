package api

import (
	"github.com/augusto/imersao5-esquenta-go/entity"
	"github.com/augusto/imersao5-esquenta-go/usecase/process_transaction"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebServer struct {
	Repository entity.TransactionRepository
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.POST("/transaction", w.process)
	e.Logger.Fatal(e.Start(":8586"))

}

func (w WebServer) process(c echo.Context) error {
	transactionDto := &process_transaction.TransactionDtoInput{}
	c.Bind(transactionDto)
	usecase := process_transaction.NewProcessTransaction(w.Repository)
	output, _ := usecase.Execute(*transactionDto)
	return c.JSON(http.StatusOK, output)
}
