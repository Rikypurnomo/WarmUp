package repository

import (
	"context"
	"fmt"

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

// func (r *cartRepository) AddToCart(ctx context.Context, productID, id int) (err error) {

// 	var quantity int
// 	err = r.DB.WithContext(ctx).
// 		Model(&models.Cart{}).
// 		Select("quantity").
// 		Find(&quantity).
// 		Error

// 		fmt.Println(quantity,"<QUANTITY>")

// 	if err != nil {
// 		return err
// 	}

// 	cart := &models.Cart{
// 		UserID:    id,
// 		ProductID: productID,
// 		Status:    false,
// 		Quantity: quantity,
// 	}

// 	err = r.DB.WithContext(ctx).Create(&cart).Error

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
