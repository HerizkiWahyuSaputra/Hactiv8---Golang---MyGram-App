package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mygram/model"
)

func GetPhotos(c *gin.Context) {
	photos, err := model.GetPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func CreatePhoto(c *gin.Context) {
	var photo model.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.CreatePhoto(&photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

func GetPhotoByID(c *gin.Context) {
	photoID := c.Param("id")
	photo, err := model.GetPhotoByID(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("id")
	var photo model.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.UpdatePhoto(photoID, &photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Param("id")
	if err := model.DeletePhoto(photoID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	
