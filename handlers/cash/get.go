package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func ReadAllCash(c *gin.Context) {
	var cashList []models.Cash
	database.DB.Find(&cashList)
	c.JSON(200, cashList)
}

func ReadCashPerId(c *gin.Context) {
	var cash models.Cash
	id := c.Params.ByName("id")
	database.DB.First(&cash, id)

	if cash.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status Not Found": "Cash Id:" + id + "not found.",
		})
		return
	}

	c.JSON(http.StatusOK, cash)
}
