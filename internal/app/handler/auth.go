package handler

import (
	"log"
	"net/http"

	"github.com/XRS0/Sigma-Network/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input repository.User

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		log.Printf("[ERROR] failed to bind User: %s", err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Printf("[ERROR] failed to create user: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{"id": id})
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		log.Printf("[ERROR] failed to bind SignInInput: %s", err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Printf("[ERROR] failed to generate token: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{"token": token})
}
