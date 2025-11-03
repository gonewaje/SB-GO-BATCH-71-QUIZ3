package controllers

import (
	"database/sql"
	"library/repository"
	"library/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	DB *sql.DB
}

func (a AuthController) Login(c *gin.Context) {
	var req structs.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := repository.GetUserByUsername(a.DB, req.Username)
	if err != nil || user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
