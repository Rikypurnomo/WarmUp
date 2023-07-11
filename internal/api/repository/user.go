package repository

import (
	"context"

	"github.com/Rikypurnomo/warmup/internal/api/models"
	"gorm.io/gorm"
)

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

type (
	UsersRepository interface {
		Register(ctx context.Context, user *models.User) (err error)
		Login(ctx context.Context, email string) (models.User, error)
	}

	userRepository struct {
		DB *gorm.DB
	}
)

func (r *userRepository) Register(ctx context.Context, user *models.User) (err error) {
	err = r.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Login(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).First(&user, "email=?", email).Error

	return user, err
}
