package handle_cash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/models"
)

func DeleteCash(c *gin.Context) {
	var cash models.Cash
	id := c.Params.ByName("id")
	database.DB.Delete(&cash, id)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Cash deletead.",
	})
}
