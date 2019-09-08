package jobs

import (
	"github.com/icadpratama/attendance/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	uname = "irsyadjpratamap"
	fname = "Irsyad Jamal"
	lname = "Pratama Putra"
	nname = "Irsyad"
	description = "This is the first user ever!"
	location = "Palbatu"
	firstUser   *models.User = &models.User{
		Email: "irsyadjpratamap@gmail.com",
		Name: &uname,
		FirstName: &fname,
		LastName: &lname,
		NickName: &nname,
		Description: &description,
		Location: &location,
	}
)

var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}