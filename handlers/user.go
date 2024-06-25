package handlers

import (
	"net/http"
	"proyek/auth"
	"proyek/helper"
	"proyek/user"
	"strconv"

	"github.com/gin-gonic/gin"

)

type userHandler struct {
	UserService user.UserService
	AuthService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}


func (uh *userHandler) CreateUser(c *gin.Context) {
	var input user.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	newUser, err := uh.UserService.Register(input)
	if err != nil {
		if err.Error() == "username already taken" {
			c.JSON(http.StatusBadRequest, helper.APIResponse("Username already taken", http.StatusBadRequest, "error", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to create user", http.StatusInternalServerError, "error", nil))
		return
	}

	token, err := uh.AuthService.GenerateToken(newUser.IdUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to generate token", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("User created successfully", http.StatusCreated, "success", user.FormatUser(newUser, token))
	c.JSON(http.StatusCreated, response)
}

func (uh *userHandler) UpdateUser(c *gin.Context) {
	var input user.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid user ID", http.StatusBadRequest, "error", nil))
		return
	}

	updatedUser, err := uh.UserService.UpdateUser(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to update user", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("User updated successfully", http.StatusOK, "success", user.FormatUser(updatedUser, ""))
	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors))
		return
	}

	loggedInUser, err := uh.UserService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.APIResponse("Invalid credentials", http.StatusUnauthorized, "error", nil))
		return
	}

	token, err := uh.AuthService.GenerateToken(loggedInUser.IdUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to generate token", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("Login successful", http.StatusOK, "success", user.FormatUser(loggedInUser, token))
	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) FetchUser(c *gin.Context){
	currentUser:= c.MustGet("currentUser").(user.UserTable)

	formatter := user.FormatUser(currentUser,"")

	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}


func (uh *userHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid user ID", http.StatusBadRequest, "error", nil))
		return
	}

	userData, err := uh.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("User not found", http.StatusInternalServerError, "error", nil))
		return
	}

	response := helper.APIResponse("User details", http.StatusOK, "success", user.FormatUser(userData,""))
	c.JSON(http.StatusOK, response)
}
