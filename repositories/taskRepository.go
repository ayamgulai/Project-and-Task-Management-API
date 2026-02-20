package repositories

import (
	"database/sql"
	"mini-jira-backend/configs"
	"mini-jira-backend/models"
)

// GET all tasks
func GetTasks() ([]models.Task, error) {
	rows, err := configs.DB.Query(`
		SELECT id, project_id, title, description, status, priority, assignee_id, created_at
		FROM tasks
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.AssigneeID,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// get task by id
func GetTaskByID(id int) (*models.Task, error) {
	row := configs.DB.QueryRow(`
		SELECT id, project_id, title, description, status, priority, assignee_id, created_at
		FROM tasks WHERE id = $1
	`, id)

	var task models.Task
	err := row.Scan(
		&task.ID,
		&task.ProjectID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&task.AssigneeID,
		&task.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func CreateTask(task models.Task, userID int) (*models.Task, error) {
	task.Status = "todo"
	task.AssigneeID = &userID
	query := `
	INSERT INTO tasks (
	project_id, title, description, status, priority, assignee_id
	) 
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id, project_id, title, description, status, priority, assignee_id, created_at`

	var createdTask models.Task

	err := configs.DB.QueryRow(query,
		task.ProjectID,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.AssigneeID,
	).Scan(
		&createdTask.ID,
		&createdTask.ProjectID,
		&createdTask.Title,
		&createdTask.Description,
		&createdTask.Status,
		&createdTask.Priority,
		&createdTask.AssigneeID,
		&createdTask.CreatedAt,
	)

	return &createdTask, err
}

func DeleteTask(id int) (bool, error) {
	res, err := configs.DB.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return false, err
	}

	affected, _ := res.RowsAffected()
	return affected > 0, nil
}

func GetTasksByProjectID(projectID int) ([]models.Task, error) {
	rows, err := configs.DB.Query(`
		SELECT id, project_id, title, description, status, priority, assignee_id, created_at
		FROM tasks
		WHERE project_id = $1
	`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.AssigneeID,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// get tasks by assignee id
func GetTasksByAssigneeID(assigneeID int) ([]models.Task, error) {
	rows, err := configs.DB.Query(`
		SELECT id, project_id, title, description, status, priority, assignee_id, created_at
		FROM tasks
		WHERE assignee_id = $1
	`, assigneeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.AssigneeID,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// update task status
func UpdateTaskStatus(id int, status string) (*models.Task, error) {
	query := `
	UPDATE tasks
	SET status = $1
	WHERE id = $2
	RETURNING id, project_id, title, description, status, priority, assignee_id, created_at`

	var updatedTask models.Task
	err := configs.DB.QueryRow(query, status, id).Scan(
		&updatedTask.ID,
		&updatedTask.ProjectID,
		&updatedTask.Title,
		&updatedTask.Description,
		&updatedTask.Status,
		&updatedTask.Priority,
		&updatedTask.AssigneeID,
		&updatedTask.CreatedAt,
	)

	return &updatedTask, err
}
