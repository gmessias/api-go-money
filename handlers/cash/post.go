package handle_cash

import (
	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/internal/core/domain"
	"github.com/gmessias/api-go-money/internal/core/ports"
	"github.com/gmessias/api-go-money/repositories"
	"github.com/gmessias/api-go-money/utils"
	"net/http"
)

func CreateCash(c *gin.Context) {
	var cash domain.Cash
	var repo ports.CashRepository = &repositories.CashRepositoryImpl{}

	if err := c.ShouldBindJSON(&cash); err != nil {
		utils.MessageBadRequest(c, err.Error())
		return
	}

	result := repo.CreateCash(&cash)
	if result.Error != nil {
		utils.MessageInternalError(c, "error creating cash record.")
		return
	}

	c.JSON(http.StatusCreated, cash)
}
