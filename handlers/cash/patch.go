package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
	utils "github.com/gmessias/api-go-money/utils"
)

func UpdateCash(c *gin.Context) {
	var cash models.Cash
	id := c.Param("id")

	result := database.DB.First(&cash, id)
	if result.Error != nil {
		utils.MessageNotFound(c, "cash not found.")
		return
	}

	if err := c.ShouldBindJSON(&cash); err != nil {
		utils.MessageBadRequest(c, err.Error())
		return
	}

	if err := database.DB.Model(&cash).UpdateColumns(cash).Error; err != nil {
		utils.MessageInternalError(c, "error updating cash record.")
		return
	}

	c.JSON(http.StatusOK, cash)
}
