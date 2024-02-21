package handle_category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
	utils "github.com/gmessias/api-go-money/utils"
)

func ReadAllCategories(c *gin.Context) {
	var categories []models.Category
	result := database.DB.Find(&categories)

	if result.Error != nil {
		utils.MessageInternalError(c, "internal server error.")
		return
	}

	if len(categories) == 0 {
		utils.MessageNotFound(c, "no categories records found.")
		return
	}

	c.JSON(http.StatusOK, categories)
}

func ReadCategoryPerId(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	result := database.DB.First(&category, id)

	if result.Error != nil || result.RowsAffected == 0 {
		utils.MessageNotFound(c, "category not found.")
		return
	}

	c.JSON(http.StatusOK, category)
}
