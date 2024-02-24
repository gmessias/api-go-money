package handle_category

import (
	"github.com/gmessias/api-go-money/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/utils"
)

func CreateCategory(c *gin.Context) {
	var category domain.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		utils.MessageBadRequest(c, err.Error())
		return
	}

	result := database.DB.Create(&category)
	if result.Error != nil {
		utils.MessageInternalError(c, "error creating category record.")
		return
	}

	c.JSON(http.StatusCreated, category)
}
