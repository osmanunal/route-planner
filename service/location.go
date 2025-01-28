package service

import (
	"context"
	"route-planner/model"

	"github.com/uptrace/bun"
)

type ILocationService interface {
	GetAll(ctx context.Context) ([]model.Location, int, error)
	GetByID(ctx context.Context, id int64) (model.Location, error)
	Create(ctx context.Context, m *model.Location) error
	Update(ctx context.Context, m model.Location) error
	Delete(ctx context.Context, id int64) error
}

type LocationService struct {
	DB bun.IDB
}

func NewLocationService(db *bun.DB) ILocationService {
	return &LocationService{
		DB: db,
	}
}

func (s *LocationService) GetAll(ctx context.Context) ([]model.Location, int, error) {
	var locations []model.Location
	count, err := s.DB.NewSelect().Model(&locations).ScanAndCount(ctx)
	return locations, count, err
}

func (s *LocationService) GetByID(ctx context.Context, id int64) (model.Location, error) {
	var data model.Location
	err := s.DB.NewSelect().Model(&data).Where("id = ?", id).Scan(ctx)
	return data, err
}

func (s *LocationService) Create(ctx context.Context, m *model.Location) error {
	_, err := s.DB.NewInsert().Model(m).Exec(ctx)
	return err
}

func (s *LocationService) Update(ctx context.Context, m model.Location) error {
	_, err := s.DB.NewUpdate().Model(&m).WherePK().Exec(ctx)
	return err
}

func (s *LocationService) Delete(ctx context.Context, id int64) error {
	_, err := s.DB.NewDelete().Model(&model.Location{}).Where("id = ?", id).Exec(ctx)
	return err
}
