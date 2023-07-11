package services

import (
	"context"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
)

func (s *ServicessInit) AddToCart(ctx context.Context, productID int, id int) error {
	err := s.RepositoryCart.AddToCart(ctx, productID, id)

	if err != nil {
		return err
	}
	return nil

}

func (s *ServicessInit) GetCartByUserID(ctx context.Context, userID int) ([]dto.CartResponse, error) {
	carts, err := s.RepositoryCart.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return carts, nil
}
