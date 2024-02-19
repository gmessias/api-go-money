package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func DeleteCash(c *gin.Context) {
	var cash models.Cash
	id := c.Params.ByName("id")

	if err := database.DB.First(&cash, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cash not found.",
		})
		return
	}

	result := database.DB.Delete(&cash, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete cash",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"success": "Cash deleted.",
	})
}
