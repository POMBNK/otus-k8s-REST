package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/POMBNK/kuberRest/internal/entity"
	"github.com/POMBNK/kuberRest/pkg/client/postgres"
)

type Storage struct {
	conn postgres.Client
}

func New(conn postgres.Client) *Storage {
	return &Storage{conn: conn}
}

func (s *Storage) CreateUser(ctx context.Context, user entity.CreateUser) (int, error) {

	insertBuilder := sq.Insert("users").PlaceholderFormat(sq.Dollar).
		Columns("last_name", "first_name", "email", "phone", "username").
		Values(user.LastName, user.FirstName, user.Email, user.Phone, user.Username).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var userID int
	err = s.conn.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *Storage) FindUserBydID(ctx context.Context, userID int) (entity.User, error) {

	selectBuilder := sq.Select("last_name", "first_name", "email", "phone", "username").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return entity.User{}, err
	}

	var user entity.User
	err = s.conn.QueryRow(ctx, query, args...).Scan(&user.LastName, &user.FirstName, &user.Email, &user.Phone, &user.Username)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *Storage) UpdateUser(ctx context.Context, bannerID int, user entity.User) error {

	updateBuilder := sq.Update("users").PlaceholderFormat(sq.Dollar).
		Set("last_name", user.LastName).
		Set("first_name", user.FirstName).
		Set("email", user.Email).
		Set("phone", user.Phone).
		Set("username", user.Username).
		Where(sq.Eq{"id": bannerID})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = s.conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (s *Storage) DeleteUser(ctx context.Context, userID int) error {

	deleteBuilder := sq.Delete("users").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = s.conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
