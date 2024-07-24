package service

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	UserBydID(ctx context.Context, userID int) (entity.User, error)
	UpdateUser(ctx, bannerID int, user entity.User)
	DeleteUser(ctx, bannerID int)
}
