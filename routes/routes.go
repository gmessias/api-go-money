package routes

import (
	"github.com/gin-gonic/gin"
	cashroutes "github.com/gmessias/api-go-money/routes/cash"
	categoryroutes "github.com/gmessias/api-go-money/routes/category"
	"log"
)

func HandleRequests() {
	r := gin.Default()

	cashroutes.AllRoutes(r)
	categoryroutes.AllRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
