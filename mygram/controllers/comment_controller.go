package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mygram/model"
)

func GetComments(c *gin.Context) {
	photoID := c.Param("id")
	comments, err := model.GetComments(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
