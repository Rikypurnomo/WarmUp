package services

import (
	"context"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
)

func (s *ServicessInit) CreateTransaction(ctx context.Context, UserID int) error {
	err := s.RepositoryTransaction.CreateTransaction(ctx, UserID)

	if err != nil {
		return err
	}

	return nil

}

func (s *ServicessInit) GetTransactionHistory(ctx context.Context, userID int) ([]dto.History, error) {
	history, err := s.RepositoryTransaction.GetTransactionHistory(ctx, userID)
	if err != nil {
		return nil, err
	}

	return history, nil
}

// func (s *ServicessInit) GetTransactionHistory(ctx context.Context, userID int) ([]models.Transaction, error) {
// 	transactions, err := s.RepositoryTransaction.GetTransactionHistory(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return transactions, nil
// }
