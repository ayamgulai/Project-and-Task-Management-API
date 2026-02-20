package controllers

import (
	"mini-jira-backend/models"
	"mini-jira-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.LoginUserInput true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var req models.LoginUserInput

	// bind & validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email and password are required",
		})
		return
	}

	// call service
	token, err := services.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}

// register user
// Register godoc
// @Summary Register new user
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param body body models.RegisterUserInput true "User payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {

	var req models.User
	// bind & validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name, email, and password are required",
			"details": err.Error(),
		})
		return
	}

	// call service
	createdUser, err := services.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user":    createdUser,
	})
}
