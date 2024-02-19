package handle_category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Params.ByName("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found.",
		})
		return
	}

	result := database.DB.Delete(&category, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete category",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"success": "Category deleted.",
	})
}
