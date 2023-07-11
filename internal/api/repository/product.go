package repository

import (
	"context"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"gorm.io/gorm"
)

func NewProductsRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		DB: db,
	}
}

type (
	ProductsRepository interface {
		DeleteProduct(ctx context.Context, productID int) error
		ListProduct(ctx context.Context, page int, limit int, search string) (product []dto.ProductResponse, count int64, err error)
		CreateProduct(ctx context.Context, product *models.Product) (err error)
	}

	productRepository struct {
		DB *gorm.DB
	}
)

func (r *productRepository) ListProduct(ctx context.Context, page int, limit int, search string) (product []dto.ProductResponse, count int64, err error) {

	offset := (page - 1) * limit
	err = r.DB.WithContext(ctx).
		Model(&models.Product{}).
		Select("products.name AS name,products.id ,products.price AS price ,products.description AS description , categories.name as category_name").
		Joins("JOIN categories ON products.category_id = categories.id").
		Where("products.deleted_at IS NUll AND products.name iliKE ?", "%"+search+"%").Offset(offset).Limit(limit).
		// Where("name LIKE ?",cari).
		Find(&product).Count(&count).Error

	if err != nil {
		return product, count, err
	}
	return product, count, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, productID int) error {
	err := r.DB.WithContext(ctx).
		Delete(&models.Product{}, "id = ?", productID).
		Error

	return err
}

func (r *productRepository) CreateProduct(ctx context.Context, product *models.Product) (err error) {
	err = r.DB.WithContext(ctx).Create(&product).Error

	if err != nil {
		return err
	}
	return nil
}
