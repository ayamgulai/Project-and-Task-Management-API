package controllers

import (
	"mini-jira-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// show all logs
// GetTasks godoc
// @Summary Get all log activities
// @Description Retrieve list of tasks activities
// @Tags TaskLog
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /taskLogs [get]
func ShowTaskLogs(c *gin.Context) {
	logs, err := services.ShowTaskLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
