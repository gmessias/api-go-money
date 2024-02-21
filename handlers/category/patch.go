package handle_category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
	utils "github.com/gmessias/api-go-money/utils"
)

func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	result := database.DB.First(&category, id)
	if result.Error != nil {
		utils.MessageNotFound(c, "category not found.")
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		utils.MessageBadRequest(c, err.Error())
		return
	}

	if err := database.DB.Model(&category).UpdateColumns(category).Error; err != nil {
		utils.MessageInternalError(c, "error updating category record.")
		return
	}

	c.JSON(http.StatusOK, category)
}
