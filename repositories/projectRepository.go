package repositories

import (
	"database/sql"
	"mini-jira-backend/configs"
	"mini-jira-backend/models"
)

// GET all projects
func GetProjects() ([]models.Project, error) {
	rows, err := configs.DB.Query(`
		SELECT id, name, description, owner_id, created_at
		FROM projects
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.OwnerID,
			&project.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

// get project by id
func GetProjectByID(id int) (*models.Project, error) {
	row := configs.DB.QueryRow(`
		SELECT id, name, description, owner_id, created_at
		FROM projects WHERE id = $1
	`, id)
	var project models.Project
	err := row.Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.OwnerID,
		&project.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if no project found
		}
		return nil, err
	}
	return &project, nil
}

// create new project
func CreateProject(project *models.Project, userID int, role string) (*models.Project, error) {
	project.OwnerID = userID
	row := configs.DB.QueryRow(`
		INSERT INTO projects (name, description, owner_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`, project.Name, project.Description, project.OwnerID)
	err := row.Scan(&project.ID, &project.CreatedAt)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// update project
func UpdateProject(projectID int, project *models.Project) (*models.Project, error) {
	row := configs.DB.QueryRow(`
		UPDATE projects
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id, name, description, owner_id, created_at
	`, project.Name, project.Description, projectID)
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.OwnerID, &project.CreatedAt)
	if err != nil {
		return nil, err
	}
	return project, nil
}
