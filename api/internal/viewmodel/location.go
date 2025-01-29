package viewmodel

import (
	"math"
	"route-planner/model"
)

type LocationRequest struct {
	Name      string  `json:"name" validate:"required,min=2,max=100"`
	Latitude  float64 `json:"latitude" validate:"required,latitude"`
	Longitude float64 `json:"longitude" validate:"required,longitude"`
	Color     string  `json:"color" validate:"required,hexcolor"`
}

func (vm LocationRequest) ToDBModel(m model.Location) model.Location {
	m.Name = vm.Name
	m.Latitude = vm.Latitude
	m.Longitude = vm.Longitude
	m.Color = vm.Color

	return m
}

type LocationResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Color     string  `json:"color"`
}

func (vm LocationResponse) ToViewModel(m model.Location) LocationResponse {
	vm.ID = m.ID
	vm.Name = m.Name
	vm.Latitude = m.Latitude
	vm.Longitude = m.Longitude
	vm.Color = m.Color

	return vm
}

type LocationSortByDistanceRequest struct {
	Latitude  float64 `json:"latitude" validate:"required,latitude"`
	Longitude float64 `json:"longitude" validate:"required,longitude"`
}

func (vm LocationSortByDistanceRequest) ToDBModel(m model.Location) model.Location {
	m.Latitude = vm.Latitude
	m.Longitude = vm.Longitude

	return m
}

type LocationSortByDistanceResponse struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Color        string  `json:"color"`
	Distance     float64 `json:"distance"`
	DistanceUnit string  `json:"distance_unit"`
}

func (vm LocationSortByDistanceResponse) ToViewModel(m model.Location) LocationSortByDistanceResponse {
	vm.ID = m.ID
	vm.Name = m.Name
	vm.Latitude = m.Latitude
	vm.Longitude = m.Longitude
	vm.Color = m.Color

	if m.Distance < 1.0 {
		vm.Distance = math.Round(m.Distance*1000*100) / 100 // Convert to meters and round to 2 decimal places
		vm.DistanceUnit = "m"
	} else {
		vm.Distance = math.Round(m.Distance*100) / 100 // Round to 2 decimal places
		vm.DistanceUnit = "km"
	}

	return vm
}
