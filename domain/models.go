package domain

import (
	"github.com/rest-api/configs"
	"github.com/rest-api/infra/logger"
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
