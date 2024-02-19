package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func ReadAllCash(c *gin.Context) {
	var cashList []models.Cash
	result := database.DB.Find(&cashList)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
		})
		return
	}

	if len(cashList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No cash records found",
		})
	}

	c.JSON(http.StatusOK, cashList)
}

func ReadCashPerId(c *gin.Context) {
	var cash models.Cash
	id := c.Param("id")

	result := database.DB.First(&cash, id)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cash with ID " + id + " not found.",
		})
		return
	}

	c.JSON(http.StatusOK, cash)
}
