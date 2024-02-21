package category_routes

import (
	"github.com/gin-gonic/gin"
	handlecategory "github.com/gmessias/api-go-money/handlers/category"
)

func AllRoutes(r *gin.Engine) {
	GetRoutes(r)
	PostRoutes(r)
	DeleteRoutes(r)
	PatchRoutes(r)
}

func GetRoutes(r *gin.Engine) {
	r.GET("/category", handlecategory.ReadAllCategories)
	r.GET("/category/:id", handlecategory.ReadCategoryPerId)
}

func PostRoutes(r *gin.Engine) {
	r.POST("/category", handlecategory.CreateCategory)
}

func DeleteRoutes(r *gin.Engine) {
	r.DELETE("/category/:id", handlecategory.DeleteCategory)
}

func PatchRoutes(r *gin.Engine) {
	r.PATCH("/category/:id", handlecategory.UpdateCategory)
}
