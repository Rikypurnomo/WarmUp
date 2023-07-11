package services

import (
	"context"

	"github.com/Rikypurnomo/warmup/internal/api/dto"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/Rikypurnomo/warmup/pkg/util"
	models_server "github.com/Rikypurnomo/warmup/server/models"
)

func (s *ServicessInit) CreateCategory(ctx context.Context, category *models.Category) error {
	err := s.RepositoryCategori.CreateCategory(ctx, &models.Category{
		Name: category.Name,
	})

	if err != nil {
		return err
	}
	return nil

}

func (s *ServicessInit) ListCategory(ctx context.Context, page int, limit int, search string) ([]dto.CategoryResponse, models_server.MetaPagination, error) {
	category, count, err := s.RepositoryCategori.ListCategory(ctx, page, limit, search)

	if err != nil {
		return nil, util.ResPagination(0, int64(page), int64(limit)), err
	}
	return category, util.ResPagination(count, int64(page), int64(limit)), nil
}
