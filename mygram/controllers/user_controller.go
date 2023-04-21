package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/muchlist/mygram-go/models"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func (ctrl UserController) Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Name cannot be empty",
		})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email cannot be empty",
		})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Password cannot be empty",
		})
		return
	}

	user.Password, err = ctrl.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to hash password",
		})
		return
	}

	err = models.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create user",
		})
		return
	}

	token, err := ctrl.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func (ctrl UserController) Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	var dbUser models.User
	err = models.DB.Where("email = ?", user.Email).First(&dbUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid email or password",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to query database",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid email or password",
		})
		return
	}

	token, err := ctrl.GenerateToken(dbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
