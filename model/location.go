package model

import "route-planner/pkg/model"

type Location struct {
	model.BaseModel

	Name      string
	Latitude  float64
	Longitude float64
	Color     string
}
