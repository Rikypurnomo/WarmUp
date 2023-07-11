package repository

import (
	"context"
	"database/sql"

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
		GetTransactionHistory(ctx context.Context, userID int) ([]models.Transaction, error)
	}

	transactionRepository struct {
		DB *gorm.DB
	}
)

func (r *transactionRepository) CreateTransaction(ctx context.Context, userID int) (err error) {
	tx := r.DB.Begin()

	// var productID int
	// err = tx.WithContext(ctx).
	// 	Model(&models.Cart{}).
	// 	Select("products.id").
	// 	Joins("JOIN products ON carts.product_id = products.id").
	// 	Where("carts.user_id = ?", userID).
	// 	Scan(&productID).
	// 	Error
	// fmt.Println(productID, "<(&&&&&&&&&&&&&&&&&&&&&&&&**(((((PRODUCTID")

	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// var totalPrice sql.NullFloat64
	// err = tx.WithContext(ctx).
	// 	Model(&models.Product{}).
	// 	Select("COALESCE(SUM(products.price * carts.quantity), 0)").
	// 	Joins("JOIN carts ON carts.product_id = products.id").
	// 	Where("carts.user_id = ? OR products.deleted_at IS NULL", userID).
	// 	Scan(&totalPrice).
	// 	Error

	var totalPrice sql.NullFloat64
	err = tx.WithContext(ctx).
		Model(&models.Cart{}).
		Select("COALESCE(SUM(products.price * carts.quantity), 0)").
		Joins("JOIN products ON carts.product_id = products.id").
		Where("carts.user_id = ?", userID).
		Scan(&totalPrice).
		Error

	// var totalPrice sql.NullInt64
	// err = tx.WithContext(ctx).
	// 	Model(&models.Product{}).
	// 	Select("SUM(products.price)").
	// 	Joins("JOIN carts ON carts.product_id = products.id").
	// 	Where("carts.user_id = ? OR products.deleted_at IS NULL", userID).
	// 	Scan(&totalPrice).
	// 	Error

	if err != nil {
		tx.Rollback()
		return err
	}
	transaction := &models.Transaction{
		UserID: userID,
		// ProductID:  productID,
		TotalPrice: int(totalPrice.Float64),
	}

	err = tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}

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

// history := &models.History{
// 	TransactionID: int(transaction.ID),
// 	TotalPrice:    transaction.TotalPrice,
// 	UserID:        transaction.UserID,
// 	ProductID:     productID,
// }
// fmt.Println(history, "<<<ini yg historyyyyyyyyyy")

func (r *transactionRepository) GetTransactionHistory(ctx context.Context, userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("user_id = ?", userID).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// var Transactions []models.Transaction
// err = r.DB.WithContext(ctx).
// 	Model(&models.Transaction{}).
// 	Select("user_id, product_id, total_price").
// 	Where("user_id = ?", userID).
// 	Find(&Transactions).
// 	Error

// fmt.Println(userID, "ini userrrrrrrrrr")
// fmt.Println("<history>", Transactions, "<>history<<")
