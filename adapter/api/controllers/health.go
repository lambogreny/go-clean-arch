package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HeathController struct {
}

func (h HeathController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}