package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) FindAllByPhotoID(photoID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("photo_id = ?", photoID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	if err := r.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}
