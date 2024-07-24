package repository

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/entity"
)

type UserRepository interface {
	CreateBanner(ctx context.Context, user entity.User) (int, error)
}
