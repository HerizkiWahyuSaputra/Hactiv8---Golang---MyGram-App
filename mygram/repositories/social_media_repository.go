package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{db}
}

func (r *SocialMediaRepository) FindAllByPhotoID(photoID uint) ([]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	if err := r.db.Where("photo_id = ?", photoID).Find(&social
