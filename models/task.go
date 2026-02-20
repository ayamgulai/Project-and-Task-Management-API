package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	AssigneeID  *int      `json:"assignee_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTaskInput struct {
	ProjectID   int    `json:"project_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Priority    string `json:"priority" binding:"required"`
}

type UpdateTaskStatusInput struct {
	Status string `json:"status" binding:"required"`
}
