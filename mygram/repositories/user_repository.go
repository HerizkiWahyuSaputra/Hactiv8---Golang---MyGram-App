package repositories

import (
	"errors"
	"fmt"
	"mygram/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindById(id uint) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User

	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("User with username %s not found", username)
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("User with email %s not found", email)
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindById(id uint) (models.User, error) {
	var user models.User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("User with ID %d not found", id)
		}
		return models.User{}, err
	}

	return user, nil
}
