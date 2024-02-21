package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
	utils "github.com/gmessias/api-go-money/utils"
)

func CreateCash(c *gin.Context) {
	var cash models.Cash

	if err := c.ShouldBindJSON(&cash); err != nil {
		utils.MessageBadRequest(c, err.Error())
		return
	}

	result := database.DB.Create(&cash)
	if result.Error != nil {
		utils.MessageInternalError(c, "error creating cash record.")
		return
	}

	c.JSON(http.StatusCreated, cash)
}
