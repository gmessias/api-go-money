package handle_category

import (
	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/internal/core/domain"
	"github.com/gmessias/api-go-money/utils"
)

func DeleteCategory(c *gin.Context) {
	var category domain.Category
	id := c.Params.ByName("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		utils.MessageNotFound(c, "category not found.")
		return
	}

	result := database.DB.Delete(&category, id)
	if result.Error != nil {
		utils.MessageInternalError(c, "failed to delete category.")
		return
	}

	utils.MessageNoContentSuccess(c, "category deleted.")
}
