package repository

import (
	"context"
	"fmt"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"gorm.io/gorm"
)

func NewCategoriesRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

type (
	CategoriesRepository interface {
		CreateCategory(ctx context.Context, category *models.Category) (err error)
		ListCategory(ctx context.Context, page int, limit int, search string) (category []dto.CategoryResponse, count int64, err error)
	}

	categoryRepository struct {
		DB *gorm.DB
	}
)

func (r *categoryRepository) CreateCategory(ctx context.Context, category *models.Category) (err error) {
	err = r.DB.WithContext(ctx).Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) ListCategory(ctx context.Context, page int, limit int, search string) (category []dto.CategoryResponse, count int64, err error) {
	offset := (page - 1) * limit

	err = r.DB.WithContext(ctx).
		Model(&models.Category{}).
		Select("categories.name,categories.id").
		Where("deleted_at IS NULL AND name LIKE ?", "%"+search+"%").
		Offset(offset).Limit(limit).
		Find(&category).
		Count(&count).
		Error

	if err != nil {
		return category, count, err
	}
	fmt.Println(category, "<CATEGORY>")

	return category, count, nil
}

// func (r *categoryRepository) ListCategory(ctx context.Context, page int, limit int, search string) (category []dto.CategoryResponse, count int64, err error) {

// 	offset := (page - 1) * limit
// 	err = r.DB.WithContext(ctx).
// 		Model(&models.Category{}).
// 		Select("categories.name").
// 		Where("categories.deleted_at IS NULL AND categoriess.name LIKE ?", "%"+search+"%").Offset(offset).Limit(limit).
// 		Find(&category).Count(&count).Error

// 	if err != nil {
// 		return category, count, err
// 	}
// 	return category, count, nil
// }
