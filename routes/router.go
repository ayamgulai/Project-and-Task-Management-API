package routes

import (
	"mini-jira-backend/controllers"
	"mini-jira-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// auth routes
	router.POST("/login", controllers.Login)

	// only admin can register new user, so we protect this route with auth middleware
	adminAuthority := router.Group("/")
	adminAuthority.Use(
		middlewares.AuthMiddleware(),
		middlewares.AdminOnly(),
	)
	{
		adminAuthority.POST("/register", controllers.Register)
		adminAuthority.GET("/taskLogs", controllers.ShowTaskLogs)
	}

	// project routes
	projectGroup := router.Group("/projects")
	projectGroup.Use(middlewares.AuthMiddleware())
	projectGroup.GET("", controllers.GetProjects)
	projectGroup.GET("/:id", controllers.GetProjectByID)
	projectGroup.GET("/:id/tasks", controllers.GetTasksByProjectID)
	projectGroup.POST("", controllers.CreateProject)
	projectGroup.PUT("/:id", controllers.UpdateProject)

	// task routes
	taskGroup := router.Group("/tasks")
	taskGroup.Use(middlewares.AuthMiddleware())
	taskGroup.GET("", controllers.GetTasks)
	taskGroup.GET("/:id", controllers.GetTaskByID)
	taskGroup.GET("/assignee/:assigneeId", controllers.GetTasksByAssigneeID)
	taskGroup.POST("", controllers.CreateTask)
	taskGroup.DELETE("/:id", controllers.DeleteTask)
	taskGroup.PUT("/:id/status", controllers.UpdateTaskStatus)

	return router

}
