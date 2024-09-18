package domain

import (
	"github.com/aro/configs"
	"github.com/aro/infra/logger"
)

// Auto-migrate database models
func MigrateModels(db *configs.Db) bool {
	// Add any new model to migrate here
	var models = []interface{}{}

	for _, s := range models {
		err := db.AutoMigrate(s)
		if err != nil {
			logger.Errorf("Failed to automigrate err: %s", err.Error())
			return false
		}
	}

	return true
}
