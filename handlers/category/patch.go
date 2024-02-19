package handle_category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	result := database.DB.First(&category, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&category).UpdateColumns(category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating category record",
		})
		return
	}

	c.JSON(http.StatusOK, category)
}
