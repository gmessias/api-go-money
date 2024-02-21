package handle_cash

import (
	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/internal/core/domain"
	"github.com/gmessias/api-go-money/internal/core/ports"
	"github.com/gmessias/api-go-money/repositories"
	utils "github.com/gmessias/api-go-money/utils"
)

func DeleteCash(c *gin.Context) {
	var cash domain.Cash
	id := c.Params.ByName("id")

	var repo ports.CashRepository = &repositories.CashRepositoryImpl{}

	if err := repo.GetCashPerId(&cash, id).Error; err != nil {
		utils.MessageNotFound(c, "cash not found.")
		return
	}

	//result := database.DB.Delete(&cash, id)
	result := repo.DeleteCash(cash, id)

	if result.Error != nil {
		utils.MessageInternalError(c, "failed to delete cash.")
		return
	}

	utils.MessageNoContentSuccess(c, "cash deleted.")
}
