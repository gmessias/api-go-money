package handle_cash

import (
	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
	utils "github.com/gmessias/api-go-money/utils"
)

func DeleteCash(c *gin.Context) {
	var cash models.Cash
	id := c.Params.ByName("id")

	if err := database.DB.First(&cash, id).Error; err != nil {
		utils.MessageNotFound(c, "cash not found.")
		return
	}

	result := database.DB.Delete(&cash, id)
	if result.Error != nil {
		utils.MessageInternalError(c, "failed to delete cash.")
		return
	}

	utils.MessageNoContentSuccess(c, "cash deleted.")
}
