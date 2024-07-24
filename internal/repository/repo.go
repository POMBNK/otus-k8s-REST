package repository

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.CreateUser) (int, error)
	FindUserBydID(ctx context.Context, userID int) (entity.User, error)
	UpdateUser(ctx context.Context, bannerID int, user entity.User) error
	DeleteUser(ctx context.Context, userID int) error
}
