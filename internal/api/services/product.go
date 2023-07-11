package services

import (
	"context"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/Rikypurnomo/warmup/pkg/util"
	models_server "github.com/Rikypurnomo/warmup/server/models"
)

func (s *ServicessInit) ListProduct(ctx context.Context, page int, limit int, search string) ([]dto.ProductResponse, models_server.MetaPagination, error) {
	product, count, err := s.RepositoryProduct.ListProduct(ctx, page, limit, search)

	if err != nil {
		return nil, util.ResPagination(0, int64(page), int64(limit)), err
	}
	return product, util.ResPagination(count, int64(page), int64(limit)), nil
}

func (s *ServicessInit) DeleteProduct(ctx context.Context, productID int) error {

	err := s.RepositoryProduct.DeleteProduct(ctx, productID)

	if err != nil {
		return err
	}
	return nil

}

func (s *ServicessInit) CreateProduct(ctx context.Context, product *models.Product) error {
	err := s.RepositoryProduct.CreateProduct(ctx, &models.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		CategoryID:  product.CategoryID,
	})

	if err != nil {
		return err
	}
	return nil

}
