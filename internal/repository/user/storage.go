package user

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/entity"
	"github.com/POMBNK/kuberRest/pkg/client/postgres"
)

type Storage struct {
	conn postgres.Client
}

func New(conn postgres.Client) *Storage {
	return &Storage{conn: conn}
}

func (s *Storage) CreateBanner(ctx context.Context, banner entity.User) (int, error) {
	//jsonb, err := json.Marshal(banner.Content)
	//if err != nil {
	//	return 0, err
	//}
	//insertBuilder := sq.Insert("banners").PlaceholderFormat(sq.Dollar).
	//	Columns("content", "feature_id", "is_active").
	//	Values(jsonb, banner.FeatureId, banner.IsActive).
	//	Suffix("RETURNING id")
	//
	//query, args, err := insertBuilder.ToSql()
	//if err != nil {
	//	return 0, err
	//}
	//
	//var bannerID int
	//err = s.conn.QueryRow(ctx, query, args...).Scan(&bannerID)
	//if err != nil {
	//	return 0, err
	//}
	//
	//insertLinkBuilder := sq.Insert("banner_tags").PlaceholderFormat(sq.Dollar).
	//	Columns("banner_id", "tag_id")
	//
	//for _, tagID := range banner.TagIds {
	//	insertLinkBuilder = insertLinkBuilder.
	//		Values(bannerID, tagID)
	//}
	//
	//query, args, err = insertLinkBuilder.ToSql()
	//if err != nil {
	//	return 0, err
	//}
	//
	//_, err = s.conn.Exec(ctx, query, args...)
	//if err != nil {
	//	return 0, err
	//}

	return 0, nil
}

func (s *Storage) UpdateBanner() {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) DeleteBanner() {
	//TODO implement me
	panic("implement me")
}
