package repositories

import (
	"errors"
	"mygram/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) Authenticate(email, password string) (*models.User, error) {
	user := new(models.User)
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}
