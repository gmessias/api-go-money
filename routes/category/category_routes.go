package category_routes

import (
	"github.com/gin-gonic/gin"
	handle_category "github.com/gmessias/api-go-money/handlers/category"
)

func AllRoutes(r *gin.Engine) {
	GetRoutes(r)
	PostRoutes(r)
	DeleteRoutes(r)
	PatchRoutes(r)
}

func GetRoutes(r *gin.Engine) {
	r.GET("/category", handle_category.ReadAllCategories)
	r.GET("/category/:id", handle_category.ReadCategoryPerId)
}

func PostRoutes(r *gin.Engine) {
	r.POST("/category", handle_category.CreateCategory)
}

func DeleteRoutes(r *gin.Engine) {
	r.DELETE("/category/:id", handle_category.DeleteCategory)
}

func PatchRoutes(r *gin.Engine) {
	r.PATCH("/category/:id", handle_category.UpdateCategory)
}
