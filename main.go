package main

import (
	"github.com/gin-gonic/gin"

	"example.com/REST-API/db"
	"example.com/REST-API/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRouts(server)

	server.Run(":8080") //localhost:8080

}
