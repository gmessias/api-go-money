package routes

import (
	"github.com/gin-gonic/gin"
	handle "github.com/gmessias/api-go-money/handlers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/cash", handle.ReadAllCash)
	r.GET("/cash/:id", handle.ReadCashPerId)
	r.POST("/cash", handle.CreateCash)

	r.Run(":8080")
}
