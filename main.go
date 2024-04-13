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

// Steps to test
// Since the authorization token expires every 2 hours, it is necessary to
// log in, then copy the token and replace the old token for the desired action
// (create, delete, update)
// Then it should work
