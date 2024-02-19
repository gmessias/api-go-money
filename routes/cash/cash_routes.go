package cash_routes

import (
	"github.com/gin-gonic/gin"
	handle_cash "github.com/gmessias/api-go-money/handlers/cash"
)

func AllRoutes(r *gin.Engine) {
	GetRoutes(r)
	PostRoutes(r)
	DeleteRoutes(r)
	PatchRoutes(r)
}

func GetRoutes(r *gin.Engine) {
	r.GET("/cash", handle_cash.ReadAllCash)
	r.GET("/cash/:id", handle_cash.ReadCashPerId)
}

func PostRoutes(r *gin.Engine) {
	r.POST("/cash", handle_cash.CreateCash)
}

func DeleteRoutes(r *gin.Engine) {
	r.DELETE("/cash/:id", handle_cash.DeleteCash)
}

func PatchRoutes(r *gin.Engine) {
	r.PATCH("/cash/:id", handle_cash.UpdateCash)
}
