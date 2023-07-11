package services

import (
	"context"
)

func (s *ServicessInit) AddToCart(ctx context.Context, productID int, id int) error {
	err := s.RepositoryCart.AddToCart(ctx, productID, id)

	if err != nil {
		return err
	}
	return nil

}
