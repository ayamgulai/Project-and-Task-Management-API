package services

import (
	"mini-jira-backend/models"
	"mini-jira-backend/repositories"
)

// show all logs
func ShowTaskLogs() ([]models.TaskLog, error) {
	return repositories.ShowTaskLogs()
}