package services

import (
	"errors"
	"log"
	"mini-jira-backend/models"
	"mini-jira-backend/repositories"
)

// get all projects
func GetProjects() ([]models.Project, error) {
	return repositories.GetProjects()
}

// get project by id
func GetProjectByID(id int) (*models.Project, error) {
	project, err := repositories.GetProjectByID(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// create new project
func CreateProject(project models.Project, userID int, role string) (*models.Project, error) {
	log.Println("owner_id:", userID)
	// Ensure the acting user exists
	user, err := repositories.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("owner not found")
	}

	createdProject, err := repositories.CreateProject(&project, userID, role)
	if err != nil {
		return nil, err
	}
	return createdProject, nil
}

// update project by id
func UpdateProject(id int, project models.Project) (*models.Project, error) {
	updatedProject, err := repositories.UpdateProject(id, &project)
	if err != nil {
		return nil, err
	}
	return updatedProject, nil
}
