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
func (s *Service) CreateUser(ctx context.Context, user entity.CreateUser) (int, error) {
	return s.storage.CreateUser(ctx, user)
}

func (s *Service) FindUserBydID(ctx context.Context, userID int) (entity.User, error) {
	return s.storage.FindUserBydID(ctx, userID)
}

func (s *Service) UpdateUser(ctx context.Context, bannerID int, user entity.User) error {
	return s.storage.UpdateUser(ctx, bannerID, user)
}

func (s *Service) DeleteUser(ctx context.Context, bannerID int) error {
	return s.storage.DeleteUser(ctx, bannerID)
}
