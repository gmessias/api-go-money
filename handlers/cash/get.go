package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/internal/core/domain"
	utils "github.com/gmessias/api-go-money/utils"
)

func ReadAllCash(c *gin.Context) {
	var cashList []domain.Cash
	result := database.DB.Find(&cashList)

	if result.Error != nil {
		utils.MessageInternalError(c, "internal server error.")
		return
	}

	if len(cashList) == 0 {
		utils.MessageNotFound(c, "no cash records found.")
		return
	}

	c.JSON(http.StatusOK, cashList)
}

func ReadCashPerId(c *gin.Context) {
	var cash domain.Cash
	id := c.Param("id")

	result := database.DB.First(&cash, id)

	if result.Error != nil || result.RowsAffected == 0 {
		utils.MessageNotFound(c, "cash not found.")
		return
	}

	c.JSON(http.StatusOK, cash)
}
