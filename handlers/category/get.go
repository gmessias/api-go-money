package handle_category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func ReadAllCategories(c *gin.Context) {
	var categories []models.Category
	result := database.DB.Find(&categories)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
		})
		return
	}

	if len(categories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No categories records found",
		})
	}

	c.JSON(http.StatusOK, categories)
}

func ReadCategoryPerId(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	result := database.DB.First(&category, id)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category with ID " + id + " not found.",
		})
		return
	}

	c.JSON(http.StatusOK, category)
}
