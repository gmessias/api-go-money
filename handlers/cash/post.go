package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func CreateCash(c *gin.Context) {
	var cash models.Cash
	if err := c.ShouldBindJSON(&cash); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&cash)
	c.JSON(http.StatusOK, cash)
}
