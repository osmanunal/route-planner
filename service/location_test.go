package service

import (
	"context"
	"database/sql"
	"route-planner/model"
	basemodel "route-planner/pkg/model"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func setupTestDB(t *testing.T) *bun.DB {
	sqldb, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	// Create tables with custom UUID handling for SQLite
	ctx := context.Background()
	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS locations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid CHAR(36) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			name TEXT NOT NULL,
			latitude REAL NOT NULL,
			longitude REAL NOT NULL,
			color TEXT NOT NULL,
			h3_index INTEGER NOT NULL DEFAULT 0
		)
	`)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func createTestLocation(t *testing.T, service ILocationService) *model.Location {
	location := &model.Location{
		BaseModel: basemodel.BaseModel{
			UUID: uuid.New(),
		},
		Name:      "Test Location",
		Latitude:  41.0082,
		Longitude: 28.9784,
		Color:     "#FF0000",
	}

	err := service.Create(context.Background(), location)
	assert.NoError(t, err)
	assert.NotZero(t, location.ID)

	return location
}

func TestLocationService_Create(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	service := NewLocationService(db)
	ctx := context.Background()

	// Test Create
	location := &model.Location{
		BaseModel: basemodel.BaseModel{
			UUID: uuid.New(),
		},
		Name:      "Test Location",
		Latitude:  41.0082,
		Longitude: 28.9784,
		Color:     "#FF0000",
	}

	err := service.Create(ctx, location)
	assert.NoError(t, err)
	assert.NotZero(t, location.ID)

	// Verify creation
	retrieved, err := service.GetByID(ctx, location.ID)
	assert.NoError(t, err)
	assert.Equal(t, location.Name, retrieved.Name)
	assert.Equal(t, location.Latitude, retrieved.Latitude)
	assert.Equal(t, location.Longitude, retrieved.Longitude)
	assert.Equal(t, location.Color, retrieved.Color)
}

func TestLocationService_GetByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	service := NewLocationService(db)
	location := createTestLocation(t, service)

	// Test GetByID
	retrieved, err := service.GetByID(context.Background(), location.ID)
	assert.NoError(t, err)
	assert.Equal(t, location.Name, retrieved.Name)
	assert.Equal(t, location.Latitude, retrieved.Latitude)
	assert.Equal(t, location.Longitude, retrieved.Longitude)
	assert.Equal(t, location.Color, retrieved.Color)

	// Test GetByID with non-existent ID
	_, err = service.GetByID(context.Background(), 99999)
	assert.Error(t, err)
}

func TestLocationService_Update(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	service := NewLocationService(db)
	location := createTestLocation(t, service)
	ctx := context.Background()

	// Test Update
	location.Name = "Updated Location"
	location.Color = "#00FF00"
	err := service.Update(ctx, *location)
	assert.NoError(t, err)

	// Verify update
	updated, err := service.GetByID(ctx, location.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Location", updated.Name)
	assert.Equal(t, "#00FF00", updated.Color)
}

func TestLocationService_GetAll(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	service := NewLocationService(db)
	ctx := context.Background()

	// Test empty list
	locations, count, err := service.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
	assert.Len(t, locations, 0)

	// Create test locations
	createTestLocation(t, service)
	createTestLocation(t, service)

	// Test GetAll with data
	locations, count, err = service.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)
	assert.Len(t, locations, 2)
}

func TestLocationService_Delete(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	service := NewLocationService(db)
	location := createTestLocation(t, service)
	ctx := context.Background()

	// Test Delete
	err := service.Delete(ctx, location.ID)
	assert.NoError(t, err)

	// Verify deletion
	_, err = service.GetByID(ctx, location.ID)
	assert.Error(t, err)

	// Test Delete non-existent record
	err = service.Delete(ctx, 99999)
	assert.NoError(t, err) // Silme işlemi var olmayan kayıt için de başarılı sayılır
}
