package services

import (
	"errors"
	"fmt"
	"log"
	"mini-jira-backend/models"
	"mini-jira-backend/repositories"
)

// get all tasks
func GetTasks() ([]models.Task, error) {
	return repositories.GetTasks()
}

// get task by id
func GetTaskByID(id int) (*models.Task, error) {
	task, err := repositories.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// create new task
func CreateTask(task models.Task, userEmail string) (*models.Task, error) {
	log.Println("userEmail:", userEmail)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil || user == nil {
		return nil, err
	}
	log.Println("user.ID:", user.ID)
	createdTask, err := repositories.CreateTask(task, user.ID)
	if err != nil {
		return nil, errors.New("error while inserting task")
	}
	log := &models.TaskLog{
		TaskID:    createdTask.ID,
		Action:    "create_task",
		OldValue:  "none",
		NewValue:  fmt.Sprintf("%s created", createdTask.Title),
		ChangedBy: userEmail,
	}
	_, err = repositories.CreateTaskLog(log)
	if err != nil {
		return nil, err
	}
	return createdTask, nil
}

// delete task by id
func DeleteTask(id int, userEmail string, role string, userId int) (isSuccess bool, err error) {
	// 1️⃣ Ambil task lama
	task, err := repositories.GetTaskByID(id)
	if err != nil || task == nil {
		return false, err
	}
	if role == "member" {
		// Member hanya boleh task miliknya
		if task.AssigneeID == nil || *task.AssigneeID != user.ID {
			return false, errors.New("forbidden: not your task")
		}
	}
	log := &models.TaskLog{
		TaskID:    id,
		Action:    "delete_task",
		OldValue:  fmt.Sprintf("%s exists", task.Title),
		NewValue:  "deleted",
		ChangedBy: userEmail,
	}
	_, err = repositories.CreateTaskLog(log)
	if err != nil {
		return false, err
	}
	isSuccess, err = repositories.DeleteTask(id)
	if err != nil {
		return false, err
	}

	return isSuccess, nil
}

// get tasks by project id
func GetTasksByProjectID(projectID int) ([]models.Task, error) {
	return repositories.GetTasksByProjectID(projectID)
}

// update task status by id
func UpdateTaskStatus(id int, status string, userEmail string, role string) (*models.TaskLog, error) {
	// 1️⃣ Ambil task lama
	task, err := repositories.GetTaskByID(id)
	if err != nil || task == nil {
		return nil, err
	}
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		return nil, err
	}
	if role == "member" {
		// Member hanya boleh task miliknya
		if task.AssigneeID == nil || *task.AssigneeID != user.ID {
			return nil, errors.New("forbidden: not your task")
		}
	}

	oldStatus := task.Status

	// 2️⃣ Validasi business rule
	if oldStatus == status {
		return nil, errors.New("status is the same")
	}

	//  validasi transition
	validTransitions := map[string][]string{
		"todo":        {"in_progress"},
		"in_progress": {"in_review", "todo"},
		"in_review":   {"in_progress", "done"},
	}

	isValid := false
	for _, v := range validTransitions[oldStatus] {
		if v == status {
			isValid = true
			break
		}
	}
	if !isValid {
		return nil, errors.New("invalid status transition")
	}
	log := &models.TaskLog{
		TaskID:    id,
		Action:    "update_status",
		OldValue:  oldStatus,
		NewValue:  status,
		ChangedBy: userEmail,
	}
	logUpdate, err := repositories.CreateTaskLog(log)
	_, err = repositories.UpdateTaskStatus(id, status)
	if err != nil {
		return nil, err
	}
	return &logUpdate, nil
}

// get tasks by assignee id
func GetTasksByAssigneeID(assigneeID int) ([]models.Task, error) {
	return repositories.GetTasksByAssigneeID(assigneeID)
}
