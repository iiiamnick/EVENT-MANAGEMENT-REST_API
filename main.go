package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iiiamnick/GOLANG--REST-API-EVENT-BOOKING-.git/db"
	"github.com/iiiamnick/GOLANG--REST-API-EVENT-BOOKING-.git/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8081") // LOCALHOST : 8081

}
