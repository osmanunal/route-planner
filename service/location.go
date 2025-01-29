package service

import (
	"context"
	"route-planner/model"
	"sort"

	"github.com/uptrace/bun"
)

type ILocationService interface {
	GetAll(ctx context.Context) ([]model.Location, int, error)
	GetByID(ctx context.Context, id int64) (model.Location, error)
	Create(ctx context.Context, m *model.Location) error
	Update(ctx context.Context, m model.Location) error
	Delete(ctx context.Context, id int64) error
	GetRoute(ctx context.Context, route *model.Location) ([]model.Location, int, error)
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

func (s *LocationService) GetRoute(ctx context.Context, route *model.Location) ([]model.Location, int, error) {
	var locations []model.Location

	_, err := s.DB.NewSelect().Model(&locations).ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	for i := range locations {
		locations[i].HaversineDistance(route.Latitude, route.Longitude)

	}

	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Distance < locations[j].Distance
	})

	return locations, len(locations), nil
}
