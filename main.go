// @title Mini Jira Backend
// @version 1.0
// @description Project and Task Management API
// @host project-and-task-management-api-production.up.railway.app
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"mini-jira-backend/configs"
	_ "mini-jira-backend/docs"
	"mini-jira-backend/routes"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
	PORT := ":8080"
	closeDB := configs.ConnectDB()
	r := routes.RegisterRoutes()
	defer closeDB()

	configs.RunMigrations()
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(PORT)
}
