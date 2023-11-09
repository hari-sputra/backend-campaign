package handler

import (
	"backend-campaign/helper"
	"backend-campaign/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewsUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := user.FormatUser(newUser, "1h23u123dy72hd71dgu17")

	response := helper.ApiResponse("Account has been Created", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := user.FormatUser(loginUser, "12h3881h318hd8182hdh")

	response := helper.ApiResponse("Logged successfully", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)

}
