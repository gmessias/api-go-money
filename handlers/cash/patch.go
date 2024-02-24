package handle_cash

import (
	"github.com/gmessias/api-go-money/internal/core/domain"
	"github.com/gmessias/api-go-money/internal/core/ports"
	"github.com/gmessias/api-go-money/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/utils"
)

func UpdateCash(c *gin.Context) {
	var cash domain.Cash
	id := c.Param("id")

	var repo ports.CashRepository = &repositories.CashRepositoryImpl{}

	result := repo.GetCashPerId(&cash, id)
	if result.Error != nil {
		utils.MessageNotFound(c, "cash not found.")
		return
	}

	if err := c.ShouldBindJSON(&cash); err != nil {
		utils.MessageBadRequest(c, err.Error())
		return
	}

	if err := repo.UpdateCash(&cash).Error; err != nil {
		utils.MessageInternalError(c, "error updating cash record.")
		return
	}

	c.JSON(http.StatusOK, cash)
}
