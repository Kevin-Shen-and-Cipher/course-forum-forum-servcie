package migrations

import (
	"course-forum/infra/database"
	"course-forum/infra/logger"
	"course-forum/models"
)

// Migrate Add list of model add for migrations
func Migrate() {
	var migrationModels = []interface{}{&models.Post{}, &models.Tag{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		logger.Errorf("Migration error: %s", err)
	}
}
