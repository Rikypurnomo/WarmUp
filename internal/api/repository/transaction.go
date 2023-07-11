package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"gorm.io/gorm"
)

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{
		DB: db,
	}
}

type (
	TransactionsRepository interface {
		CreateTransaction(ctx context.Context, userID int) (err error)
		// GetTransactionHistory(ctx context.Context) ([]models.Transaction, error)
		GetTransactionHistory(ctx context.Context, userID int) ([]dto.History, error)
	}

	transactionRepository struct {
		DB *gorm.DB
	}
)

func (r *transactionRepository) CreateTransaction(ctx context.Context, userID int) (err error) {
	tx := r.DB.Begin()

	var totalPrice sql.NullFloat64
	err = tx.WithContext(ctx).
		Model(&models.Cart{}).
		Select("COALESCE(SUM(products.price * carts.quantity), 0)").
		Joins("JOIN products ON carts.product_id = products.id").
		Where("carts.user_id = ?", userID).
		Scan(&totalPrice).
		Error

	if err != nil {
		tx.Rollback()
		return err
	}
	transaction := &models.Transaction{
		UserID:     userID,
		TotalPrice: int(totalPrice.Float64),
	}

	err = tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var cart int64
	err = tx.WithContext(ctx).
		Model(&models.Cart{}).
		Where("user_id = ? AND (status = false)", userID).
		Count(&cart).
		Error
	if err != nil {
		return err
	}

	fmt.Println(cart, "<<hasil count")
	if cart == 0 {
		return errors.New("User already has a checked-out cart or existing transaction")
	}

	err = tx.WithContext(ctx).
		Model(&models.Cart{}).
		Where("user_id = ? AND status = 'false'", userID).
		Update("status", "true").
		Error
	if err != nil {
		return err
	}

	history := &models.History{
		UserID:        userID,
		TransactionID: int(transaction.ID),
	}

	err = tx.Create(&history).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil
	}

	return nil
}

func (r *transactionRepository) GetTransactionHistory(ctx context.Context, userID int) ([]dto.History, error) {

	var history []dto.History
	err := r.DB.WithContext(ctx).
		Model(&models.Product{}).
		Select("products.name AS name, users.full_name AS full_name, transactions.id AS transaction_id, categories.name AS category_name").
		Joins("JOIN carts ON carts.product_id = products.id").
		Joins("JOIN categories ON categories.id = products.category_id").
		Joins("JOIN users ON carts.user_id = users.id").
		Joins("JOIN transactions ON transactions.user_id = users.id").
		Where("transactions.user_id = ?", userID).
		Find(&history).Error

	if err != nil {
		return nil, err
	}
	fmt.Println(history, "<(*&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&)")

	return history, nil
}
