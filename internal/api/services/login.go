package services

import (
	"context"

	authdto "github.com/Rikypurnomo/warmup/internal/api/dto/auth"
	"github.com/Rikypurnomo/warmup/internal/api/models"
	"github.com/Rikypurnomo/warmup/internal/api/validator"
	"github.com/Rikypurnomo/warmup/pkg/util"
)

func (s *ServicessInit) Login(ctx context.Context, loginRequest *models.Login) (*authdto.LoginTokenResponse, error) {
	user, err := s.RepositoryAuth.Login(ctx, loginRequest.Email)
	if err != nil {
		return nil, err
	}
	res, err := util.ComparePassword(user.Password, loginRequest.Password)

	if res == false {
		return nil, err
	}
	token, err := validator.GenerateToken(user.Email, int(user.ID))

	if err != nil {
		return nil, err
	}

	err = validator.CreateAuth(ctx, user.ID, &token, &user)

	return &token, nil
}
