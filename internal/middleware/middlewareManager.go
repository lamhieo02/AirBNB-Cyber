package middleware

import (
	"context"
	"go01-airbnb/config"
	usermodel "go01-airbnb/internal/user/model"
)

type UserRepository interface {
	FindDataWithCondition(ctx context.Context, condition map[string]any) (*usermodel.User, error)
}

type middlewareManager struct {
	cfg      *config.Config
	userRepo UserRepository
}

func NewMiddlewareManager(cfg *config.Config, userRepo UserRepository) *middlewareManager {
	return &middlewareManager{cfg: cfg, userRepo: userRepo}
}
