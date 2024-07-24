package user

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/entity"
	"github.com/POMBNK/kuberRest/internal/repository"
	"github.com/POMBNK/kuberRest/internal/repository/tx"
)

type Service struct {
	storage repository.UserRepository
	tx      tx.Tx
}

func New(storage repository.UserRepository, tx tx.Tx) *Service {
	return &Service{
		storage: storage,
		tx:      tx,
	}
}
func (s Service) CreateUser(ctx context.Context, user entity.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UserBydID(ctx context.Context, userID int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateUser(ctx, bannerID int, user entity.User) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteUser(ctx, bannerID int) {
	//TODO implement me
	panic("implement me")
}
