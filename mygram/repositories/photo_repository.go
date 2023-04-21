package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{db}
}

func (r *PhotoRepository) FindAll() ([]models.Photo, error) {
	var photos []models.Photo
	if err := r.db.Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func (r *PhotoRepository) FindByID(id uint) (*models.Photo, error) {
	photo := new(models.Photo)
	if err := r.db.Where("id = ?", id).First(photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *PhotoRepository) Create(photo *models.Photo) error {
	if err := r.db.Create(photo).Error; err != nil {
		return err
	}
	return nil
}

func (r *PhotoRepository) Update(photo *models.Photo) error {
	if err := r.db.Save(photo).Error; err != nil {
		return err
	}
	return nil
}

func (r *PhotoRepository) Delete(photo *models.Photo) error {
	if err := r.db.Delete(photo).Error; err != nil {
		return err
	}
	return nil
}
