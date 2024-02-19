package handle_category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Params.ByName("id")
	database.DB.First(&category, id)

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&category).UpdateColumns(category)
	c.JSON(http.StatusOK, category)
}
