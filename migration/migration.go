package migration

import "route-planner/model"

func GetMigrationModels() []interface{} {
	return []interface{}{
		&model.Location{},
	}
}
