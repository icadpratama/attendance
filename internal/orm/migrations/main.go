package migrations

import (
	"fmt"

	log "github.com/icadpratama/attendance/internal/logger"

	"github.com/icadpratama/attendance/internal/orm/migration/jobs"
	"github.com/icadpratama/attendance/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}

func ServiceAutoMigration(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("[Migration.InitSchema] Initializing database schema")
		switch db.Dialect().GetName() {
		case "postgres":
			db.Exec("create extension \"uuid-ossp\";")

		}
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}

		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
	})
	return m.Migrate()
}
