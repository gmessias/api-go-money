package routes

import (
	"github.com/gin-gonic/gin"
	handle "github.com/gmessias/api-go-money/handlers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/cash", handle.GetAllCash)
	r.GET("/:movimentacao", handle.GetIdCash)

	r.Run(":9001")
}
