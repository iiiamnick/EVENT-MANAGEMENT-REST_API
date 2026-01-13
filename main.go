package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iiiamnick/EVENT-MANAGEMENT-REST_API/db"
	"github.com/iiiamnick/EVENT-MANAGEMENT-REST_API/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8081") // LOCALHOST : 8081

}
