package services

import (
	"context"

	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/Rikypurnomo/warmup/pkg/util"
)

func (s *ServicessInit) Register(ctx context.Context, user *models.User) error {
	hash, err := util.Generatehash(user.Password)
	err = s.RepositoryAuth.Register(ctx, &models.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: hash,
	})

	if err != nil {
		return err
	}
	return nil

}
