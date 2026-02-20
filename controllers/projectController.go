package controllers

import (
	"mini-jira-backend/models"
	"mini-jira-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get all projects
// GetProjects godoc
// @Summary Get all projects
// @Description Retrieve list of projects
// @Tags Project
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects [get]
func GetProjects(c *gin.Context) {
	projects, err := services.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// get project by id
// GetProjectByID godoc
// @Summary Get project by ID
// @Description Get single project by its ID
// @Tags Project
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /projects/{id} [get]
func GetProjectByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}
	project, err := services.GetProjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}

// create new project
// CreateProject godoc
// @Summary Create new project
// @Description Create a project
// @Tags Project
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.CreateProjectInput true "Project payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects [post]
func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetInt("user_id")
	role := c.GetString("role")
	createdProject, err := services.CreateProject(project, userID, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"project": createdProject})
}

// update project by id
// UpdateProject godoc
// @Summary Update project
// @Description Update project by ID
// @Tags Project
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param body body models.UpdateProjectInput true "Project payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /projects/{id} [put]
func UpdateProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}
	userID := c.GetInt("user_id")
	role := c.GetString("role")
	projectData, err := services.GetProjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	if role != "admin" && projectData.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: not project owner"})
		return
	}

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedProject, err := services.UpdateProject(id, project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": updatedProject})
}
