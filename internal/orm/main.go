package orm

import (
    log "github.com/icadpratama/attendance/internal/logger"

    "github.com/icadpratama/attendance/internal/orm/migrations"

    "github.com/icadpratama/attendance/pkg/utils"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/jinzhu/gorm"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

type ORM struct {
    DB *gorm.DB
}

func init() {
    dialect = utils.MustGet("GORM_DIALECT")
    dsn = utils.MustGet("GORM_CONNECTION_DSN")
    seedDB = utils.MustGetBool("GORM_SEED_DB")
    logMode = utils.MustGetBool("GORM_LOGMODE")
    autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

func Factory() (*ORM, error) {
    db, err := gorm.Open(dialect, dsn)
    if err != nil {
        log.Panic("[ORM] err: ", err)
    }
    orm := &ORM{
        DB: db,
	}

    db.LogMode(logMode)

	if autoMigrate {
        err = migrations.ServiceAutoMigration(orm.DB)
    }

	log.Info("[ORM] Database connection initialized.")
    return orm, err
}