package services

import (
	"context"
)

func (s *ServicessInit) GetBalance(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{
		"balance": 0,
	}, nil
}
