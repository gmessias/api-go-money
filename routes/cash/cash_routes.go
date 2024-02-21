package cash_routes

import (
	"github.com/gin-gonic/gin"
	handlecash "github.com/gmessias/api-go-money/handlers/cash"
)

func AllRoutes(r *gin.Engine) {
	GetRoutes(r)
	PostRoutes(r)
	DeleteRoutes(r)
	PatchRoutes(r)
}

func GetRoutes(r *gin.Engine) {
	r.GET("/cash", handlecash.ReadAllCash)
	r.GET("/cash/:id", handlecash.ReadCashPerId)
}

func PostRoutes(r *gin.Engine) {
	r.POST("/cash", handlecash.CreateCash)
}

func DeleteRoutes(r *gin.Engine) {
	r.DELETE("/cash/:id", handlecash.DeleteCash)
}

func PatchRoutes(r *gin.Engine) {
	r.PATCH("/cash/:id", handlecash.UpdateCash)
}
