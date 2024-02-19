package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func UpdateCash(c *gin.Context) {
	var cash models.Cash
	id := c.Param("id")

	result := database.DB.First(&cash, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cash not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&cash); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&cash).UpdateColumns(cash).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating cash record",
		})
		return
	}

	c.JSON(http.StatusOK, cash)
}
