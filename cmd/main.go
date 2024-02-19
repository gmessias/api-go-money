package cmd

import (
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
