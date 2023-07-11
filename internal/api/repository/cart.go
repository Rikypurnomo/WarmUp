package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"gorm.io/gorm"
)

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{
		DB: db,
	}
}

type (
	CartsRepository interface {
		AddToCart(ctx context.Context, productID, userID int) (err error)
		GetCartByUserID(ctx context.Context, userID int) ([]dto.CartResponse, error)
	}

	cartRepository struct {
		DB *gorm.DB
	}
)

func (r *cartRepository) AddToCart(ctx context.Context, productID, userID int) (err error) {
	var cart models.Cart
	err = r.DB.WithContext(ctx).
		Model(&models.Cart{}).
		Where("product_id = ? AND user_id = ?", productID, userID).
		First(&cart).
		Error

	if err != nil {
		cart.Quantity = 1
		cart.UserID = userID
		cart.ProductID = productID
		err = r.DB.WithContext(ctx).Create(&cart).Error
	} else {
		cart.Quantity++
		err = r.DB.WithContext(ctx).Save(&cart).Error
	}

	fmt.Println(cart.Quantity, "<iniQQ>>>>")

	if err != nil {
		return err
	}

	return nil
}

func (r *cartRepository) GetCartByUserID(ctx context.Context, userID int) ([]dto.CartResponse, error) {

	var carts []dto.CartResponse
	err := r.DB.WithContext(ctx).
		Model(&models.Cart{}).
		Select("products.name,carts.quantity, users.full_name as full_name, categories.name as category_name").
		Joins("JOIN products ON carts.product_id = products.id").
		Joins("JOIN users ON carts.user_id = users.id").
		Joins("JOIN categories ON products.category_id = categories.id").
		Where("carts.user_id = ? AND carts.status = ?", userID, false).
		Find(&carts).
		Error
	if err != nil {
		return nil, err
	}

	if len(carts) == 0 {
		return nil, errors.New("No carts found")
	}
	fmt.Println(carts, "<(*&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&)")

	return carts, nil
}
