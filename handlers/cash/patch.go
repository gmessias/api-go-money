package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func UpdateCash(c *gin.Context) {
	var cash models.Cash
	id := c.Params.ByName("id")
	database.DB.First(&cash, id)

	if err := c.ShouldBindJSON(&cash); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&cash).UpdateColumns(cash)
	c.JSON(http.StatusOK, cash)
}
