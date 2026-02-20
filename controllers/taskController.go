package controllers

import (
	"mini-jira-backend/models"
	"mini-jira-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get all tasks

// GetTasks godoc
// @Summary Get all tasks
// @Description Retrieve list of tasks
// @Tags Task
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	tasks, err := services.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// get task by id
// GetTaskByID godoc
// @Summary Get task by ID
// @Description Get single task by its ID
// @Tags Task
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /tasks/{id} [get]
func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	task, err := services.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

// create new task
// CreateTask godoc
// @Summary Create new task
// @Description Create a task
// @Tags Task
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.CreateTaskInput true "Task payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userEmail := c.GetString("user_email")
	createdTask, err := services.CreateTask(task, userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"task": createdTask})
}

// delete task by id
// DeleteTask godoc
// @Summary Delete task
// @Description Delete a task by ID
// @Tags Task
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	userEmail := c.GetString("user_email")
	userId := c.GetInt("user_id")
	userRole := c.GetString("role")
	isSuccess, err := services.DeleteTask(id, userEmail, userRole, userId)
	if !isSuccess {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

// update task status by id
// UpdateTaskStatus godoc
// @Summary Update task status
// @Description Update the status of a task
// @Tags Task
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param body body models.UpdateTaskStatusInput true "Task Status Update payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tasks/{id}/status [put]
func UpdateTaskStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userEmail := c.GetString("user_email")
	userRole := c.GetString("role")
	logUpdate, err := services.UpdateTaskStatus(id, req.Status, userEmail, userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task status updated successfully", "log": logUpdate})
}

// get tasks by project id
// GetTasksByProjectID godoc
// @Summary Get tasks by project ID
// @Description Retrieve tasks for a specific project
// @Tags Task
// @Security ApiKeyAuth
// @Produce json
// @Param project_id path int true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /projects/{project_id}/tasks [get]
func GetTasksByProjectID(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}
	tasks, err := services.GetTasksByProjectID(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// get tasks by assignee id
// GetTasksByAssigneeID godoc
// @Summary Get tasks by assignee ID
// @Description Retrieve tasks assigned to a user
// @Tags Task
// @Security ApiKeyAuth
// @Produce json
// @Param assignee_id path int true "Assignee ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /tasks/assignee/{assignee_id}/ [get]
func GetTasksByAssigneeID(c *gin.Context) {
	assigneeID, err := strconv.Atoi(c.Param("assigneeId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid assignee ID"})
		return
	}
	tasks, err := services.GetTasksByAssigneeID(assigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
