package handle_category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func ReadAllCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)
	c.JSON(200, categories)
}

func ReadCategoryPerId(c *gin.Context) {
	var category models.Category
	id := c.Params.ByName("id")
	database.DB.First(&category, id)

	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status Not Found": "Category Id:" + id + "not found.",
		})
		return
	}

	c.JSON(http.StatusOK, category)
}
