package main

import (
	"gotodo/db"
	"gotodo/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("GOTODO_ENV")
	if env == "" {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")

	db.InitDB(env)

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
