package category_routes

import "github.com/gin-gonic/gin"

func AllRoutes(r *gin.Engine) {
	GetRoutes(r)
	PostRoutes(r)
	DeleteRoutes(r)
	PatchRoutes(r)
}

func GetRoutes(r *gin.Engine) {
	r.GET("/category")
	r.GET("/category/:id")
}

func PostRoutes(r *gin.Engine) {
	r.POST("/category")
}

func DeleteRoutes(r *gin.Engine) {
	r.DELETE("/category/:id")
}

func PatchRoutes(r *gin.Engine) {
	r.PATCH("/category/:id")
}
