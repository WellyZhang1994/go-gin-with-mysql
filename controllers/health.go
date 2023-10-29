package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/WellyZhang1994/go-gin-with-mysql/models"
)

type HealthController struct{}
var healthModel = new(models.Health)

func (h HealthController) Status(c *gin.Context) {
	version, err := healthModel.GetVersion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"dbVersion": version})
}
