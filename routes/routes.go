package routes

import (
	"github.com/gin-gonic/gin"
	cash_routes "github.com/gmessias/api-go-money/routes/cash"
	category_routes "github.com/gmessias/api-go-money/routes/category"
)

func HandleRequests() {
	r := gin.Default()

	cash_routes.AllRoutes(r)
	category_routes.AllRoutes(r)

	r.Run(":8080")
}
